package gkafka

import (
	"context"
	"errors"
	"strings"

	"github.com/IBM/sarama"

	mq "github.com/daheige/athena/internal/infras/prioritymq"
)

var _ sarama.ConsumerGroupHandler = (*consumerGroupHandler)(nil)

// ErrSubHandlerInvalid sub handler invalid
var ErrSubHandlerInvalid = errors.New("subHandler is nil")

// consumerGroupHandler impl sarama.ConsumerGroupHandler
// consumer groups require Version to be >= V0_10_2_0
type consumerGroupHandler struct {
	ctx               context.Context
	groupID           string // consumer group group_id
	commitOffsetBlock bool
	logger            mq.Logger

	// 优先级队列
	highMsgChan   chan *sarama.ConsumerMessage
	mediumMsgChan chan *sarama.ConsumerMessage
	lowMsgChan    chan *sarama.ConsumerMessage
	normalMsgChan chan *sarama.ConsumerMessage

	stop    chan struct{}
	handler mq.SubHandler
}

// Setup is run at the beginning of a new session, before ConsumeClaim.
func (c *consumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error {
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
// but before the offsets are committed for the very last time.
func (c *consumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
// Once the Messages() channel is closed, the Handler must finish its processing
// loop and exit.
func (c *consumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession,
	claim sarama.ConsumerGroupClaim) error {
	defer mq.Recovery(c.logger)

	// 启动消费逻辑
	c.logger.Printf("kafka start consume message...")
	go func() {
		c.startConsume(c.ctx, sess)
	}()

	// note: the message key of kafka may be nil,if c.key is not empty,it must be eq msg.key
	for msg := range claim.Messages() {
		c.logger.Printf("kafka received topic:%v group_id:%v partition:%d offset:%d key:%s -- value:%s\n",
			msg.Topic, c.groupID, msg.Partition, msg.Offset, msg.Key, msg.Value)
		// put msg into channel
		c.putMsgToChan(msg)
	}

	return nil
}

// 根据队列优先级将msg放入对应的chan中
func (c *consumerGroupHandler) putMsgToChan(msg *sarama.ConsumerMessage) {
	switch true {
	case strings.HasSuffix(msg.Topic, "high"):
		c.highMsgChan <- msg
	case strings.HasSuffix(msg.Topic, "medium"):
		c.mediumMsgChan <- msg
	case strings.HasSuffix(msg.Topic, "low"):
		c.lowMsgChan <- msg
	default:
		c.normalMsgChan <- msg
	}
}

// start consumer msg
func (c *consumerGroupHandler) startConsume(ctx context.Context, sess sarama.ConsumerGroupSession) {
	defer mq.Recovery(c.logger)

	// 启动优先级消费逻辑
	for {
		select {
		case msg := <-c.highMsgChan: // 优先处理高优先级
			err := c.handlerMsg(c.ctx, msg, sess)
			if err != nil {
				c.logger.Printf("failed to handler high msg topic:%v group_id:%v err:%v\n", msg.Topic, c.groupID, err)
			}
		case msg := <-c.mediumMsgChan: // 处理中等优先级
			select {
			case highMsg := <-c.highMsgChan:
				// 先将中优先级消息放回队列
				c.logger.Printf(
					"put the medium message topic:%v group_id:%v back into the origin channel", msg.Topic, c.groupID,
				)
				c.mediumMsgChan <- msg

				// 处理高优先级的消息
				err := c.handlerMsg(c.ctx, highMsg, sess)
				if err != nil {
					c.logger.Printf(
						"failed to handler high msg topic:%v group_id:%v err:%v\n", highMsg.Topic, c.groupID, err,
					)
				}
			default:
				// 如果没有高优先级的消息，就直接消费
				err := c.handlerMsg(c.ctx, msg, sess)
				if err != nil {
					c.logger.Printf(
						"failed to handler medium msg topic:%v group_id:%v err:%v\n",
						msg.Topic, c.groupID, err,
					)
				}
			}
		case msg := <-c.lowMsgChan:
			select {
			case highMsg := <-c.highMsgChan:
				// 将低优先级消息放回队列
				c.logger.Printf(
					"put the low message topic:%v group_id:%v back into the origin channel", msg.Topic, c.groupID,
				)
				c.lowMsgChan <- msg

				// 处理高优先级的消息
				err := c.handlerMsg(c.ctx, highMsg, sess)
				if err != nil {
					c.logger.Printf(
						"failed to handler high msg topic:%v group_id:%v err:%v\n", highMsg.Topic, c.groupID, err,
					)
				}
			default:
				err := c.handlerMsg(c.ctx, msg, sess)
				if err != nil {
					c.logger.Printf(
						"failed to handler low msg topic:%v group_id:%v err:%v\n", msg.Topic, c.groupID, err,
					)
				}
			}
		case msg := <-c.normalMsgChan:
			err := c.handlerMsg(c.ctx, msg, sess)
			if err != nil {
				c.logger.Printf(
					"failed to handler normal msg topic:%v group_id:%v err:%v\n", msg.Topic, c.groupID, err,
				)
			}
		case <-c.stop:
			c.logger.Printf("kafka consumer group has stopped\n")
			return
		}
	}
}

// handlerMsg consumer msg
func (c *consumerGroupHandler) handlerMsg(ctx context.Context, msg *sarama.ConsumerMessage,
	sess sarama.ConsumerGroupSession) error {
	defer mq.Recovery(c.logger)

	c.logger.Printf(
		"handler kafka consume message,topic:%v group_id:%v partition:%d offset:%d\n",
		msg.Topic, c.groupID, msg.Partition, msg.Offset,
	)
	err := c.handler(ctx, msg.Value)
	if err != nil {
		return err
	}

	// mark message as processed
	sess.MarkMessage(msg, "")

	// Commit the offset to the backend for kafka
	// Note: calling Commit performs a blocking synchronous operation.
	if c.commitOffsetBlock {
		sess.Commit()
	}

	return nil
}
