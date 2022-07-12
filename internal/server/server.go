package server

import (
	"fizzbuzz/internal/conf"
)

func Start(cfg conf.Config) {
	router := setRouter()

	router.Run(":8080")
}
