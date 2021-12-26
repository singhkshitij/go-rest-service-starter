package config

type Configuration struct {
	logLevel       string
}

func LogLevel() string {
	return (*config).logLevel
}
