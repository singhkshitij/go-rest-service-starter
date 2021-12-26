package config

import "github.com/spf13/viper"

type RedisConfig struct {
	Enabled bool
	Host    string
	Port    string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Name     string
}

type Configuration struct {
	env         string
	logLevel    string
	port        string
	redisConfig RedisConfig
	dbConfig    DatabaseConfig
}

func GetConfiguration() *Configuration {
	return &Configuration{
		env:      getStringOrPanic("ENV"),
		logLevel: getStringOrPanic("LOG_LEVEL"),
		port:     getStringOrPanic("PORT"),
		redisConfig: RedisConfig{
			Enabled: getBoolOrPanic("REDIS_ENABLED"),
			Host:    getStringOrPanic("REDIS_HOST"),
			Port:    getStringOrPanic("REDIS_PORT"),
		},
		dbConfig: DatabaseConfig{
			Host:     getStringOrPanic("DATABASE_HOST"),
			Port:     getStringOrPanic("DATABASE_PORT"),
			Username: getStringOrPanic("DATABASE_USER"),
			Password: viper.GetString("DATABASE_PASSWORD"),
			Name:     getStringOrPanic("DATABASE_NAME"),
		},
	}
}

func setDefaults() {
	viper.SetDefault("ENV", "dev")
	viper.SetDefault("LOG_LEVEL", "debug")
	viper.SetDefault("PORT", "8080")
	viper.SetDefault("REDIS_ENABLED", "false")
	viper.SetDefault("REDIS_PORT", "6379")
	viper.SetDefault("REDIS_HOST", "localhost")
}

func LogLevel() string {
	return (*config).logLevel
}

func GetEnv() string {
	return (*config).env
}

func Port() string {
	return (*config).port
}

func RedisConf() RedisConfig {
	return (*config).redisConfig
}

func DbConfig() DatabaseConfig{
	return (*config).dbConfig
}
