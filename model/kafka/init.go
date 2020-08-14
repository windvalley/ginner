package kafka

import (
	"context"
	"time"

	"github.com/Shopify/sarama"

	cfg "use-gin/config"
	"use-gin/logger"
)

var (
	Consumer sarama.ConsumerGroup
	Producer sarama.AsyncProducer
)

func InitKafkaConsumer() {
	var err error
	config := sarama.NewConfig()
	config.Net.KeepAlive = 10 * time.Second
	config.Producer.Return.Successes = true
	config.Producer.Timeout = 20 * time.Second
	config.Version = sarama.V2_0_0_0

	Consumer, err = sarama.NewConsumerGroup(
		cfg.Config().Kafka.Brokers,
		cfg.Config().Kafka.ConsumerGroup,
		config,
	)
	if err != nil {
		panic(err)
	}
}

func InitKafkaProducer() {
	var err error
	config := sarama.NewConfig()
	config.Net.KeepAlive = 10 * time.Second
	config.Producer.Return.Successes = true
	config.Producer.Timeout = 20 * time.Second

	Producer, err = sarama.NewAsyncProducer(cfg.Config().Kafka.Brokers, config)
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
