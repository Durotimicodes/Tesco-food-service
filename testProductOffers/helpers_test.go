package testProductOffers

import (
	"reflect"
	"testing"

	"github.com/moretonb/moj-durotimicodes-challenge/models"
	"github.com/moretonb/moj-durotimicodes-challenge/productoffers"
)

func TestGetProductNameByCode(t *testing.T) {
	tests := []struct {
		testName string
		input    string
		want     string
	}{
		{testName: "Test-Case-One", input: "FR1", want: "Fruit tea"},
		{testName: "Test-Case-Two", input: "FR2", want: "Sorry do not have this Product"},
		{testName: "Test-Case-Three", input: "SR1", want: "Strawberries"},
		{testName: "Test-Case-Four", input: "SR2", want: "Sorry do not have this Product"},
		{testName: "Test-Case-Five", input: "CF1", want: "Coffee"},
		{testName: "Test-Case-Six", input: "CF2", want: "Sorry do not have this Product"},
	}

	for _, tc := range tests {
		got := productoffers.GetProductNameByCode(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("name : %s, expected: %v, got : %v", tc.testName, tc.want, got)
		}
	}

}

func TestGetProductPriceByCode(t *testing.T) {

	tests := []struct {
		testName string
		input    string
		want     float64
	}{
		{testName: "Test-Case-One", input: "FR1", want: 3.11},
		{testName: "Test-Case-Two", input: "FR2", want: 0},
		{testName: "Test-Case-Three", input: "SR1", want: 5.00},
		{testName: "Test-Case-Four", input: "SR2", want: 0},
		{testName: "Test-Case-Five", input: "CF1", want: 11.23},
		{testName: "Test-Case-Six", input: "CF2", want: 0},
	}

	for _, tc := range tests {
		got := productoffers.GetProductPriceByCode(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("name : %s, expected: %v, got : %v", tc.testName, tc.want, got)
		}
	}

}

func TestGetTotalQuantitySoldProducts(t *testing.T) {

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
	p2 := models.Product{
		ProductId:        1,
		ProductCode:      "SR1",
		ProductQuantity:  3,
		ProductPrice:     5.00,
		ProductDiscount:  0,
		HaveDiscount:     true,
		HaveSpecialOffer: false,
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
		want     uint
	}{

		{testName: "Test-Case-One", input: prodOne, want: 13},
		{testName: "Test-Case-Two", input: prodTwo, want: 17},
		{testName: "Test-Case-Three", input: prodThree, want: 0},
	}

	for _, tc := range tests {
		got, _ := productoffers.GetTotalQuantitySoldProducts(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("name : %s, expected: %v, got : %v", tc.testName, tc.want, got)
		}
	}

}

func TestGetTotalPriceOfProduct(t *testing.T) {

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
		want     float64
	}{

		{testName: "Test-Case-One", input: prodOne, want: 46.099999999999994},
		{testName: "Test-Case-Two", input: prodTwo, want: 172.22},
		{testName: "Test-Case-Three", input: prodThree, want: 0},
	}

	for _, tc := range tests {
		got, _ := productoffers.GetTotalPriceOfProduct(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("name : %s, expected: %v, got : %v", tc.testName, tc.want, got)
		}
	}
}
