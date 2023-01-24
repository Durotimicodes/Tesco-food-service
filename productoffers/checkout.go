package productoffers

import (
	"fmt"
	"log"

	"github.com/moretonb/moj-durotimicodes-challenge/models"
)

/*checkout can scan items
 */

func CheckOutProducts(products []models.Product) models.CheckoutBill {

	//a checkout bill
	// var checkoutBill models.CheckoutBill
	var id uint
	var totalPrice float64

	//scan items
	for i, _ := range products {
		if len(products) > 1 {
			id = uint(i)
		}
	}

	totalPrice, err := GetTotalPriceOfProduct(products)
	if err != nil {
		log.Fatalf("%v Error is calculating total Price:", err)
	}

	checkOut := models.CheckoutBill{
		OrderId:           id,
		PurchasedProducts: products,
		TotalPrice:        totalPrice,
	}

	fmt.Println("CHECKOUT", checkOut)

	return checkOut
}
