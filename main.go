package main

import (
	"flag"
	"gpk/logger"
	"log"
	"supplementary-inspection/pool"
	"supplementary-inspection/route"
	"supplementary-inspection/service"

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
	if dat := service.AnalyzeDatFile("test.dat"); dat != nil {
		dat.MakeFile()
	}
	go pool.Run()
	route.RunHttpServer(*serverPort)
}
