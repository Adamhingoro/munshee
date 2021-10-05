package http

import (
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"munshee/services/autho"
	"munshee/services/chartofaccounts"
	"net/http"
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

type HttpService struct {
	Router *mux.Router
	COAService chartofaccounts.ChartOfAccountsService
}

func (service *HttpService) InitializeRouter(COAService chartofaccounts.ChartOfAccountsService) {
	log.Info("Initializing the router")
	service.Router = mux.NewRouter()
	service.Router.Use(loggingMiddleware)
	service.COAService = COAService
	// Adding routes to the router
	service.Router.HandleFunc("/accounts", service.COAService.HttpGetAll).Methods("GET")
	service.Router.HandleFunc("/account", service.COAService.HttpCreate).Methods("POST")
	service.Router.HandleFunc("/account/{accountId}", service.COAService.HttpUpdate).Methods("PUT")
	autho.SetOAuthRoutes(service.Router)
}

func (service *HttpService) Start(listenAddress string){
	log.WithField("ListenAddress" , listenAddress).Info("Starting to listen")
	log.Fatal(http.ListenAndServe(listenAddress , service.Router))
}





