package productoffers

import (
	"fmt"
	"log"

	"github.com/moretonb/moj-durotimicodes-challenge/models"
)

// checkout can scan items
func CheckOutProducts(products []models.Product) (models.CheckoutBill, float64, uint) {

	var numOfItemScanned uint
	var totalPrice float64

	//scan items
	for i, _ := range products {
		if len(products) > 1 {
			numOfItemScanned = 1
			i++
			numOfItemScanned = uint(i)
		}
	}

	totalPrice, err := GetTotalPriceOfProduct(products)
	if err != nil {
		log.Fatalf("%v Error is calculating total Price:", err)
	}

	totalScannedItems, err := GetTotalQuantitySoldProducts(products)
	if err != nil {
		log.Fatalf("%v Error is calculating total quantity:", err)
	}

	checkOut := models.CheckoutBill{
		TypesOfItemScanned:   numOfItemScanned,
		PurchasedProducts:    products,
		TotalQuantityScanned: totalScannedItems,
		TotalPrice:           totalPrice,
	}

	fmt.Printf("THANK YOU FOR SHOPPING AT THE COOP ☺️\n")
	fmt.Println("-----Receipt-----")
	fmt.Printf("Types of Item Product Scanned : %d\nProduct Purchased : %v\nTotal Products Sold : %d\nTotal Price : %v Pounds\n ", numOfItemScanned, products, totalScannedItems, totalPrice)

	return checkOut, totalPrice, totalScannedItems
}
