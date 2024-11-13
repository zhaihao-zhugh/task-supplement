package main

import (
	"flag"
	"log"

	logger "gpk/logger"

	"github.com/spf13/viper"
)

var (
	configFile = flag.String("c", "supplementary.yaml", "config file path")
	serverPort = flag.Int("p", 8080, "server port")
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
}

func main() {

}
