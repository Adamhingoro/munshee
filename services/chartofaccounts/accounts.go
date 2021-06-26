package chartofaccounts

import (
	"errors"
	"munshee/internal/base_model"
	"munshee/internal/generic_event"
)
import (
	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
)

type Account struct {
	base_model.BaseModel
	Code       string    `validate:"required,numeric,min=1",json:"code"`
	FullCode   string
	Name       string    `validate:"required",json:"name"`
	Children   []Account `gorm:"foreignkey:ParentId",json:"-"`
	ParentId   *uint
}

func (service *COAService) Validate(account *Account , is_new_record bool) *generic_event.GenericEvent{
	log.WithField("Account" , account).Info("Validating the Account using go-validator/v10")
	var validate *validator.Validate
	validate = validator.New()

	// Here we handle the Parent Relation and the FullCode algo
	if account.ParentId != nil{
		if service.GetAccount(*account.ParentId).Status == generic_event.FAILED{
			return generic_event.GenericEventInvalidData(errors.New("Unable to find the parent you are trying to assign."))
		}
		account.FullCode = service.getAccountInitialCode(*account.ParentId) + account.Code
	} else {
		account.FullCode = account.Code
	}

	// Check if account already exists with the same code
	if is_new_record {
		if service.GetAccountByCode(account.FullCode).Status == generic_event.SUCCESSFUL{
			return generic_event.GenericEventInvalidData(errors.New("Account with code " + account.FullCode + " already exists. try different code"))
		}
	} else {
		existingAccountWithSameCodeServiceResponse := service.GetAccountByCode(account.FullCode)
		if existingAccountWithSameCodeServiceResponse.Status == generic_event.SUCCESSFUL{
			existingAccountWithSameCode := existingAccountWithSameCodeServiceResponse.Payload.(Account)
			if existingAccountWithSameCode.ID != account.ID{
				return generic_event.GenericEventInvalidData(errors.New("Account with code " + account.Code + " already exists. try different code"))
			}
		}
	}

	err := validate.Struct(account)
	if err != nil{
		// if it is invalid validation rule
		if _, ok := err.(*validator.InvalidValidationError); ok {
			log.WithField("ValidationErrors" , err).Info("Invalid Validation Rules Error")
			return generic_event.GenericEventError(err)
		}

		// if it is actual data validation error
		log.WithField("ValidationErrors" , err).Info("Error validating the accounts")
		return generic_event.GenericEventInvalidData(err.(validator.ValidationErrors))
	} else {
		// if everything is fine and data is good to go
		return generic_event.GenericEventSuccess(true)
	}
}

