package main

import (
	"flag"
	"fmt"
	"log"

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
	fmt.Println("read config success")
	fmt.Println(viper.AllSettings())
}

func main() {
	fmt.Println(viper.GetString("microservices.basic-data"))
}
