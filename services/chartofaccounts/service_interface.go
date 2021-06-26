package chartofaccounts

import (
	"munshee/internal/generic_event"
	"net/http"
)

//ChartOfAccountsService
type ChartOfAccountsService interface {
	AddAccount(account *Account) *generic_event.GenericEvent
	DeleteAccount(id uint) *generic_event.GenericEvent
	UpdateAccount(id uint, account *Account) *generic_event.GenericEvent
	GetAccount(id uint) *generic_event.GenericEvent
	GetAccounts() *generic_event.GenericEvent
	GetAccountByCode(code string) *generic_event.GenericEvent
	InitializeDBSchema() *generic_event.GenericEvent
	AddMainHeadAccounts() *generic_event.GenericEvent
	Validate(account *Account , is_new_record bool) *generic_event.GenericEvent
	HttpGetAll(w http.ResponseWriter, r *http.Request)
	HttpCreate(w http.ResponseWriter, r *http.Request)
	HttpUpdate(w http.ResponseWriter, r *http.Request)
	HttpDelete(w http.ResponseWriter, r *http.Request)
}


