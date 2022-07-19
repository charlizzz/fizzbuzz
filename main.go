package main

import (
	"fizzbuzz/internal/cli"
	"fizzbuzz/internal/conf"
	"fizzbuzz/internal/server"
)

func main() {
	env := cli.Parse()
	server.Start(conf.NewConfig(env))
}
