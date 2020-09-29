package kafka

import (
	"context"
	"time"

	"github.com/Shopify/sarama"

	cfg "ginner/config"
	"ginner/logger"
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
	Consumer, err = sarama.NewConsumerGroup(
		cfg.Conf().Kafka.Brokers,
		cfg.Conf().Kafka.ConsumerGroup,
		getConfig(),
	)
	if err != nil {
		logger.Log.Fatalf("init kafka consumer failed: %v", err)
	}
}

// InitProducer kafka producer initialization
func InitProducer() {
	var err error
	Producer, err = sarama.NewAsyncProducer(
		cfg.Conf().Kafka.Brokers,
		getConfig(),
	)
	if err != nil {
		logger.Log.Fatalf("init kafka producer failed: %v", err)
	}

	go func(Producer sarama.AsyncProducer) {
		errors := Producer.Errors()
		success := Producer.Successes()
		for {
			select {
			case err := <-errors:
				if err != nil {
					logger.Log.Errorf("kafka producer error: %v", err)
				}
			case <-success:
			}
		}
	}(Producer)
}

func getConfig() *sarama.Config {
	config := sarama.NewConfig()
	config.Net.KeepAlive = cfg.Conf().Kafka.Keepalive * time.Second
	config.Producer.Return.Successes = true
	config.Version = sarama.V2_0_0_0
	return config
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
			if err := Consumer.Consume(ctx, topics, handler); err != nil {
				logger.Log.Errorf("consume error: %v", err)
			}
		}
	}
}
