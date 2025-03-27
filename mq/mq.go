package mq

import (
	"gpk/rabbitmq"
	"time"

	"gpk/logger"
)

var Mqconfig *rabbitmq.MQConfig
var MQ *rabbitmq.MQ

var (
	DatabaseModify    = "database.modify"
	FrontAcpointValue = "front.acpoint.value"
	FrontAcpointFinal = "front.acpoint.final"
)

func Run() {
	defer func() {
		if r := recover(); r != nil {
			time.Sleep(10 * time.Second)
			Run()
		}
	}()
	logger.Info("mq config: ", Mqconfig)
	MQ = rabbitmq.New(Mqconfig)

	// 定义生产者和消费者
	NewPlatformComsumer()
	NewProducer()
}

var Producer *rabbitmq.Producer

func NewProducer() {
	defer func() {
		if r := recover(); r != nil {
			time.Sleep(10 * time.Second)
			NewProducer()
		}
	}()
	Producer = MQ.NewProducer(&rabbitmq.ChannelConfig{
		Type:     "topic",
		Exchange: "platform",
	})
}

func PublishMsg(action string, data interface{}) error {
	msg := map[string]interface{}{
		"action": action,
		"data":   data,
	}
	return Producer.PublishMsgWithKey(action, msg)
}
