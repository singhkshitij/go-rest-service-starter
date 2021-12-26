package api

import (
	"github.com/singhkshitij/golang-rest-service-starter/src/config"
	"github.com/singhkshitij/golang-rest-service-starter/src/logger"
)

func NewServer() {
	logger.Info("Server started " + config.LogLevel())
}
