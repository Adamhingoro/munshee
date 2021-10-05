package autho

import (
	log "github.com/sirupsen/logrus"
	"munshee/internal/httphelpers"
	"net/http"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {

	session, err := Store.Get(r, "auth-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// returning the logged in user info

	log.WithFields(log.Fields{
		"session": session.Values,
	}).Info("User authenticated successfully")
	httphelpers.RespondHttpWithJSON(w , 200, session.Values["profile"])
}
