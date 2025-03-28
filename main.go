package main

import (
	"flag"
	"gpk/logger"
	"log"
	"supplementary-inspection/basicdata"
	"supplementary-inspection/mq"
	"supplementary-inspection/pool"
	"supplementary-inspection/route"

	"github.com/spf13/viper"
)

var (
	configFile = flag.String("c", "supplementary.yaml", "config file path")
	serverPort = flag.Int("p", 8000, "server port")
)

func getConfig() {
	flag.Parse()
	viper.SetConfigFile(*configFile)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	var cfg logger.LogConfig
	viper.UnmarshalKey("logger", &cfg)
	logger.Init(&cfg)
	logger.Info("read config success")
	logger.Info(viper.AllSettings())

	viper.UnmarshalKey("http-host", &pool.HttpHost)
	viper.UnmarshalKey("analysis-host", &pool.AnalysisHost)
	viper.UnmarshalKey("mq", &mq.Mqconfig)

	pool.AnalyzeTimeout = viper.GetInt("settings.analysis-timeout")
	basicdata.BaseUrl = viper.GetString("microservices.basic-data")
}

func main() {
	getConfig()
	go pool.Run()
	go basicdata.Init()
	go mq.Run()
	route.RunHttpServer(*serverPort)
}
