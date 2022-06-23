package main

//CanDrink provides a boolean following age
func canDrink(age int) bool {
	if age >= 18 {
		return true
	}
	return false
}

//CanBuyAlcool provides a boolean following age and country
func canBuyAlcool(age int, country string) bool {
	if country == "US" && age >= 21 {
		return true
	} else if country == "FR" && age >= 18 {
		return true
	}
	return false
}
