package testProductOffers

import (
	"reflect"
	"testing"

	"github.com/moretonb/moj-durotimicodes-challenge/productoffers"
)

func TestBuyOneGetOneFree(t *testing.T) {

	testSpecialOffer := []struct {
		testName string
		pQuanity uint
		pCode    string
		pPrice   float64
		wantOne  uint
		wantTwo  bool
	}{
		{testName: "Test-Case-One", pQuanity: 3, pCode: "FR1", pPrice: 3.10, wantOne: 1, wantTwo: false},
		{testName: "Test-Case-Two", pQuanity: 3, pCode: "SR1", pPrice: 5.10, wantOne: 3, wantTwo: false},
		{testName: "Test-Case-Three", pQuanity: 3, pCode: "FR1", pPrice: 11.54, wantOne: 3, wantTwo: false},
	}

	for _, tc := range testSpecialOffer {
		_, gotOne, gotTwo := productoffers.BuyOneGetOneFree(tc.pQuanity, tc.pCode, tc.pPrice)
		if !reflect.DeepEqual(tc.wantOne, gotOne) {
			if !reflect.DeepEqual(tc.wantTwo, gotTwo) {
				t.Fatalf("name : %s, expectedOne: %v, expectedTwo: %v, gotOne : %v, gotTwo : %v", tc.testName, tc.wantOne, tc.wantTwo, gotOne, gotTwo)
			}

		}
	}

}

func TestDiscountOnAProduct(t *testing.T) {

	testSpecialOffer := []struct {
		testName string
		pQuanity uint
		pCode    string
		pPrice   float64
		wantOne  uint
		wantTwo  bool
	}{
		{testName: "Test-Case-One", pQuanity: 3, pCode: "FR1", pPrice: 3.10, wantOne: 3, wantTwo: false},
		{testName: "Test-Case-Two", pQuanity: 3, pCode: "SR1", pPrice: 5.10, wantOne: 3, wantTwo: false},
		{testName: "Test-Case-Three", pQuanity: 3, pCode: "FR1", pPrice: 11.54, wantOne: 3, wantTwo: false},
	}

	for _, tc := range testSpecialOffer {
		_, gotOne, gotTwo := productoffers.BuyOneGetOneFree(tc.pQuanity, tc.pCode, tc.pPrice)
		if !reflect.DeepEqual(tc.wantOne, gotOne) {
			if !reflect.DeepEqual(tc.wantTwo, gotTwo) {
				t.Fatalf("name : %s, expectedOne: %v, expectedTwo: %v, gotOne : %v, gotTwo : %v", tc.testName, tc.wantOne, tc.wantTwo, gotOne, gotTwo)
			}

		}
	}

}
