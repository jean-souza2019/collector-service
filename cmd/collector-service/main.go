package main

import (
	"github.com/jean-souza2019/collector-service/api/server"
	"github.com/jean-souza2019/collector-service/configs"
	"github.com/jean-souza2019/collector-service/pkg/mongo"
)

func main() {
	configs.LoadEnvironments()

	mongo.Connect()

	server.Initialize()
}
