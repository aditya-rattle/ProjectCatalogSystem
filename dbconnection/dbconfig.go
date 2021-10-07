package dbconnection

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

func Dbconfiguration() *gorm.DB {
	var db *gorm.DB
	var err error
	dbURI := fmt.Sprintf("host=localhost user=postgres dbname=postgres sslmode=disable password=aditya port=5432")

	db, err = gorm.Open("postgres", dbURI)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connected to database successfully")
	}

	return db
}
