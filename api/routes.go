package api

import (
	"github.com/aditya/ProjectCatalog/implementation"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func Routes(db1 *gorm.DB){
	r := mux.NewRouter()
	catalogService := &implementation.DbImplementation{Db: db1,Dbsales:db1}
	catalogController := CatalogController{CatalogService: catalogService}
	r.HandleFunc("/api/create/product", catalogController.CreateProduct).Methods("POST")
	r.HandleFunc("/api/show/product", catalogController.ShowProducts).Methods("GET")
	r.HandleFunc("/api/show/product/{id}", catalogController.ShowProductById).Methods("GET")
	r.HandleFunc("/api/update/product/{id}", catalogController.UpdateProduct).Methods("PUT")
	r.HandleFunc("/api/buy/product/{id}", catalogController.BuyProduct).Methods("PUT")
	r.HandleFunc("/api/delete/product/{id}", catalogController.DeleteProduct).Methods("DELETE")
	r.HandleFunc("/api/top5/product", catalogController.TopProduct).Methods("GET")


	handler := cors.Default().Handler(r)

	log.Fatal(http.ListenAndServe(":8080", handler))
}


