package mq

import (
	"gpk/logger"
	"gpk/rabbitmq"
	"supplementary-inspection/basicdata"
	"time"

	"github.com/tidwall/gjson"
)

func NewPlatformComsumer() {
	defer func() {
		if r := recover(); r != nil {
			time.Sleep(10 * time.Second)
			NewPlatformComsumer()
		}
	}()
	cfg := &rabbitmq.ChannelConfig{
		Type:     "topic",
		Exchange: "platform",
		Queue:    "linkage_consumer",
		Key:      []string{"database.modify"},
	}
	MQ.NewConsumer(cfg, platformHandler)
}

func platformHandler(body []byte) {
	logger.Info("platform message: ", string(body))
	action := gjson.GetBytes(body, "action").String()
	switch action {
	case DatabaseModify:
		var ids []int64
		table := gjson.GetBytes(body, "data.table").String()
		action := gjson.GetBytes(body, "data.action").String()
		for _, v := range gjson.GetBytes(body, "data.ids").Array() {
			ids = append(ids, v.Int())
		}
		handleDataModify(table, action, ids...)
	}
}

func handleDataModify(table, action string, ids ...int64) {
	switch table {
	case "task":
		if action == "create" || action == "update" {
			basicdata.TaskMap.GetData()
		}
	case "patrolpoint":
		if action == "create" || action == "update" {
			basicdata.PatrolPointMap.GetData()
		}
	}
}
