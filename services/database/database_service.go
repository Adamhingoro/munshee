package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)
// This is just a database service which returns the GormDB which will be used in the other services

const LOG_HEAD  = "DATABASE_SERVICE"

var (
	instance *gorm.DB
)
var once sync.Once

func InitializeDB() *gorm.DB{
	log("Requesting DB Connection");
	once.Do(func() {
		log("Connecting to the Database Once");
		// TODO Pass the database credentials from the top-level class as the Db configuration
		dsn := "root:qw4hddqcrg@tcp(127.0.0.1:3306)/munshee?charset=utf8mb4&parseTime=True&loc=Local"
		log("Using DSN " + dsn)
		db , err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log("Error while connecting to the database")
			log(err)
		} else {
			log("Connection Successful")
		}
		instance = db
	})
	log("Connection Returned");
	return instance
}

func log(log ...interface{}){
	fmt.Print("[" + LOG_HEAD + "]");
	fmt.Println(log);
}
