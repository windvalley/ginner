package kafka

import (
	"context"
	"time"

	"github.com/Shopify/sarama"

	cfg "use-gin/config"
	"use-gin/logger"
)

var (
	// Consumer consumer client instance of kafka
	Consumer sarama.ConsumerGroup
	// Producer producer client instance of kafka
	Producer sarama.AsyncProducer
)

// InitConsumer kafka consumer initialization
func InitConsumer() {
	var err error
	config := sarama.NewConfig()
	config.Net.KeepAlive = 10 * time.Second
	config.Producer.Return.Successes = true
	config.Producer.Timeout = 20 * time.Second
	config.Version = sarama.V2_0_0_0

	Consumer, err = sarama.NewConsumerGroup(
		cfg.Conf().Kafka.Brokers,
		cfg.Conf().Kafka.ConsumerGroup,
		config,
	)
	if err != nil {
		panic(err)
	}
}

// InitProducer kafka producer initialization
func InitProducer() {
	var err error
	config := sarama.NewConfig()
	config.Net.KeepAlive = 10 * time.Second
	config.Producer.Return.Successes = true
	config.Producer.Timeout = 20 * time.Second

	Producer, err = sarama.NewAsyncProducer(cfg.Conf().Kafka.Brokers, config)
	if err != nil {
		panic(err)
	}

	go func(Producer sarama.AsyncProducer) {
		errors := Producer.Errors()
		success := Producer.Successes()
		for {
			select {
			case err := <-errors:
				if err != nil {
					logger.Log.Errorf("kafka run failed:%#v", err)
				}
			case <-success:
			}
		}
	}(Producer)
}

// ConsumeTopics start to consume the topics
func ConsumeTopics(
	ctx context.Context,
	cancel func(),
	timeout time.Duration,
	topics []string,
	handler sarama.ConsumerGroupHandler,
) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(timeout):
			logger.Log.Errorf("task %v timeout", handler)
			cancel()
		default:
			if err := Consumer.Consume(
				ctx, topics, handler); err != nil {
				logger.Log.Errorf("consume error: %v", err)
			}
		}
	}
}
