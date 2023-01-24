package productoffers

// in-memory-data-store
// key value pair for products and their corrensponding prices
var ProductNameAndPrice = map[float64]string{
	3.11:  "FR1",
	5.00:  "SR1",
	11.23: "CF1",
}

// key value pair for products and their corresponding codes
var ProductNameAndCode = map[string]string{
	"Fruit tea":    "FR1",
	"Strawberries": "SR1",
	"Coffee":       "CF1",
}

// special offer deal setttings
var SpecialOffers = map[string]uint{
	"Fruit tea": 1,
}

// key value pair to store product discounts
var DiscountOnProduct = map[string]float64{
	"Strawberries": 0.10,
}

// list of products on buy-one-get-one-free offer
var ProductsOnGetOneFreeOffer = []string{"Fruit tea"}

// list of products on discount offer
var ProductsOnDiscountOffer = []string{"Strawberries"}
