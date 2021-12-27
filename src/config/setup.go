package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var cfgFile string
var config *Configuration

func GetConfigFile() string {
	return cfgFile
}

// InitConfig reads in config file and ENV variables if set.
func InitConfig() {
	setDefaults()
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigName("config")
		viper.AddConfigPath("./")
		viper.AddConfigPath("../")
		viper.AddConfigPath("../../")
		viper.SetConfigType("yaml")
	}

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error while reading config file", err)
	}
	config = GetConfiguration()

	if viper.ConfigFileUsed() == "" {
		fmt.Println("Failed to load config file, falling back to env vars or default configs")
	} else {
		cfgFile = viper.ConfigFileUsed()
		fmt.Println("Config file used: " + cfgFile)
	}
}
