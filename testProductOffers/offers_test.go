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
		want     bool
	}{
		{testName: "Test-Case-One", pQuanity: 3, pCode: "FR1", pPrice: 3.10, want: true},
		{testName: "Test-Case-One", pQuanity: 3, pCode: "SR1", pPrice: 5.10, want: true},
		{testName: "Test-Case-One", pQuanity: 3, pCode: "FR1", pPrice: 11.54, want: true},
	}

	for _, tc := range testSpecialOffer {
		got := productoffers.BuyOneGetOneFree(tc.pQuanity, tc.pCode, tc.pPrice)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("name : %s, expected: %v, got : %v", tc.testName, tc.want, got)
		}
	}

}

func TestDiscountOnAProduct(t *testing.T) {

	testSpecialOffer := []struct {
		testName string
		pQuanity uint
		pCode    string
		pPrice   float64
		want     bool
	}{
		{testName: "Test-Case-One", pQuanity: 3, pCode: "FR1", pPrice: 3.10, want: true},
		{testName: "Test-Case-One", pQuanity: 3, pCode: "SR1", pPrice: 5.10, want: true},
		{testName: "Test-Case-One", pQuanity: 3, pCode: "FR1", pPrice: 11.54, want: true},
	}

	for _, tc := range testSpecialOffer {
		got := productoffers.DiscountOnAProduct(tc.pQuanity, tc.pCode, tc.pPrice)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("name : %s, expected: %v, got : %v", tc.testName, tc.want, got)
		}
	}

}
