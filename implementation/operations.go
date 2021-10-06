package implementation

import (
	"errors"
	"github.com/aditya/ProjectCatalog/models"
	"sync"
	"time"
)

type TimeStampStructure struct{
	prodId int64
	timeInSec int64
}

type Inmemoryimplement struct{
	ProductArr []models.InMemoryProduct
	mutex sync.Mutex
}
var mp map[TimeStampStructure]int64



func Initializer(){
	mp =make(map[TimeStampStructure]int64)
	return
}

func min(a int64, b int64) int64{
	if a<=b{
		return a
	}

	return b

}

func (h *Inmemoryimplement) CreateProduct(product models.InMemoryProduct) error{
	h.mutex.Lock()

	for _, prod:=range h.ProductArr{
		if prod.Name==product.Name{
			return errors.New("product already exist")
		}
	}
	product.Id=int64(len(h.ProductArr)+1)
	h.ProductArr = append(h.ProductArr,product)

	h.mutex.Unlock()
	return nil

}

func (h *Inmemoryimplement) ShowProducts() []models.InMemoryProduct{

	return h.ProductArr
}

func (h *Inmemoryimplement) ShowProductById(productId int64) models.InMemoryProduct{
	h.mutex.Lock()
	for _,product:=range h.ProductArr{
		if product.Id==productId{
			return product
		}
	}
	h.mutex.Unlock()
	return models.InMemoryProduct{}
}

func (h *Inmemoryimplement) UpdateProduct(updatedProduct models.InMemoryProduct,productName string) error{
	h.mutex.Lock()
	for i,item:=range h.ProductArr{
		if item.Name==productName{
			updatedProduct.Id = int64(i+1)

			updatedProduct.Quantity=updatedProduct.Quantity+item.Quantity
			h.ProductArr[i] = updatedProduct
			return nil
		}
	}
	h.mutex.Unlock()
	return errors.New("product do not exist")
}

func (h *Inmemoryimplement) BuyProduct(productId , productQuantity int64) error{
	h.mutex.Lock()
	for i,item:=range h.ProductArr{
		if item.Id==productId{
			currCount:= item.Quantity
			reqCount:= productQuantity
			if currCount-reqCount < 0 {
				return errors.New("Product not available.")
			}

			Idint:=item.Id
			timestamp:=time.Now().Unix()

			local:= TimeStampStructure{
				prodId: Idint,
				timeInSec: timestamp,
			}
			mp[local]=reqCount

			currCount-=reqCount

			h.ProductArr[i].Quantity=currCount
			return nil
		}
	}
	h.mutex.Unlock()
	return errors.New("item not available")
}

//func (h *Inmemoryimplement) TopProduct() []string{
//	currTime:=time.Now().Unix()
//	prodQuantity:=make([]int64,len(h.ProductArr)+1)
//	const timeHour=3600
//	for k,v :=range mp {
//		if currTime-k.timeInSec<=timeHour{
//			prodQuantity[k.prodId]+=v
//		}
//	}
//
//	productsTop:=make([][2]int64,0)
//
//	for ind,val:=range prodQuantity{
//		if val>0 {
//			arr := [2]int64{val, int64(ind)}
//			productsTop = append(productsTop, arr)
//		}
//	}
//
//	sort.Slice(productsTop, func(i, j int) bool {
//		return productsTop[i][0] > productsTop[j][0]
//	})
//
//	listProducts:=make([]string,0)
//	l:=len(productsTop)
//	var i int64
//	for i=0;i< min(5,int64(l));i++{
//		ind:=productsTop[i][1]-1
//		listProducts=append(listProducts,h.ProductArr[ind].Name)
//	}
//
//	return listProducts
//}

