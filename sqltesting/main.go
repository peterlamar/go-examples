package main

import (
	"github.com/peterlamar/go-examples/sqltesting/api"
	"github.com/peterlamar/go-examples/sqltesting/database"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.InfoLevel)
	log.SetReportCaller(true)
	log.Debug("Two Tier Cache API Service Starting")
	database.ConnectCache()
	database.ConnectDB()

	router := api.SetupRouter() // Setup router paths
	router.Run()                // listen and serve on 0.0.0.0:8080
	log.Debug("Two Tier Cache API Service Stopping")
}
