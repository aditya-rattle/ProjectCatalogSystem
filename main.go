package main

import (
	"github.com/aditya/ProjectCatalog/api"
	"github.com/aditya/ProjectCatalog/databaseconnection"

	_ "github.com/lib/pq"
)


func main() {
	db1:=databaseconnection.Dbconfig()
	api.Routes(db1)
	defer db1.Close()

}
