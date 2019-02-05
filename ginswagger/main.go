package main

import (
	"github.com/peterlamar/go-examples/ginswagger/api"
	_ "github.com/peterlamar/go-examples/ginswagger/docs"
	log "github.com/sirupsen/logrus"
)

// @title swagger example
// @version 0.0.1
// @description This API is used to show go-swagger capability
// @host localhost:8080
// @BasePath /
func main() {
	log.SetLevel(log.InfoLevel)
	log.Debug("Swagger example service starting")

	router := api.SetupRouter() // Setup router paths
	router.Run()                // listen and serve on 0.0.0.0:8080

	log.Debug("Swagger example service stopping")
}
