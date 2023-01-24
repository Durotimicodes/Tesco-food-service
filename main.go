package main

import (
	"fmt"

	"github.com/moretonb/moj-durotimicodes-challenge/models"
)

func main() {
	p1 := models.Product{
		ProductId:        1,
		ProductCode:      "FR1",
		ProductQuantity:  10,
		ProductPrice:     3.11,
		ProductDiscount:  0,
		HaveDiscount:     false,
		HaveSpecialOffer: false,
	}

	p3 := models.Product{
		ProductId:        1,
		ProductCode:      "CF1",
		ProductQuantity:  10,
		ProductPrice:     11.23,
		ProductDiscount:  0,
		HaveDiscount:     false,
		HaveSpecialOffer: false,
	}

	p2 := models.Product{
		ProductId:        1,
		ProductCode:      "SR1",
		ProductQuantity:  2,
		ProductPrice:     5.00,
		ProductDiscount:  0,
		HaveDiscount:     false,
		HaveSpecialOffer: false,
	}

	prod := []models.Product{p1, p2, p3}
	fmt.Println(prod)
}
