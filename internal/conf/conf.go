package conf

import (
	"os"
	"strconv"

	"github.com/rs/zerolog/log"
)

const (
	hostKey = "FB_HOST"
	portKey = "FB_PORT"
)

type Config struct {
	Host string
	Port string
	Env  string
}

func NewConfig(env string) Config {
	host, ok := os.LookupEnv(hostKey)
	if !ok || host == "" {
		logAndPanic(hostKey)
	}

	port, ok := os.LookupEnv(portKey)
	if !ok || port == "" {
		if _, err := strconv.Atoi(port); err != nil {
			logAndPanic(portKey)
		}
	}

	return Config{
		Host: host,
		Port: port,
		Env:  env,
	}
}

func NewTestConfig() Config {
	testConfig := NewConfig("dev")
	return testConfig
}

func logAndPanic(envVar string) {
	log.Panic().Str("envVar", envVar).Msg("ENV variable not set or value not valid")
}
