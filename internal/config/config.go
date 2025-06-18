package config

type Config interface {
	Port() int
	LogLevel() string
}

type config struct {
	port     int
	logLevel string
}

func (c config) LogLevel() string {
	return c.logLevel
}

func (c config) Port() int {
	return c.port
}

func MustLoad() Config {
	return config{
		port:     9345,
		logLevel: "debug",
	}
}
