package databaseconnection

import (
	"fmt"
	"github.com/aditya/ProjectCatalog/models"
	"github.com/jinzhu/gorm"
)

func Dbconfig() *gorm.DB{
	var db1 *gorm.DB
	var err error
	dbURI := fmt.Sprintf("host=localhost user=postgres dbname=postgres sslmode=disable password=aditya port=5432")

	db1, err = gorm.Open("postgres", dbURI)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connected to database successfully")
	}


	db1.AutoMigrate(&models.DbProduct{})
	//db1.AutoMigrate(&models.Transaction{})
	return db1
}
