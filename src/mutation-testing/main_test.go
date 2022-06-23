package main

import "testing"

func TestCanDrink(t *testing.T) {
	testCases := []struct {
		age      int
		expected bool
	}{
		{18, true},
		{10, false},
	}
	for _, testCase := range testCases {
		result := canDrink(testCase.age)
		if result != testCase.expected {
			t.Errorf("Error : canDrink(%d) return %t, expected %t", testCase.age, result, testCase.expected)
		}
	}
}

func TestCanBuyAlcool(t *testing.T) {
	testCases := []struct {
		age      int
		country  string
		expected bool
	}{
		{18, "FR", true},
		{10, "FR", false},
		{21, "US", true},
		{10, "US", false},
	}
	for _, testCase := range testCases {
		result := canBuyAlcool(testCase.age, testCase.country)
		if result != testCase.expected {
			t.Errorf("Error : canBuyAlcool(%d, %s) return %t, expected %t", testCase.age, testCase.country, result, testCase.expected)
		}
	}
}
