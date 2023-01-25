package productoffers

import (
	"fmt"
	"log"

	"github.com/moretonb/moj-durotimicodes-challenge/models"
)

// checkout can scan items
func CheckOutProducts(products []models.Product) (models.CheckoutBill, float64, uint) {

	var typeOfItemsScanned uint
	var totalPrice float64

	//scan items
	for i, _ := range products {
		if len(products) > 1 {
			typeOfItemsScanned = 1
			i++
			typeOfItemsScanned = uint(i)
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
		TypesOfItemScanned:   typeOfItemsScanned,
		PurchasedProducts:    products,
		TotalQuantityScanned: totalScannedItems,
		TotalPrice:           totalPrice,
	}

	fmt.Println("THANK YOU FOR SHOPPING AT THE COOP ☺️\n -----Receipt-----")
	fmt.Printf("Types of Item Product Scanned : %d\nList Of Products Purchased : %v\nTotal Products Sold : %d\nTotal Price : %v Pounds\n ", typeOfItemsScanned, products, totalScannedItems, totalPrice)

	return checkOut, totalPrice, totalScannedItems
}
