package config

type Config interface {
	Port() string
	LogLevel() string
}

type config struct {
	port     string
	logLevel string
}

func (c config) LogLevel() string {
	return c.logLevel
}

func (c config) Port() string {
	return c.port
}

func MustLoad() Config {
	return &config{
		port:     "8080",
		logLevel: "debug",
	}
}
