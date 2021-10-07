package main

import (
	"github.com/aditya/ProjectCatalog/api"
	"github.com/aditya/ProjectCatalog/dbconnection"
	"github.com/aditya/ProjectCatalog/models"
	_ "github.com/lib/pq"
)


func main() {
	db1:=dbconnection.Dbconfiguration()
	db1.AutoMigrate(&models.DbProduct{})
	db1.AutoMigrate(&models.Transaction{})
	api.Routes(db1)
	defer db1.Close()
}
