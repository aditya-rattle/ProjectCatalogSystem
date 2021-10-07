package main

import (
	"fmt"
	"github.com/aditya/ProjectCatalog/api"
	"github.com/aditya/ProjectCatalog/models"
	"github.com/jinzhu/gorm"

	_ "github.com/lib/pq"
)


func main() {
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
	db1.AutoMigrate(&models.Transaction{})

	api.Routes(db1)
	defer db1.Close()

}
