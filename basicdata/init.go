package basicdata

import (
	"github.com/spf13/viper"
)

var baseUrl string

func Init() {
	baseUrl = viper.GetString("microservices.basic-data")
	PatrolPointMap.GetData()
	TaskMap.GetData()
}
