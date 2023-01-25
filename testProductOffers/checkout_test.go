package testProductOffers

import (
	"reflect"
	"testing"

	"github.com/moretonb/moj-durotimicodes-challenge/models"
	"github.com/moretonb/moj-durotimicodes-challenge/productoffers"
)

func TestCheckOutProducts(t *testing.T) {

	//Product one
	p1 := models.Product{
		ProductId:        1,
		ProductCode:      "FR1",
		ProductQuantity:  10,
		ProductPrice:     3.11,
		ProductDiscount:  0,
		HaveDiscount:     true,
		HaveSpecialOffer: true,
	}

	//Product two
	p2 := models.Product{
		ProductId:        1,
		ProductCode:      "SR1",
		ProductQuantity:  3,
		ProductPrice:     5.00,
		ProductDiscount:  0,
		HaveDiscount:     true,
		HaveSpecialOffer: true,
	}

	//Product three
	p3 := models.Product{
		ProductId:        1,
		ProductCode:      "CF1",
		ProductQuantity:  14,
		ProductPrice:     11.23,
		ProductDiscount:  0,
		HaveDiscount:     true,
		HaveSpecialOffer: false,
	}

	//Product four
	p4 := models.Product{
		ProductId:        1,
		ProductCode:      "CF1",
		ProductQuantity:  0,
		ProductPrice:     11.23,
		ProductDiscount:  0,
		HaveDiscount:     true,
		HaveSpecialOffer: false,
	}

	//a group/slice of products
	prodOne := []models.Product{p1, p2}
	prodTwo := []models.Product{p2, p3}
	prodThree := []models.Product{p4}

	tests := []struct {
		testName string
		input    []models.Product
		wantOne  float64
		wantTwo  uint
	}{

		{testName: "Test-Case-One", input: prodOne, wantOne: 46.099999999999994, wantTwo: 2},
		{testName: "Test-Case-Two", input: prodTwo, wantOne: 172.22, wantTwo: 2},
		{testName: "Test-Case-Three", input: prodThree, wantOne: 0, wantTwo: 1},
	}

	for _, tc := range tests {
		_, gotOne, gotTwo := productoffers.CheckOutProducts(tc.input)
		if !reflect.DeepEqual(tc.wantOne, gotOne) {
			if !reflect.DeepEqual(tc.wantTwo, gotTwo) {
				t.Fatalf("name : %s, expectedOne: %v, expectedTwo: %v, gotOne : %v, gotTwo : %v", tc.testName, tc.wantOne, tc.wantTwo, gotOne, gotTwo)
			}
		}
	}

}
