package implementation

import (
	"errors"
	"fmt"
	"github.com/aditya/ProjectCatalog/models"
	"github.com/jinzhu/gorm"
)

type DbImplementation struct {
	Db *gorm.DB
}

//Add Product
func (h *DbImplementation) CreateProduct(product models.UserProduct) error {
	//handle error
	dbprod:=models.DbProduct{
		Name:product.Name,
		Price: product.Price,
		Quantity: product.Quantity,
		Description:product.Description,
	}
	fmt.Println(dbprod)

	h.Db.Create(&dbprod)

	return nil
}

//Available Product
func (h *DbImplementation) ShowProducts() []models.DbProduct {
	var allProd []models.DbProduct

	h.Db.Find(&allProd)
	return allProd
}

//Get Product By Id
func (h *DbImplementation) ShowProductById(productId int64) models.DbProduct {
	var product models.DbProduct
	h.Db.First(&product, productId)
	return product
}


//Buy Product
func (h *DbImplementation) BuyProduct(productId, productQuantity int64) error {
	var product models.DbProduct
	h.Db.First(&product,productId)
	if product.Name==""{
		return errors.New("item not found")
	}

	if product.Quantity>=productQuantity {
		product.Quantity = product.Quantity - productQuantity
		h.Db.Save(&product)
		return nil
	}

	return errors.New("item is  not sufficient to fulfil request")

}




// Update the Product
func (h *DbImplementation) UpdateProduct( updatedProduct models.DbProduct , productId int64) error {
	var product models.DbProduct
	h.Db.First(&product,productId)

	if product.Name==""{
		return errors.New("product not found")
	}

	if updatedProduct.Quantity<0{
		return errors.New("quantity is invalid")
	}

	if updatedProduct.Price<0{
		return errors.New("price of the product is invalid")
	}

	//if int64(product.ID) == productId {

		//fmt.Println("hi",initialQuantity)
		if updatedProduct.Price!=0{
			product.Price=updatedProduct.Price
		}
		if updatedProduct.Quantity!=0{
			initialQuantity:=product.Quantity
			product.Quantity = initialQuantity + updatedProduct.Quantity
		}
		//fmt.Println("hi",product.Quantity,updatedProduct.Quantity)

		h.Db.Save(&product)
		return nil
	//}

}

func (h *DbImplementation) DeleteProduct(productId int64) error {
	var product models.DbProduct
	h.Db.First(&product,productId)
	if product.Name==""{
		return errors.New("some error")
	}

	h.Db.Delete(&product)

	return nil
}




