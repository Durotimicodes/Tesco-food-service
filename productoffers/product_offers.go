package productoffers

import (
	"github.com/moretonb/moj-durotimicodes-challenge/models"
)

// product type specification
type ByOneGetOneSpecification struct {
	Product []models.Product
}

// get one product is satisfied
func (b *ByOneGetOneSpecification) IsSatisfied(productes *models.Product) bool {

	var productNames string

	//get Product name
	for _, v := range b.Product {
		productNames += GetProductNameByCode(v.ProductCode)
		break
	}
	//if product contains in the list of special offers return satisfied else return false
	if _, exist := SpecialOffers[productNames]; exist {
		return true
	}

	return false
}

// product discount specification
type ProductDiscountSpecification struct {
	Product []models.Product
}

// discount on product is satisfied
func (d *ProductDiscountSpecification) IsSatisfied(productes *models.Product) bool {

	var productNames string
	var productPrice float64

	//get Product name and price
	for _, v := range d.Product {
		productNames += GetProductNameByCode(v.ProductCode)
		productPrice += GetProductPriceByCode(v.ProductCode)
		break
	}

	//if product contains in the list of special offers return satisfied else return false
	if _, exist := SpecialOffers[productNames]; exist {
		return true
	}
	if _, exist := ProductNameAndPrice[productPrice]; exist {
		return true
	}

	return false
}

// specification settings for extension
type SpecificationSettings struct{}

// CheckSpecifications checks if product matches the spec, if yes it implements the specification else returns the real price
func (s *SpecificationSettings) CheckSpecifications(product []models.Product, spec models.OfferSpecification) []models.Product {
	result := make([]models.Product, 0)

	//get Product name and Product Price
	for i, v := range product {
		if spec.IsSatisfied(&v) {
			//if spec is satisfied implement the Buy-One-Get-One function
			xtraProduct := BuyOneGetOneFree(v.ProductQuantity, v.ProductCode, v.ProductPrice)
			result = append(result, xtraProduct)
		}
		if spec.IsSatisfied(&v) {
			//if spec is satisfied implement the Discount Product function
			discounProduct := DiscountOnAProduct(v.ProductQuantity, v.ProductCode, v.ProductPrice)
			result = append(result, discounProduct)
		} else {
			//else just return the product list
			result = append(result, product[i])
		}
	}
	return result

}
