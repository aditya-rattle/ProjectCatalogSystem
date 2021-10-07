package implementation

import (
	"errors"
	"fmt"
	"github.com/aditya/ProjectCatalog/models"
	"github.com/jinzhu/gorm"
	"time"
)

type DbImplementation struct {
	Db *gorm.DB
	Dbsales *gorm.DB
}

// implementation of create product function

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

// implementation of show products function

func (h *DbImplementation) ShowProducts() []models.DbProduct {
	var allProd []models.DbProduct

	h.Db.Find(&allProd)
	return allProd
}

// implementation of show product by id function

func (h *DbImplementation) ShowProductById(productId int64) models.DbProduct {
	var product models.DbProduct
	h.Db.First(&product, productId)
	return product
}

// implementation of buy product function

func (h *DbImplementation) BuyProduct(productId, productQuantity int64) error {
	var product models.DbProduct
	h.Db.First(&product,productId)
	if product.Name==""{
		return errors.New("item not found")
	}

	if product.Quantity>=productQuantity {
		product.Quantity = product.Quantity - productQuantity

		productSales:=models.Transaction{
			Name:product.Name,
			Quantity:productQuantity,
			OrderTime: time.Now(),
		}

		h.Dbsales.NewRecord(productSales)
		h.Dbsales.Create(&productSales)

		h.Db.Save(&product)

		return nil
	}


	return errors.New("item is  not sufficient to fulfil request")

}




// implementation of update product function

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


		if updatedProduct.Price!=0{
			product.Price=updatedProduct.Price
		}
		if updatedProduct.Quantity!=0{
			initialQuantity:=product.Quantity
			product.Quantity = initialQuantity + updatedProduct.Quantity
		}

		h.Db.Save(&product)
		return nil


}
// implementation of delete product function


func (h *DbImplementation) DeleteProduct(productId int64) error {
	var product models.DbProduct
	h.Db.First(&product,productId)
	if product.Name==""{
		return errors.New("some error")
	}

	h.Db.Delete(&product)
	return nil
}

// implementation of top 5 products in last hour  function

func (h *DbImplementation) TopProduct() []models.Transaction {
	var BestSellers []models.Transaction
	sqlStr := "SELECT name, SUM(quantity) AS Total FROM (SELECT name, quantity FROM transactions WHERE order_time > NOW() - INTERVAL '1 hour') Temp GROUP BY name ORDER BY Total DESC LIMIT 5"
	rows, err := h.Dbsales.Raw(sqlStr).Rows()

	if err != nil {
		fmt.Println("hello")
		return BestSellers
	}

	for rows.Next() {
		var p models.Transaction
		var name string
		var quantity int64
		_ = rows.Scan(&name, &quantity)
		p.Name = name
		p.Quantity=quantity
		BestSellers = append(BestSellers, p)
	}
	return BestSellers

}

