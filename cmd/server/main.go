package main

import (
	log "github.com/sirupsen/logrus"
	"munshee/services/chartofaccounts"
	"munshee/services/http"
	"os"
)

func init() {
	// TODO: Add a config for json or text based logs
	//log.SetFormatter(&log.JSONFormatter{})
	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func main() {
	log.Info("Munshee Starting Http Server")
	log.Info("Initializing the ChartOfAccounts Service")
	coaService := chartofaccounts.NewChartOfAccountsService()
	httpservice := &http.HttpService{}
	httpservice.InitializeRouter(coaService)
	httpservice.Start(":8002")
}