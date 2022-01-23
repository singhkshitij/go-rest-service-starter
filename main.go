package main

import (
	"github.com/singhkshitij/golang-rest-service-starter/src/api"
	"github.com/singhkshitij/golang-rest-service-starter/src/cache"
	"github.com/singhkshitij/golang-rest-service-starter/src/config"
	"github.com/singhkshitij/golang-rest-service-starter/src/logger"
	"log"
)

func init() {
	config.InitConfig()
	logger.Must(logger.NewLogger("")) //replace with config file if logs need to be put in a log file
	err := cache.Setup()
	if err != nil {
		panic("err while setting up redis " + err.Error())
	}
}

func main() {
	server, err := api.NewServer()
	if err != nil {
		log.Fatal(err)
	}
	server.Start()
}
