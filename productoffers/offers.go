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
	for i := 0; i < len(ProductsOnGetOneFreeOffer); i++ {
		//if product is found on the list of special get-one offers and the prices are similar, add an extra one on each product else just return product price
		if prodName == ProductsOnGetOneFreeOffer[i] && prodPrice == pPrice {
			newProductQuantity = prodQuantity + prodQuantity
			sOffer = true
		} else {
			newProductQuantity += prodQuantity
		}
	}

	//populate the values with product model
	finalProduct := models.Product{
		ProductId:        productID,
		ProductPrice:     pPrice,
		ProductQuantity:  newProductQuantity,
		ProductCode:      prodCode,
		HaveSpecialOffer: sOffer,
		ProductName:      prodName,
	}

	return finalProduct
}

// DiscountOnAProduct takes in a product code, quantity and price and checks if product has a discount offer
func DiscountOnAProduct(prodQuantity uint, prodCode string, prodPrice float64) models.Product {

	var productID uint
	var sOffer bool
	var prodName string
	var newProductPrice, discountPercentage float64

	//get product name from map
	for k, v := range ProductNameAndCode {
		if v == prodCode {
			prodName = k
		}
	}

	//get discount percentage
	for k, v := range DiscountOnProduct {
		if k == prodName {
			discountPercentage = v
		}
	}

	//get product price from function
	pPrice := GetProductPriceByCode(prodCode)

	for i := 0; i < len(ProductsOnDiscountOffer); i++ {
		/*if product is found on the list of discount offers, the prices are similar and ordered quantity purchased is 
		 greater or equal to 3, add a discount on each product else just return product price
		 */
		if prodName == ProductsOnDiscountOffer[i] && prodPrice == pPrice && prodQuantity >= 3 {
			reduction := pPrice * discountPercentage
			newProductPrice = pPrice - reduction
		} else {
			newProductPrice = pPrice
		}

	}

	//populate the values into object product model
	finalProduct := models.Product{
		ProductId:        productID,
		ProductPrice:     newProductPrice,
		ProductQuantity:  prodQuantity,
		ProductCode:      prodCode,
		HaveSpecialOffer: sOffer,
		ProductName:      prodName,
	}

	return finalProduct
}
