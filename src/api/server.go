package api

import (
	"fmt"
	"github.com/spf13/viper"
)

func NewServer() {
	fmt.Println("Server started", viper.Get("LOG_LEVEL"))
}

