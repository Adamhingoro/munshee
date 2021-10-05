package autho

import "github.com/gorilla/mux"

import (
	"encoding/gob"
	"fmt"

	"github.com/gorilla/sessions"
)

var (
	Store *sessions.FilesystemStore
)

func InitializeSessionStore() error {
	fmt.Println("Initializing the Session Store")
	Store = sessions.NewFilesystemStore("", []byte("something-very-secret"))
	gob.Register(map[string]interface{}{})
	return nil
}

func SetOAuthRoutes(router *mux.Router){
	InitializeSessionStore()
	router.HandleFunc("/callback" , CallbackHandler)
	router.HandleFunc("/login" , LoginHandler)
	router.HandleFunc("/user" , UserHandler)
	router.HandleFunc("/logout" , LogoutHandler)
}
