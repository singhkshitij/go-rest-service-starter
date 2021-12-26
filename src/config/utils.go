package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func panicIfError(err error) {
	if err != nil {
		panic(fmt.Errorf("unable to load config: %v", err))
	}
}

func checkKey(key string) {
	if !viper.IsSet(key) {
		panicIfError(fmt.Errorf("%s key is not set", key))
	}
}

func getStringOrPanic(key string) string {
	checkKey(key)
	v := viper.GetString(key)
	return v
}

func getBoolOrPanic(key string) bool {
	checkKey(key)
	v := viper.GetBool(key)
	return v
}
