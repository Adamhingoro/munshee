package autho

import (
	"net/http"
	"net/url"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {

	domain := "adamhingoro.eu.auth0.com"

	logoutUrl, err := url.Parse("https://" + domain)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logoutUrl.Path += "/v2/logout"
	parameters := url.Values{}

	var scheme string
	if r.TLS == nil {
		scheme = "http"
	} else {
		scheme = "https"
	}

	returnTo, err := url.Parse(scheme + "://" +  r.Host)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	parameters.Add("returnTo", returnTo.String())
	parameters.Add("client_id", "rPyxM9VIegTVlECGP0AihbO6IZNVHPvB")
	logoutUrl.RawQuery = parameters.Encode()

	http.Redirect(w, r, logoutUrl.String(), http.StatusTemporaryRedirect)
}
