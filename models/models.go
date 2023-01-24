package models

// custom data types
type ProductCode map[string]string

type ProductAndPrice map[string]float64

// a struct product type
type Product struct {
	ProductId        uint
	ProductName      string
	ProductCode      string
	ProductPrice     float64
	ProductQuantity  uint
	ProductDiscount  float32
	HaveDiscount     bool
	HaveSpecialOffer bool
}

// a checkout bill type
type CheckoutBill struct {
	TypesOfItemScanned   uint
	PurchasedProducts    []Product
	TotalQuantityScanned uint
	TotalPrice           float64
}

// CheckBuyOneGetOneFree settings
type CheckBuyOneGetOneFree struct{}

// offer specification interface
type OfferSpecification interface {
	IsSatisfied(productes *Product) bool
}
