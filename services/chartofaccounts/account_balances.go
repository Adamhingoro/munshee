package chartofaccounts

import "munshee/internal/base_model"

type AccountBalances struct {
	base_model.BaseModel
	Balance    int64   `validate:"required,numeric`
	AccountID int
	Account   Account  `json:"-"`
}

type AccountBalanceHistory struct {
	base_model.BaseModel
	PreviousBalance    int64   `validate:"required,numeric`
	NewBalance    int64   `validate:"required,numeric`
	AccountID int
	Account  Account  `json:"-"`
}