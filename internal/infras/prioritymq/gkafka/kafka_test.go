package gkafka

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/daheige/athena/internal/infras/prioritymq"
)

func TestMQPublish(t *testing.T) {
	brokers := []string{"localhost:9092"}
	mq, err := New(
		brokers,
		WithLogger(prioritymq.LoggerFunc(log.Printf)),
		WithProducerTimeout(10*time.Second),
		WithGracefulWait(3*time.Second),
		WithDialTimeout(10*time.Second),
	)

	if err != nil {
		log.Fatalf("init mq error: %v", err)
	}

	for i := 0; i < 100; i++ {
		err = mq.Publish(
			context.Background(),
			prioritymq.NewTopicName("my-test", prioritymq.High),
			[]byte(fmt.Sprintf("hello world %d from high topic", i)),
			prioritymq.WithPubName("my-test"),
		)
		if err != nil {
			log.Printf("publish error: %v", err)
			continue
		}
		err = mq.Publish(
			context.Background(),
			prioritymq.NewTopicName("my-test", prioritymq.Low),
			[]byte(fmt.Sprintf("hello world %d from low topic", i)),
			prioritymq.WithPubName("my-test1"),
		)
		if err != nil {
			log.Printf("publish error: %v", err)
			continue
		}

		err = mq.Publish(
			context.Background(),
			prioritymq.NewTopicName("my-test", prioritymq.Medium),
			[]byte(fmt.Sprintf("hello world %d from medium topic", i)),
			prioritymq.WithPubName("my-test2"),
		)
		if err != nil {
			log.Printf("publish error: %v", err)
			continue
		}
	}

}

func TestMQSubscribe(t *testing.T) {
	brokers := []string{"localhost:9092"}
	mq, err := New(
		brokers,
		WithLogger(prioritymq.LoggerFunc(log.Printf)),
		WithConsumerAutoCommitInterval(3*time.Second),
		WithConsumerOffsetsInitial(-1),
		WithGracefulWait(3*time.Second),
		WithDialTimeout(10*time.Second),
	)

	if err != nil {
		log.Fatalf("init mq error: %v", err)
	}

	topics := []string{
		prioritymq.NewTopicName("my-test", prioritymq.High),
		prioritymq.NewTopicName("my-test", prioritymq.Medium),
		prioritymq.NewTopicName("my-test", prioritymq.Low),
	}
	err = mq.Subscribe(context.Background(), topics, "group-1", func(ctx context.Context, value []byte) error {
		log.Println("current msg value:", string(value))
		return nil
	})
}
