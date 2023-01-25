package main

import (
	"github.com/moretonb/moj-durotimicodes-challenge/models"
	"github.com/moretonb/moj-durotimicodes-challenge/productoffers"
)

func main() {

	//Product one
	p1 := models.Product{
		ProductId:        1,
		ProductCode:      "FR1",
		ProductQuantity:  10,
		ProductPrice:     3.11,
		ProductDiscount:  0,
		HaveDiscount:     false,
		HaveSpecialOffer: true,
	}

	//Product two
	p3 := models.Product{
		ProductId:        1,
		ProductCode:      "CF1",
		ProductQuantity:  10,
		ProductPrice:     11.23,
		ProductDiscount:  0,
		HaveDiscount:     true,
		HaveSpecialOffer: false,
	}

	//Product three
	p2 := models.Product{
		ProductId:        1,
		ProductCode:      "SR1",
		ProductQuantity:  3,
		ProductPrice:     5.00,
		ProductDiscount:  0,
		HaveDiscount:     true,
		HaveSpecialOffer: false,
	}

	//a group/slice of products
	prod := []models.Product{p1, p2, p3}

	//specification checker
	// getOneFreeSpec := productoffers.ByOneGetOneSpecification{Product: prod}
	discountSpec := productoffers.ProductDiscountSpecification{Product: prod}

	//product spec
	ps := productoffers.SpecificationSettings{}
	var listOfProducts []models.Product

	//iterate products over check specfication 
	for _, v := range ps.CheckSpecifications(prod, &discountSpec) {
		listOfProducts = append(listOfProducts, v)
	}

	//implementing checkout
	productoffers.CheckOutProducts(listOfProducts)
}
