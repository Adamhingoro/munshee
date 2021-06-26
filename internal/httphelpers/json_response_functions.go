package httphelpers

import (
	"encoding/json"
	"net/http"
)

type APIResponseCode int

const (
	SUCCESS APIResponseCode = 200
	FOUND APIResponseCode = 302
	BAD_CLIENT_REQEUST APIResponseCode = 400
	UNAUTHORIZED APIResponseCode = 401
	FORBIDDEN APIResponseCode = 403
	NOT_FOUND APIResponseCode = 404
	NOT_ALLOWED APIResponseCode = 405
	INTERNAL_SERVER_ERROR APIResponseCode = 500
	NOT_IMPLEMENTED APIResponseCode = 501
	SERVICE_NOT_AVAILIBLE APIResponseCode = 503
)

func RespondHttpBadClientRequest(w http.ResponseWriter){
	RespondHttpWithError(w , int(BAD_CLIENT_REQEUST) , "Invalid Request. Please check the request and try again.")
}

func RespondHttpInvalidData(w http.ResponseWriter , err error){
	RespondHttpWithError(w , int(BAD_CLIENT_REQEUST) , err.Error())
}

func RespondHttpUnauthorized(w http.ResponseWriter){
	RespondHttpWithError(w , int(UNAUTHORIZED) , "You are not authenticated.")
}

func RespondHttpForbidden(w http.ResponseWriter){
	RespondHttpWithError(w , int(FORBIDDEN) , "You are not authorized to continue this request")
}

func RespondHttpNotFound(w http.ResponseWriter){
	RespondHttpWithError(w , int(NOT_FOUND) , "Unable to find what you're looking for")
}


func RespondHttpWithError(w http.ResponseWriter, code int, message string) {
	RespondHttpWithJSON(w, code, map[string]string{"error": message})
}

func RespondHttpWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
