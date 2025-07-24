package main

import (
	"circuitcamel.com/internal/server"
	"circuitcamel.com/internal/utils"
)

func main() {
	conf := utils.LoadConfig()
	server.StartServer(conf)
}
