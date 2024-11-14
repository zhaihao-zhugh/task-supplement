package main

import (
	"flag"
	"gpk/logger"
	"log"
	_ "supplementary-inspection/data"
	"supplementary-inspection/pool"
	"supplementary-inspection/route"

	"github.com/spf13/viper"
)

var (
	configFile = flag.String("c", "supplementary.yaml", "config file path")
	serverPort = flag.Int("p", 8000, "server port")
)

func init() {
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
}

func main() {
	go pool.Run()
	route.RunHttpServer(*serverPort)
}
