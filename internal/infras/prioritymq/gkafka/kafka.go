package gkafka

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/IBM/sarama"

	mq "github.com/daheige/athena/internal/infras/prioritymq"
	"github.com/daheige/athena/internal/infras/prioritymq/backoff"
)

var _ mq.MQ = (*kafkaImpl)(nil)

type kafkaImpl struct {
	client       sarama.Client
	logger       mq.Logger
	stop         chan struct{}
	gracefulWait time.Duration
}

// New 创建一个mq实例
func New(brokers []string, opts ...Option) (mq.MQ, error) {
	opt := Options{
		brokers:                    brokers,
		gracefulWait:               3 * time.Second,
		consumerAutoCommitInterval: 3 * time.Second,
		producerTimeout:            10 * time.Second,
		logger:                     mq.DummyLogger,
		dialTimeout:                30 * time.Second,
		consumerOffsetsInitial:     sarama.OffsetNewest,
	}

	for _, o := range opts {
		o(&opt)
	}

	k := &kafkaImpl{
		logger:       opt.logger,
		gracefulWait: opt.gracefulWait,
		stop:         make(chan struct{}, 1),
	}

	// create client
	var err error
	k.client, err = initClient(opt)
	if err != nil {
		return nil, err
	}

	return k, nil
}

// Publish 发送消息
func (k *kafkaImpl) Publish(ctx context.Context, topic string, msg []byte, opts ...mq.PubOption) error {
	select {
	case <-k.stop:
		return errors.New("current mq has stopped")
	default:
	}

	// publish options
	opt := mq.PublishOptions{}
	for _, o := range opts {
		o(&opt)
	}

	message := &sarama.ProducerMessage{
		Topic: topic, Value: sarama.ByteEncoder(msg),
	}

	if opt.Key != "" {
		// The partitioning key for this message. Pre-existing Encoders include
		// StringEncoder and ByteEncoder.
		message.Key = sarama.StringEncoder(opt.Key)
	}

	// create producer
	var producer sarama.SyncProducer
	producer, err := sarama.NewSyncProducerFromClient(k.client)
	if err != nil {
		k.logger.Printf("new kafka producer err:%v\n", err)
		return err
	}

	defer func() {
		_ = producer.Close()
	}()

	// send message
	var (
		partition int32
		offset    int64
	)
	partition, offset, err = producer.SendMessage(message)
	if err != nil {
		return err
	}

	k.logger.Printf("kafka producer partitionID: %d; offset:%d, value: %s\n", partition, offset, string(msg))

	return nil
}

// Subscribe 消费消息
func (k *kafkaImpl) Subscribe(ctx context.Context, topics []string, groupID string,
	handler mq.SubHandler, opts ...mq.SubOption) error {
	if handler == nil {
		return ErrSubHandlerInvalid
	}

	opt := mq.SubscribeOptions{
		Name:       groupID, // consumer group_id
		BufferSize: 1024,    // 默认1024
	}

	for _, o := range opts {
		o(&opt)
	}

	if ctx == nil {
		ctx = context.Background()
	}

	k.logger.Printf("subscribe message from kafka receive topics:%v group_id:%v msg...\n", topics, opt.Name)

	consumerGroup, err := sarama.NewConsumerGroupFromClient(opt.Name, k.client)
	if err != nil {
		panic(fmt.Errorf("new kafka consumer group:%v topics:%v err:%v", opt.Name, topics, err))
	}

	defer func() {
		_ = consumerGroup.Close()
	}()

	done := make(chan struct{}, 1)
	go func() {
		defer mq.Recovery(k.logger)
		defer func() {
			done <- struct{}{}
		}()

		consumerHandler := &consumerGroupHandler{
			ctx:               ctx,
			groupID:           opt.Name,
			commitOffsetBlock: opt.CommitOffsetBlock,
			logger:            k.logger,
			handler:           handler,
			stop:              k.stop,
			highMsgChan:       make(chan *sarama.ConsumerMessage, opt.BufferSize),
			mediumMsgChan:     make(chan *sarama.ConsumerMessage, opt.BufferSize),
			lowMsgChan:        make(chan *sarama.ConsumerMessage, opt.BufferSize),
			normalMsgChan:     make(chan *sarama.ConsumerMessage, opt.BufferSize),
		}

		for {
			select {
			case <-k.stop:
				return
			case consumeErr := <-consumerGroup.Errors():
				k.logger.Printf("kafka received topics:%v channel:%v handler msg err:%v\n",
					topics, opt.Name, consumeErr)
				backoff.Sleep(1)
			default:
				// Consume() should be called continuously in an infinite loop
				// Because Consume() needs to be executed again after each rebalance to restore the connection
				// The Join Group request is not initiated until the Consume starts. If the current consumer
				// becomes the leader of the consumer group after joining, the rebalance process will also be
				// performed to re-allocate
				// The topics and partitions that each consumer group in the group needs to consume,
				// and the consumption starts after the last Sync Group
				consumeErr := consumerGroup.Consume(ctx, topics, consumerHandler)
				if consumeErr != nil {
					k.logger.Printf("received topics:%v channel:%v handler msg err:%v\n",
						topics, opt.Name, consumeErr)
					continue
				}
			}
		}
	}()

	<-done
	return nil
}

// Shutdown 平滑退出
func (k *kafkaImpl) Shutdown(ctx context.Context) error {
	k.gracefulStop(ctx)
	close(k.stop)
	return nil
}

func (k *kafkaImpl) gracefulStop(ctx context.Context) {
	defer k.logger.Printf("subscribe msg exit successfully\n")

	if ctx == nil {
		ctx = context.Background()
	}

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(ctx, k.gracefulWait)
	defer cancel()

	// Doesn't block if no service run, but will otherwise wait
	// until the timeout deadline.
	// Optionally, you could run it in a goroutine and block on
	// if your application should wait for other services
	// to finalize based on context cancellation.
	done := make(chan struct{}, 1)
	var err = make(chan error, 1)
	go func() {
		defer close(done)

		err <- k.client.Close()
	}()

	<-done
	<-ctx.Done()

	k.logger.Printf("subscribe msg shutting down,err:%v\n", <-err)
}

func initClient(o Options) (sarama.Client, error) {
	// kafka sarama config
	config := sarama.NewConfig()
	config.Net.DialTimeout = o.dialTimeout
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	config.Producer.Timeout = o.producerTimeout

	// consumer config
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.AutoCommit.Enable = true
	config.Consumer.Offsets.AutoCommit.Interval = o.consumerAutoCommitInterval
	config.Consumer.Offsets.Initial = o.consumerOffsetsInitial
	config.Consumer.Fetch.Default = 1024 * 1024
	if o.user != "" { // user/pwd auth
		config.Net.SASL.Enable = true
		config.Net.SASL.User = o.user
		config.Net.SASL.Password = o.password
	}

	// create kafka client
	client, err := sarama.NewClient(o.brokers, config)
	if err != nil {
		return nil, err
	}

	return client, nil
}
