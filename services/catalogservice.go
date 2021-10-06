package services

import "github.com/aditya/ProjectCatalog/models"

type ICatalogService interface{
	CreateProduct(product models.UserProduct) error
	ShowProducts() []models.DbProduct
	ShowProductById(productId int64) models.DbProduct
	UpdateProduct(models.DbProduct, int64) error
	BuyProduct(productId int64, productQuantity int64) error
	DeleteProduct(productId int64) error
	//TopProduct() []string
}
