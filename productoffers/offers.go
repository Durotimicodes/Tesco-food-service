package productoffers

import "github.com/moretonb/moj-durotimicodes-challenge/models"

// BuyOneGetOneFree takes in a product code, quantity and price and checks if product has a special offer
func BuyOneGetOneFree(prodQuantity uint, prodCode string, prodPrice float64) models.Product {
	var newProductQuantity, productID uint
	var sOffer bool
	var prodName string

	//get product name from map
	for k, v := range ProductNameAndCode {
		if v == prodCode {
			prodName = k
		}
	}

	//get product price from function
	pPrice := GetProductPriceByCode(prodCode)

	//checks
	for i := 0; i < len(ProductsOnOffer); i++ {
		//if product is found in the list of offers and the prices are similar add an extra one on each product else just return product price
		if prodName == ProductsOnOffer[i] && prodPrice == pPrice {
			newProductQuantity = prodQuantity + prodQuantity
			sOffer = true
		} else {
			newProductQuantity += prodQuantity
		}
	}

	//populate the values with product model
	prd := models.Product{
		ProductId:        productID,
		ProductQuantity:  newProductQuantity,
		ProductCode:      prodCode,
		HaveSpecialOffer: sOffer,
		ProductName:      prodName,
	}

	return prd
}

// DiscountOnAProduct takes in a product code, quantity and price and checks if product has a discount offer
func DiscountOnAProduct(prodQuantity uint, prodCode string, prodPrice float64) bool {

	return false
}
