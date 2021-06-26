package main

import (
	log "github.com/sirupsen/logrus"
	"munshee/internal/generic_event"
	"munshee/services/chartofaccounts"
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
	log.Info("Munshee Setup Starting")

	log.Info("Initializing the ChartOfAccounts Database")
	coaService := chartofaccounts.NewChartOfAccountsService()
	if(coaService.InitializeDBSchema().Status == generic_event.SUCCESSFUL){
		log.Info("DB Schema Initialized")
	} else {
		log.Error("Error while creating Chart Of Accounts Table")
	}

	log.Info("Adding Main Head Accounts")
	if(coaService.AddMainHeadAccounts().Status == generic_event.SUCCESSFUL){
		log.Info("Main Head Accounts Added Complete")
	} else {
		log.Error("Error while adding main head accounts")
	}

	log.Info("Setup Complete")
}
