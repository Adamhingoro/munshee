package chartofaccounts

import (
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"munshee/internal/generic_event"
	"munshee/internal/httphelpers"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func init() {
	// TODO:  Add a config for json or text based logs
	//log.SetFormatter(&log.JSONFormatter{})
	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

// HttpGetAll To Get All the Accounts from the database
func (service *COAService) HttpGetAll(w http.ResponseWriter, r *http.Request) {
	// TODO Will use header.get for setting and getting the request details like user, scope and roles rights
	name := r.Header.Get("Sample")
	log.Info(name)


	log.WithField("HttpHandler" , "ChartOfAccounts").Info("HttpGetAll Called")
	count, _ := strconv.Atoi(r.FormValue("count"))
	start, _ := strconv.Atoi(r.FormValue("start"))

	if count > 10 || count < 1 {
		count = 10
	}
	if start < 0 {
		start = 0
	}

	serviceResponse := service.GetAccounts()
	if(serviceResponse.Status == generic_event.SUCCESSFUL){
		httphelpers.RespondHttpWithJSON(w, http.StatusOK, serviceResponse.Payload)
	} else {
		httphelpers.RespondHttpBadClientRequest(w)
	}

}

func (service *COAService) HttpCreate(w http.ResponseWriter, r *http.Request) {
	log.WithField("HttpHandler" , "ChartOfAccounts").Info("HttpCreate Called")

	reqBody, _ := ioutil.ReadAll(r.Body)
	var account Account
	json.Unmarshal(reqBody , &account)

	// Removing any leading zeros
	account.Code = strings.TrimLeft(account.Code , "0")

	serviceResponse := service.AddAccount(&account)
	if(serviceResponse.Status == generic_event.SUCCESSFUL){
		httphelpers.RespondHttpWithJSON(w, http.StatusOK, serviceResponse.Payload)
	} else if(serviceResponse.Status == generic_event.INVALID_DATA) {
		httphelpers.RespondHttpInvalidData(w , serviceResponse.Error)
	} else {
		httphelpers.RespondHttpBadClientRequest(w)
	}

}

func (service *COAService) HttpUpdate(w http.ResponseWriter, r *http.Request) {
	log.WithField("HttpHandler" , "ChartOfAccounts").Info("HttpCreate Called")

	// getting the variable from the URL
	vars := mux.Vars(r)
	accountId, err := strconv.ParseUint(vars["accountId"],10,32)
	if err != nil{
		httphelpers.RespondHttpBadClientRequest(w)
		return
	}


	// Checking if the account exists or not
	serviceResponse := service.GetAccount(uint(accountId))
	if(serviceResponse.Status == generic_event.FAILED){
		httphelpers.RespondHttpNotFound(w)
		return
	}

	// Casting the Interface to the Account
	existingAccount , ok := serviceResponse.Payload.(Account)
	if !ok {
		httphelpers.RespondHttpBadClientRequest(w)
		return
	}

	// Loading the new values form the json body
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody , &existingAccount)

	// Updating the account with the given values
	serviceResponse = service.UpdateAccount(uint(accountId), &existingAccount)
	if(serviceResponse.Status == generic_event.SUCCESSFUL){
		httphelpers.RespondHttpWithJSON(w, http.StatusOK, serviceResponse.Payload)
	} else if(serviceResponse.Status == generic_event.INVALID_DATA) {
		httphelpers.RespondHttpInvalidData(w , serviceResponse.Error)
	} else {
		httphelpers.RespondHttpBadClientRequest(w)
	}

}

func (service *COAService) HttpDelete(w http.ResponseWriter, r *http.Request) {
	log.WithField("HttpHandler" , "ChartOfAccounts").Info("HttpDelete Called")

	// getting the variable from the URL
	vars := mux.Vars(r)
	accountId, err := strconv.ParseUint(vars["accountId"],10,32)
	if err != nil{
		httphelpers.RespondHttpBadClientRequest(w)
		return
	}

	// Checking if the account exists or not
	serviceResponse := service.GetAccount(uint(accountId))
	if(serviceResponse.Status == generic_event.FAILED){
		httphelpers.RespondHttpNotFound(w)
		return
	}

	// Updating the account with the given values
	serviceResponse = service.DeleteAccount(uint(accountId))
	if(serviceResponse.Status == generic_event.SUCCESSFUL){
		httphelpers.RespondHttpWithJSON(w, http.StatusOK, serviceResponse.Payload)
	} else if(serviceResponse.Status == generic_event.INVALID_DATA) {
		httphelpers.RespondHttpInvalidData(w , serviceResponse.Error)
	} else {
		httphelpers.RespondHttpBadClientRequest(w)
	}

}