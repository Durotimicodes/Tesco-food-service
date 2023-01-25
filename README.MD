# CHECKOUT SHOPPING CART

## Getting Started

### Prerequistes 

To start using this application these prerequistes must be followed sequentially
- Install **[Go](https://go.dev/)**: (https://go.dev/doc/devel/release) (now version **1.16+** is required).

- Clone repository on the local system using any of these address 

```sh
$ git clone git@github.com:moretonb/moj-durotimicodes-challenge.git
```

```sh 
$ git clone https://github.com/moretonb/moj-durotimicodes-challenge.git
```
### Usage
- Update and install all package dependencies by running this command in the terminal

```sh
go mod tidy
```
- Compile the main application by running this command in the terminal
```sh
go run main.go
```
- Check test cases : invoke test cases by running this command in the terminal 
```sh 
go test
```
- Check test case coverage : go tool has a built in branch coverage. We can invoke it like this:
```sh
go test -coverprofile=c.out
```
## Documentation

### Approach To Solving The Challenge
- Following software development best practices, started the design by creating different packages/folders to seperate concerns.

- Designed the fundamental models/structure, custom types and data stores for the application. For example, Product and CheckoutBill structs in the models package.

- Since i have few products i decided to store data in-memory using data-structures like Maps and slice because they are both extensible and unique for various cases. For instance, although data in a map returns a random order the key remains unique.  

- Declared the basic functions in accordance with the user-story and wrote minimal unit tests that failed.

- To make the program easy to read and reduce the complexity of functions and methods i built easy to use helper functions.

- Implemented the buy-one-get-one free and discount on product functions.

- To enable extension/flexibility of the methods i used the (OPEN CLOSED DESIGN PATTERN PRINCIPLE). As such, the method is opened for extension should incase the CEO or COO change their mind but closed for modification. 

- Created (SpecificationSettings structure) that can hold any further feature the CEO or COO may require in the future.

- To implement the checkout function, i declared helper functions such as *GetTotalPriceOfProduct* and *GetTotalQuantitySoldProducts* and checked for errors should in case an error occurs while checking-out.

- Printed out the bill / receipt to standout.

- Build test cases

### Main Package
```sh
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

```
### Helper Funtions
```sh
// GetProductNameByCode gets the product name by the product code

func GetProductNameByCode(productCode string) string {}

// GetProductPriceByCode gets the product price by the product code

func GetProductPriceByCode(productCode string) float64 {}

// GetProductionBonusQuantityByCode get the additional bonus quantity added to product

func GetProductBonusQuanityByCode(productCode string) uint {}


// GetTotalQuantitySoldProducts gets the total product sold

func GetTotalQuantitySoldProducts(product []models.Product) (uint, error) {}

// GetTotalPriceOfProduct gets the total price of product sold

func GetTotalPriceOfProduct(product []models.Product) (float64, error) {}
```

### Methods
```sh
// product type specification

type ByOneGetOneSpecification struct {
	Product []models.Product
}

// get one product is satisfied

func (b *ByOneGetOneSpecification) IsSatisfied(productes *models.Product) bool {}


// product discount specification

type ProductDiscountSpecification struct {
	Product []models.Product
}

// discount on product is satisfied

func (d *ProductDiscountSpecification) IsSatisfied(productes *models.Product) bool {}

// CheckSpecifications checks if product matches the spec, if yes it implements the specification else returns the real price

func (s *SpecificationSettings) CheckSpecifications(product []models.Product, spec models.OfferSpecification) []models.Product {}

```


### Functions
```sh

// BuyOneGetOneFree takes in a product code, quantity and price and checks if product has a special offer

func BuyOneGetOneFree(prodQuantity uint, prodCode string, prodPrice float64) models.Product {}


// DiscountOnAProduct takes in a product code, quantity and price and checks if product has a discount offer

func DiscountOnAProduct(prodQuantity uint, prodCode string, prodPrice float64) models.Product {}

```

### Reference
 - UDEMY [https://www.udemy.com/course/design-patterns-go/]


