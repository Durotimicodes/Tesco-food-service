package productoffers

// GetProductNameByCode gets the product name by the product code
func GetProductNameByCode(productCode string) string {
	for k, v := range ProductNameAndCode {
		//if you can find product code return the product name else return out of stock
		if v == productCode {
			return k
		}
	}
	return "Product out of stock"
}

// GetProductPriceByCode gets the product price by the product code
func GetProductPriceByCode(productCode string) float64 {
	for k, v := range ProductNameAndPrice {
		//if product is found return the product price else return zero
		if v == productCode {
			return k
		}
	}
	return 0
}

// GetProductionBonusQuantityByCode get the additional bonus quantity added to product
func GetProductBonusQuanityByCode(productCode string) uint {
	for k, v := range SpecialOffers {
		if k == productCode {
			return v
		}
	}
	return 0
}

