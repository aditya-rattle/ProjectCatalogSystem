package handlers

import (
	"bytes"
	"errors"
	"github.com/aditya/ProjectCatalog/api"
	mock_services "github.com/aditya/ProjectCatalog/mock/services"
	"github.com/aditya/ProjectCatalog/models"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateProductWhenProductCreated(t *testing.T){
	ctrl:=gomock.NewController(t)
	defer ctrl.Finish()
	//id:=2
	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	inputProd:=models.UserProduct{Name:"Laptop", Price: 200000, Quantity: 100, Description: "Tool"}

	BodyProduct := []byte(`{"ID": 1, "CreatedAt": "2021-10-05T22:31:07.006403+05:30", "UpdatedAt": "2021-10-06T00:06:29.649208+05:30", "DeletedAt": null, "name": "Laptop", "description": "Tool", "price": 200000, "quantity": 100
}`)

	req, _ := http.NewRequest("POST", "/api/create/product", bytes.NewBuffer(BodyProduct))
	MockServices :=mock_services.NewMockICatalogService(ctrl)
	MockServices.EXPECT().CreateProduct(inputProd).Return(nil)
	ProductController:=api.CatalogController{CatalogService: MockServices}

	r.HandleFunc("/api/create/product",ProductController.CreateProduct ).Methods("POST")
	r.ServeHTTP(rr, req)
	StringResponse:=`"Product added successfully"`
	str:=strings.TrimSpace(rr.Body.String())
	assert.Equal(t, StringResponse,str)
	assert.Equal(t, 200,rr.Code)
}

func TestCreateProductWhenProductNotCreated(t *testing.T){
	ctrl:=gomock.NewController(t)
	defer ctrl.Finish()
	//id:=2
	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	inputProd:=models.UserProduct{Name:"Laptop", Price: 200000, Quantity: 100, Description: "Tool"}

	BodyProduct := []byte(`{"ID": 1, "CreatedAt": "2021-10-05T22:31:07.006403+05:30", "UpdatedAt": "2021-10-06T00:06:29.649208+05:30", "DeletedAt": null, "name": "Laptop", "description": "Tool", "price": 200000, "quantity": 100
}`)

	req, _ := http.NewRequest("POST", "/api/create/product", bytes.NewBuffer(BodyProduct))
	MockServices :=mock_services.NewMockICatalogService(ctrl)
	MockServices.EXPECT().CreateProduct(inputProd).Return(errors.New("some error"))
	ProductController:=api.CatalogController{CatalogService: MockServices}

	r.HandleFunc("/api/create/product",ProductController.CreateProduct ).Methods("POST")
	r.ServeHTTP(rr, req)
	StringResponse:=`"Encountered error while adding product"`
	str:=strings.TrimSpace(rr.Body.String())
	assert.Equal(t, StringResponse,str)
	assert.Equal(t, 200,rr.Code)
}

func TestShowProductsWhenCatalogNotEmpty(t *testing.T){
	ctrl:=gomock.NewController(t)
	defer ctrl.Finish()
	//id:=2
	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	//inputProd:=models.UserProduct{Name:"Laptop", Price: 200000, Quantity: 100, Description: "Tool"}

	BodyProduct := []byte(`{"ID": 1, "CreatedAt": "2021-10-05T22:31:07.006403+05:30", "UpdatedAt": "2021-10-06T00:06:29.649208+05:30", "DeletedAt": null, "name": "Laptop", "description": "Tool", "price": 200000, "quantity": 100
}`)

	outputProduct:=[]models.DbProduct{{Name:"Laptop", Price: 200000, Quantity: 100, Description: "Tool"}}

	req, _ := http.NewRequest("GET", "/api/show/product", bytes.NewBuffer(BodyProduct))
	MockServices :=mock_services.NewMockICatalogService(ctrl)
	MockServices.EXPECT().ShowProducts().Return(outputProduct)
	ProductController:=api.CatalogController{CatalogService: MockServices}

	r.HandleFunc("/api/show/product",ProductController.ShowProducts ).Methods("GET")
	r.ServeHTTP(rr, req)
	var StringResponse string
	StringResponse=`[{"ID":0,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"name":"Laptop","price":200000,"quantity":100,"description":"Tool"}]`
	str:=strings.TrimSpace(rr.Body.String())
	assert.Equal(t, StringResponse,str)
	assert.Equal(t, 200,rr.Code)
}

func TestShowProductsWhenCatalogEmpty(t *testing.T){
	ctrl:=gomock.NewController(t)
	defer ctrl.Finish()
	//id:=2
	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	//inputProd:=models.UserProduct{Name:"Laptop", Price: 200000, Quantity: 100, Description: "Tool"}

	BodyProduct := []byte(`{"ID": 1, "CreatedAt": "2021-10-05T22:31:07.006403+05:30", "UpdatedAt": "2021-10-06T00:06:29.649208+05:30", "DeletedAt": null, "name": "Laptop", "description": "Tool", "price": 200000, "quantity": 100
}`)

	outputProduct:=[]models.DbProduct{}

	req, _ := http.NewRequest("GET", "/api/show/product", bytes.NewBuffer(BodyProduct))
	MockServices :=mock_services.NewMockICatalogService(ctrl)
	MockServices.EXPECT().ShowProducts().Return(outputProduct)
	ProductController:=api.CatalogController{CatalogService: MockServices}

	r.HandleFunc("/api/show/product",ProductController.ShowProducts ).Methods("GET")
	r.ServeHTTP(rr, req)
	var StringResponse string
	StringResponse="\"catalog is empty, please add items\"\n[]"
	str:=strings.TrimSpace(rr.Body.String())
	assert.Equal(t, StringResponse,str)
	assert.Equal(t, 200,rr.Code)
}

func TestShowProductsByIDWhenAvailable(t *testing.T){
	ctrl:=gomock.NewController(t)
	defer ctrl.Finish()
	var id int64 =1
	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	//inputProd:=models.UserProduct{Name:"Laptop", Price: 200000, Quantity: 100, Description: "Tool"}

	BodyProduct := []byte(`{"ID": 1, "CreatedAt": "2021-10-05T22:31:07.006403+05:30", "UpdatedAt": "2021-10-06T00:06:29.649208+05:30", "DeletedAt": null, "name": "Laptop", "description": "Tool", "price": 200000, "quantity": 100
}`)

	outputProduct:=models.DbProduct{Name:"Laptop", Price: 200000, Quantity: 100, Description: "Tool"}

	req, _ := http.NewRequest("GET", "/api/show/product/1", bytes.NewBuffer(BodyProduct))
	MockServices :=mock_services.NewMockICatalogService(ctrl)
	MockServices.EXPECT().ShowProductById(id).Return(outputProduct)
	ProductController:=api.CatalogController{CatalogService: MockServices}

	r.HandleFunc("/api/show/product/{id}",ProductController.ShowProductById).Methods("GET")
	r.ServeHTTP(rr, req)
	var StringResponse string
	StringResponse="{\"ID\":0,\"CreatedAt\":\"0001-01-01T00:00:00Z\",\"UpdatedAt\":\"0001-01-01T00:00:00Z\",\"DeletedAt\":null,\"name\":\"Laptop\",\"price\":200000,\"quantity\":100,\"description\":\"Tool\"}"
	str:=strings.TrimSpace(rr.Body.String())
	assert.Equal(t, StringResponse,str)
	assert.Equal(t, 200,rr.Code)
}

func TestShowProductsByIDWhenNotAvailable(t *testing.T){
	ctrl:=gomock.NewController(t)
	defer ctrl.Finish()
	var id int64 =1
	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	//inputProd:=models.UserProduct{Name:"Laptop", Price: 200000, Quantity: 100, Description: "Tool"}

	BodyProduct := []byte(`{"ID": 1, "CreatedAt": "2021-10-05T22:31:07.006403+05:30", "UpdatedAt": "2021-10-06T00:06:29.649208+05:30", "DeletedAt": null, "name": "Laptop", "description": "Tool", "price": 200000, "quantity": 100
}`)

	outputProduct:=models.DbProduct{}

	req, _ := http.NewRequest("GET", "/api/show/product/1", bytes.NewBuffer(BodyProduct))
	MockServices :=mock_services.NewMockICatalogService(ctrl)
	MockServices.EXPECT().ShowProductById(id).Return(outputProduct)
	ProductController:=api.CatalogController{CatalogService: MockServices}

	r.HandleFunc("/api/show/product/{id}",ProductController.ShowProductById).Methods("GET")
	r.ServeHTTP(rr, req)
	var StringResponse string
	StringResponse=`"Item does not exists"`
	str:=strings.TrimSpace(rr.Body.String())
	assert.Equal(t, StringResponse,str)
	assert.Equal(t, 200,rr.Code)
}

func TestUpdateProductsWhenAvailable(t *testing.T){
	ctrl:=gomock.NewController(t)
	defer ctrl.Finish()
	var id int64 =1
	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	inputProd:=models.DbProduct{Name:"Laptop", Price: 200000, Quantity: 100, Description: "Tool"}

	BodyProduct := []byte(`{"name": "Laptop", "description": "Tool", "price": 200000, "quantity": 100}`)


	req, _ := http.NewRequest("PUT", "/api/update/product/1", bytes.NewBuffer(BodyProduct))
	MockServices :=mock_services.NewMockICatalogService(ctrl)
	MockServices.EXPECT().UpdateProduct(inputProd,id).Return(nil)
	ProductController:=api.CatalogController{CatalogService: MockServices}

	r.HandleFunc("/api/update/product/{id}",ProductController.UpdateProduct).Methods("PUT")
	r.ServeHTTP(rr, req)
	var StringResponse string
	StringResponse=`"Product updated successfully"`
	str:=strings.TrimSpace(rr.Body.String())
	assert.Equal(t, StringResponse,str)
	assert.Equal(t, 200,rr.Code)
}

func TestUpdateProductsWhenNotAvailable(t *testing.T){
	ctrl:=gomock.NewController(t)
	defer ctrl.Finish()
	var id int64 =1
	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	inputProd:=models.DbProduct{Name:"Laptop", Price: 200000, Quantity: 100, Description: "Tool"}

	BodyProduct := []byte(`{"name": "Laptop", "description": "Tool", "price": 200000, "quantity": 100}`)


	req, _ := http.NewRequest("PUT", "/api/update/product/1", bytes.NewBuffer(BodyProduct))
	MockServices :=mock_services.NewMockICatalogService(ctrl)
	MockServices.EXPECT().UpdateProduct(inputProd,id).Return(errors.New("some error"))
	ProductController:=api.CatalogController{CatalogService: MockServices}

	r.HandleFunc("/api/update/product/{id}",ProductController.UpdateProduct).Methods("PUT")
	r.ServeHTTP(rr, req)
	var StringResponse string
	StringResponse=`"Product not found"`
	str:=strings.TrimSpace(rr.Body.String())
	assert.Equal(t, StringResponse,str)
	assert.Equal(t, 200,rr.Code)
}


func TestDeleteProductWhenAvailable(t *testing.T){
	ctrl:=gomock.NewController(t)
	defer ctrl.Finish()
	var id int64 =1
	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	//inputProd:=models.DbProduct{Name:"Laptop", Price: 200000, Quantity: 100, Description: "Tool"}

	BodyProduct := []byte(`{"name": "Laptop", "description": "Tool", "price": 200000, "quantity": 100}`)


	req, _ := http.NewRequest("DELETE", "/api/delete/product/1", bytes.NewBuffer(BodyProduct))
	MockServices :=mock_services.NewMockICatalogService(ctrl)
	MockServices.EXPECT().DeleteProduct(id).Return(nil)
	ProductController:=api.CatalogController{CatalogService: MockServices}

	r.HandleFunc("/api/delete/product/{id}",ProductController.DeleteProduct).Methods("DELETE")
	r.ServeHTTP(rr, req)
	var StringResponse string
	StringResponse=`"item deleted successfully"`
	str:=strings.TrimSpace(rr.Body.String())
	assert.Equal(t, StringResponse,str)
	assert.Equal(t, 200,rr.Code)
}


func TestDeleteProductWhenNotAvailable(t *testing.T){
	ctrl:=gomock.NewController(t)
	defer ctrl.Finish()
	var id int64 =1
	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	//inputProd:=models.DbProduct{Name:"Laptop", Price: 200000, Quantity: 100, Description: "Tool"}

	BodyProduct := []byte(`{"name": "Laptop", "description": "Tool", "price": 200000, "quantity": 100}`)


	req, _ := http.NewRequest("DELETE", "/api/delete/product/1", bytes.NewBuffer(BodyProduct))
	MockServices :=mock_services.NewMockICatalogService(ctrl)
	MockServices.EXPECT().DeleteProduct(id).Return(errors.New("some error"))
	ProductController:=api.CatalogController{CatalogService: MockServices}

	r.HandleFunc("/api/delete/product/{id}",ProductController.DeleteProduct).Methods("DELETE")
	r.ServeHTTP(rr, req)
	var StringResponse string
	StringResponse=`"there was an error deleting item"`
	str:=strings.TrimSpace(rr.Body.String())
	assert.Equal(t, StringResponse,str)
	assert.Equal(t, 200,rr.Code)
}

func TestBuyProductWhenAvailable(t *testing.T){
	ctrl:=gomock.NewController(t)
	defer ctrl.Finish()
	var id int64 =1
	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	inputProd:=models.DbProduct{Name:"Laptop", Price: 200000, Quantity: 100, Description: "Tool"}

	BodyProduct := []byte(`{"name": "Laptop", "description": "Tool", "price": 200000, "quantity": 100}`)


	req, _ := http.NewRequest("PUT", "/api/buy/product/1", bytes.NewBuffer(BodyProduct))
	MockServices :=mock_services.NewMockICatalogService(ctrl)
	MockServices.EXPECT().BuyProduct(id,inputProd.Quantity).Return(nil)
	ProductController:=api.CatalogController{CatalogService: MockServices}

	r.HandleFunc("/api/buy/product/{id}",ProductController.BuyProduct).Methods("PUT")
	r.ServeHTTP(rr, req)
	var StringResponse string
	StringResponse=`"Congratulations, your purchase was successful."`
	str:=strings.TrimSpace(rr.Body.String())
	assert.Equal(t, StringResponse,str)
	assert.Equal(t, 200,rr.Code)
}

func TestBuyProductWhenNotAvailable(t *testing.T){
	ctrl:=gomock.NewController(t)
	defer ctrl.Finish()
	var id int64 =1
	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	inputProd:=models.DbProduct{Name:"Laptop", Price: 200000, Quantity: 100, Description: "Tool"}

	BodyProduct := []byte(`{"name": "Laptop", "description": "Tool", "price": 200000, "quantity": 100}`)


	req, _ := http.NewRequest("PUT", "/api/buy/product/1", bytes.NewBuffer(BodyProduct))
	MockServices :=mock_services.NewMockICatalogService(ctrl)
	MockServices.EXPECT().BuyProduct(id,inputProd.Quantity).Return(errors.New("some error"))
	ProductController:=api.CatalogController{CatalogService: MockServices}

	r.HandleFunc("/api/buy/product/{id}",ProductController.BuyProduct).Methods("PUT")
	r.ServeHTTP(rr, req)
	var StringResponse string
	StringResponse=`"Item not available."`
	str:=strings.TrimSpace(rr.Body.String())
	assert.Equal(t, StringResponse,str)
	assert.Equal(t, 200,rr.Code)
}
