package main

import (
	"github.com/aditya/ProjectCatalog/api"
	"github.com/aditya/ProjectCatalog/dbconnection"
	"github.com/aditya/ProjectCatalog/models"
	_ "github.com/lib/pq"
)

func main() {
	db := dbconnection.Dbconfiguration()
	db.AutoMigrate(&models.DbProduct{})
	db.AutoMigrate(&models.Transaction{})
	api.Routes(db)
	defer db.Close()
}
