package config

type Configuration struct {
	logLevel string
	port     string
}

func GetConfiguration() *Configuration {
	return &Configuration{
		logLevel: getStringOrPanic("LOG_LEVEL"),
		port:     getStringOrPanic("PORT"),
	}
}

func LogLevel() string {
	return (*config).logLevel
}

func Port() string {
	return (*config).port
}
