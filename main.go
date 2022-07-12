package main

import (
	"fizzbuzz/internal/conf"
	"fizzbuzz/internal/server"
)

func main() {
	server.Start(conf.NewConfig())
}
