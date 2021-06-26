package chartofaccounts

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"munshee/internal/generic_event"
	"munshee/services/database"
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

const LOG_HEAD  = "CHART_OF_ACCOUNTS_SERVICE"

type COAService struct{
	Db *gorm.DB
}


///////////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////// Public Functions //////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////////////////////////////////////

func NewChartOfAccountsService() *COAService{
	log.WithField("Service" , "ChartOfAccounts").Info("Creating new ChartOfAccount Service")
	serv := new(COAService)
	serv.Db = database.InitializeDB()
	return serv
}

func (service *COAService) AddAccount(account *Account) *generic_event.GenericEvent {
	log.WithField("Service" , "ChartOfAccounts").Info("AddAccount Called")
	validationResponse := service.Validate(account , true)
	if(validationResponse.Status == generic_event.INVALID_DATA){
		log.WithField("ValidationError" , validationResponse.Error).Info("While Creating Account")
		return validationResponse
	} else if(validationResponse.Status == generic_event.SUCCESSFUL){
		// Adding Account into the database
		err := service.Db.Create(account).Error
		if err != nil{
			log.WithField("DatabaseError" , err).Info("While Creating Account")
			return generic_event.GenericEventError(err)
		}
		return generic_event.GenericEventSuccess(account)
	} else {
		log.WithField("Unknown Error" , validationResponse).Info("While Creating Account")
		return generic_event.GenericEventError(errors.New("Unknown Error while creating the account"))
	}
}

func (service *COAService) DeleteAccount(id uint) *generic_event.GenericEvent {
	log.WithField("Service" , "ChartOfAccounts").Info("DeleteAccount Called")
	serviceResponse := service.GetAccount(id)
	if serviceResponse.Status == generic_event.FAILED{
		return generic_event.GenericEventError(errors.New("Unable to find the record you are looking for"))
	}
	if err := service.Db.Delete(serviceResponse.Payload.(Account)).Error; err != nil {
		return generic_event.GenericEventError(err)
	}
	return generic_event.GenericEventSuccess(true)
	panic("implement me")
}

func (service *COAService) UpdateAccount(id uint, account *Account) *generic_event.GenericEvent {
	log.WithField("Service" , "ChartOfAccounts").Info("UpdateAccount Called")
	validationResponse := service.Validate(account , false)
	if(validationResponse.Status == generic_event.INVALID_DATA){
		log.WithField("ValidationError" , validationResponse.Error).Info("While Updating Account")
		return validationResponse
	} else if(validationResponse.Status == generic_event.SUCCESSFUL){
		// Adding Account into the database
		err := service.Db.Save(account).Error
		if err != nil{
			log.WithField("DatabaseError" , err).Info("While Updating Account")
			return generic_event.GenericEventError(err)
		}
		return generic_event.GenericEventSuccess(account)
	} else {
		log.WithField("Unknown Error" , validationResponse).Info("While Updating Account")
		return generic_event.GenericEventError(errors.New("Unknown Error while Updating the account"))
	}
	panic("implement me")
}

func (service *COAService) GetAccount(id uint) *generic_event.GenericEvent {
	log.WithField("Service" , "ChartOfAccounts").Info("GetAccount Called")
	var coa Account
	err := service.Db.Preload("Children").First(&coa, "id = ?", id).Error
	if err != nil{
		log.WithField("Service" , "ChartOfAccounts").Error("Got Error while quering for Account Id ")
		log.WithField("Service" , "ChartOfAccounts").Error(err)
		return generic_event.GenericEventError(err)
	}
	return generic_event.GenericEventSuccess(coa)
}

func (service *COAService) GetAccounts() *generic_event.GenericEvent {
	log.WithField("Service" , "ChartOfAccounts").Info("GetAccount Called")
	var coas []Account
	err := service.Db.Preload("Children").Find(&coas).Error
	if err != nil{
		log.WithField("Service" , "ChartOfAccounts").Error("Got Error while quering for Accounts ")
		log.WithField("Service" , "ChartOfAccounts").Error(err)
		return generic_event.GenericEventError(err)
	}
	return generic_event.GenericEventSuccess(coas)
}

func (service *COAService) GetChilden(parentId uint) *generic_event.GenericEvent {
	log.WithField("Service" , "ChartOfAccounts").Info("GetAccount Called")
	var coas []Account
	err := service.Db.Preload("Children").Find(&coas , "parent_id = ?" , parentId).Error
	if err != nil{
		log.WithField("Service" , "ChartOfAccounts").Error("Got Error while quering for Accounts ")
		log.WithField("Service" , "ChartOfAccounts").Error(err)
		return generic_event.GenericEventError(err)
	}
	return generic_event.GenericEventSuccess(coas)
}

func (service *COAService) GetAllByCode(code string) *generic_event.GenericEvent {
	log.WithField("Service" , "ChartOfAccounts").Info("GetAccount Called")
	var coas []Account
	err := service.Db.Preload("Children").Find(&coas , "code = ?" , code).Error
	if err != nil{
		log.WithField("Service" , "ChartOfAccounts").Error("Got Error while quering for Accounts ")
		log.WithField("Service" , "ChartOfAccounts").Error(err)
		return generic_event.GenericEventError(err)
	}
	return generic_event.GenericEventSuccess(coas)
}

func (service *COAService) GetAccountByCode(code string) *generic_event.GenericEvent {
	log.WithField("Service" , "ChartOfAccounts").Info("GetAccountByCode Called")
	var coa Account
	err := service.Db.First(&coa, "full_code = ?", code).Error
	if err != nil{
		log.WithField("Service" , "ChartOfAccounts").Error("Got Error while quering for Account Code ")
		log.WithField("Service" , "ChartOfAccounts").Error(err)
		if err == gorm.ErrRecordNotFound{
			return generic_event.GenericEventError(err)
		}
	}
	return generic_event.GenericEventSuccess(coa)
}

func (service *COAService) InitializeDBSchema() *generic_event.GenericEvent{
	log.WithField("Service" , "ChartOfAccounts").Info("Creating the DB Schema")
	service.Db.AutoMigrate(&Account{})
	service.Db.AutoMigrate(&AccountBalances{})
	service.Db.AutoMigrate(&AccountBalanceHistory{})
	return generic_event.GenericEventSuccess(true)
}

func (service *COAService) AddMainHeadAccounts() *generic_event.GenericEvent{
	log.WithField("Service" , "ChartOfAccounts").Info("AddMainHeadAccounts Called")
	mainaccounts := getMainHeadAccounts()
	for _, mainaccount := range mainaccounts {
		if (service.GetAccountByCode(mainaccount.Code).Status == generic_event.SUCCESSFUL) {
			log.WithField("Service" , "ChartOfAccounts").Info("MainAccount of " + mainaccount.Code + " " + mainaccount.Name + " Already exists")
		} else {
			log.WithField("Service" , "ChartOfAccounts").Info("Creating MainAccount of " + mainaccount.Code + " " + mainaccount.Name)
			service.AddAccount(mainaccount)
		}
	}
	return generic_event.GenericEventSuccess(true)
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////// Private Functions //////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (service *COAService) getAccountInitialCode(parentId uint) string{
	parentAccount := service.GetAccount(parentId).Payload.(Account)
	if(parentAccount.ParentId != nil){
		return service.getAccountInitialCode(*parentAccount.ParentId) + parentAccount.Code
	} else {
		return parentAccount.Code
	}
}

func getMainHeadAccounts() [6]*Account{
	log.WithField("Service" , "ChartOfAccounts").Info("Generating the main heads")
	var mainAccounts = [6]*Account{}
	mainAccounts[0] = &Account{
		Code:     "10",
		FullCode: "10",
		Name:     "Assets",
		Children: nil,
		ParentId: nil,
	}
	mainAccounts[1] = &Account{
		Code:     "20",
		FullCode: "20",
		Name:     "Liabilities",
		Children: nil,
		ParentId: nil,
	}
	mainAccounts[2] = &Account{
		Code:     "30",
		FullCode: "30",
		Name:     "Owner's Capital",
		Children: nil,
		ParentId: nil,
	}
	mainAccounts[3] = &Account{
		Code:     "40",
		FullCode: "40",
		Name:     "Revenues",
		Children: nil,
		ParentId: nil,
	}
	mainAccounts[4] = &Account{
		Code:     "50",
		FullCode: "50",
		Name:     "Expenses",
		Children: nil,
		ParentId: nil,
	}
	mainAccounts[5] = &Account{
		Code:     "60",
		FullCode: "60",
		Name:     "Owners Draw",
		Children: nil,
		ParentId: nil,
	}

	return mainAccounts
}
