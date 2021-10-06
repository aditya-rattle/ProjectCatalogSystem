package api

import (
	"encoding/json"

	"github.com/aditya/ProjectCatalog/models"
	"github.com/aditya/ProjectCatalog/services"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)
type Ihandlers interface {
	CreateProduct(w http.ResponseWriter, r *http.Request)
	ShowProducts(w http.ResponseWriter, r *http.Request)
	ShowProductById(w http.ResponseWriter, r *http.Request)
	UpdateProduct(w http.ResponseWriter, r *http.Request)
	BuyProduct(w http.ResponseWriter, r *http.Request)
	DeleteProduct(w http.ResponseWriter, r *http.Request)
}

type CatalogController struct{
	CatalogService services.ICatalogService
}
func headerJson(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type", "application/json")
}
func (h *CatalogController) CreateProduct(w http.ResponseWriter, r *http.Request){
	headerJson(w,r)
	var newProduct models.UserProduct
	_ = json.NewDecoder(r.Body).Decode(&newProduct)
	err := h.CatalogService.CreateProduct(newProduct)
	if err==nil{
		json.NewEncoder(w).Encode("Product added successfully")
		return
	}

	json.NewEncoder(w).Encode("Encountered error while adding product")
}

func (h *CatalogController) ShowProducts(w http.ResponseWriter,r *http.Request){
	headerJson(w,r)
	product:=h.CatalogService.ShowProducts()
	if len(product)==0{
		json.NewEncoder(w).Encode("catalog is empty, please add items")

	}
	json.NewEncoder(w).Encode(product)
}

func (h *CatalogController) ShowProductById(w http.ResponseWriter, r *http.Request){
	headerJson(w,r)
	params:=mux.Vars(r)
	productId,_:=strconv.ParseInt(params["id"],10,0)
	product:=h.CatalogService.ShowProductById(productId)

	emptyProduct:=models.DbProduct{}
	if product==emptyProduct{
		json.NewEncoder(w).Encode("Item does not exists")
		return
	}
	json.NewEncoder(w).Encode(product)
}

func (h *CatalogController) UpdateProduct(w http.ResponseWriter, r *http.Request){
	headerJson(w,r)
	params:=mux.Vars(r)
	var updatedProduct models.DbProduct
	_ = json.NewDecoder(r.Body).Decode(&updatedProduct)
	productId,_:=strconv.ParseInt(params["id"],10,0)
	err :=h.CatalogService.UpdateProduct(updatedProduct,productId)
	if err==nil{
		json.NewEncoder(w).Encode("Product updated successfully")
		return
	}
	json.NewEncoder(w).Encode("Product not found")

}

func (h *CatalogController) BuyProduct(w http.ResponseWriter, r *http.Request){
	headerJson(w,r)
	params:=mux.Vars(r)
	var purchaseProduct models.DbProduct
	_ = json.NewDecoder(r.Body).Decode(&purchaseProduct)

	productId,_:=strconv.ParseInt(params["id"],10,0)
	productQuantity:=purchaseProduct.Quantity

	err:=h.CatalogService.BuyProduct(productId, productQuantity)

	if err==nil{
		json.NewEncoder(w).Encode("Congratulations, your purchase was successful.")
		return
	}

	json.NewEncoder(w).Encode("Item not available.")

}

func (h *CatalogController) DeleteProduct(w http.ResponseWriter, r *http.Request){
	headerJson(w,r)
	params:=mux.Vars(r)
	productId,_:=strconv.ParseInt(params["id"],10,0)
	err:=h.CatalogService.DeleteProduct(productId)
	if err==nil{
		json.NewEncoder(w).Encode("item deleted successfully")
		return
	}
	json.NewEncoder(w).Encode("there was an error deleting item")


}
//func (h *CatalogController) TopProduct(w http.ResponseWriter, r *http.Request){
//	headerJson(w,r)
//
//	listOfProducts:=h.CatalogService.TopProduct()
//
//	if len(listOfProducts)==0{
//		json.NewEncoder(w).Encode("There are no purchases in the last one hour.")
//		return
//	}
//
//	json.NewEncoder(w).Encode(listOfProducts)
//
//}