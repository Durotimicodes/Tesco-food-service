package productoffers

import (
	"github.com/moretonb/moj-durotimicodes-challenge/models"
)

// product type specification
type ByOneGetOneSpecification struct {
	Product []models.Product
}

// one product is satisfied
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

type SpecificationSettings struct{}

func (s *SpecificationSettings) CheckSpecifications(product []models.Product, spec models.OfferSpecification) []models.Product {
	result := make([]models.Product, 0)

	//get Product name and Product Price
	for i, v := range product {
		if spec.IsSatisfied(&v) {
			s := BuyOneGetOneFree(v.ProductQuantity, v.ProductCode, v.ProductPrice)
			result = append(result, s)
		} else {
			result = append(result, product[i])
		}
	}

	return result

}

// product discount specification
type ProductDiscountSpecification struct {
	Product models.Product
}

func (d *ProductDiscountSpecification) IsSatisfied(productes *models.Product) bool {

	//get Product name
	productNames := GetProductNameByCode(d.Product.ProductCode)
	productPrice := GetProductPriceByCode(d.Product.ProductCode)

	//if product contains in the list of special offers return satisfied else return false
	if _, exist := SpecialOffers[productNames]; exist {
		return true
	}
	if _, exist := ProductNameAndPrice[productPrice]; exist {
		return true
	}

	return false
}
