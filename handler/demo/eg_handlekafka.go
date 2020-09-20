package demo

import (
	"context"
	"fmt"
	"time"

	"github.com/Shopify/sarama"
	"github.com/gin-gonic/gin"
	"github.com/pquerna/ffjson/ffjson"

	cfg "use-gin/config"
	"use-gin/db/kafka"
	"use-gin/errcode"
	"use-gin/handler"
	"use-gin/logger"
)

const consumeTaskTimeout = 6 * time.Minute

type produceTask struct {
	ID       string
	StrDemo1 string `json:"str_demo1"`
	StrDemo2 string `json:"str_demo2"`
}

type consumeTask struct {
	ID       string
	StrDemo1 string `json:"str_demo1"`
	StrDemo2 string `json:"str_demo2"`
}

type consumeHandler struct {
	c         *gin.Context
	taskID    string
	ctxCancel func()
}

func (consumeHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (consumeHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (h consumeHandler) ConsumeClaim(
	_ sarama.ConsumerGroupSession,
	claim sarama.ConsumerGroupClaim,
) error {
	requestID, ok := h.c.Get("requestID")
	if !ok {
		requestID = "null"
	}

	for v := range claim.Messages() {
		// do not crash when received the invalid topic messages.
		defer func() {
			if err := recover(); err != nil {
				logger.Log.Errorf("request_id: %s, meet up panic: %v", requestID, err)
			}
		}()

		var ct consumeTask
		if err := ffjson.Unmarshal(v.Value, &ct); err != nil {
			logger.Log.Errorf(
				"request_id: %s, kafka message unmarshal: %v, error: %v", requestID, v, err)
			continue
		}
		logger.Log.Debugf(
			"request_id: %s, received from %s: %+v",
			requestID,
			cfg.Conf().Kafka.ConsumerTopic,
			ct,
		)

		if ct.ID == h.taskID {

			// your specific logic

			h.ctxCancel()
			return nil
		}
	}
	return nil
}

// HandleKafkaDemo a demo of handle kafka
func HandleKafkaDemo(c *gin.Context) {
	// producer
	produceTask := produceTask{}
	if err := addTaskToTopic(c, produceTask); err != nil {
		err1 := errcode.New(errcode.InternalServerError, err)
		handler.SendResponse(c, err1, nil)
		return
	}
	handler.SendResponse(c, nil, nil)

	// consumer
	taskID := produceTask.ID
	consumeTopics := []string{cfg.Conf().Kafka.ConsumerTopic}
	ctx, cancel := context.WithCancel(context.Background())
	handler := consumeHandler{c, taskID, cancel}

	go kafka.ConsumeTopics(
		ctx,
		cancel,
		consumeTaskTimeout,
		consumeTopics,
		handler,
	)
}

func addTaskToTopic(c *gin.Context, task produceTask) error {
	requestID, ok := c.Get("requestID")
	if !ok {
		requestID = "null"
	}

	reqStr, err := ffjson.Marshal(task)
	if err != nil {
		return fmt.Errorf("task serialization error: %v", err)
	}

	topic := cfg.Conf().Kafka.ProducerTopic
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(reqStr),
	}
	logger.Log.Debugf(
		"request_id: %s, put task into topic %s, task: %s", requestID, topic, reqStr)

	kafka.Producer.Input() <- msg
	return nil
}
