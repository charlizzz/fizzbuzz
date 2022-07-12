package conf

import (
	"log"
	"os"
	"strconv"
)

const (
	hostKey = "FB_HOST"
	portKey = "FB_PORT"
)

type Config struct {
	Host string
	Port string
}

func NewConfig() Config {
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
	}
}

func logAndPanic(envVar string) {
	log.Println("ENV variable not set or value not valid: ", envVar)
	panic(envVar)
}
