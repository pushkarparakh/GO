package main

import "fmt"

func splitEmail(email string) (string, string) {
	/*{
		username, domain := "", ""
	}*/
	//this will generate compiler error because the variables are declared in a block scope and are not accessible outside of it.
	username, domain := "", ""
	for i, r := range email {
		if r == '@' {
			username = email[:i]
			domain = email[i+1:]
			break
		}
	}
	return username, domain
}

func placeOrder(productID string, quantity int, accountBalance float64) (bool, float64) {
	price := calcPrice(productID, quantity)
	if price > accountBalance {
		return false, accountBalance
	}
	stock := amountInStock(productID)
	if stock < quantity {
		return false, accountBalance
	}
	accountBalance -= price
	return true, accountBalance
}

// Don't touch below this line

func calcPrice(productID string, quantity int) float64 {
	return priceList(productID) * float64(quantity)
}

func priceList(productID string) float64 {
	if productID == "1" {
		return 1.50
	} else if productID == "2" {
		return 2.25
	} else if productID == "3" {
		return 3.00
	} else if productID == "4" {
		return 1.00
	} else if productID == "5" {
		return 2.50
	} else if productID == "6" {
		return 8.99
	} else if productID == "7" {
		return 22.50
	} else if productID == "8" {
		return 50.00
	} else if productID == "9" {
		return 999.99
	} else {
		return 0.00
	}
}

func amountInStock(productID string) int {
	if productID == "1" {
		return 11
	} else if productID == "2" {
		return 25
	} else if productID == "3" {
		return 4
	} else if productID == "4" {
		return 6
	} else if productID == "5" {
		return 50
	} else if productID == "6" {
		return 2
	} else if productID == "7" {
		return 0
	} else if productID == "8" {
		return 99
	} else if productID == "9" {
		return 1
	} else {
		return 0
	}
}

func main() {
	email := "user@example.com"
	username, domain := splitEmail(email)
	fmt.Printf("Username: %s, Domain: %s\n", username, domain)

}
