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
func main() {
	email := "user@example.com"
	username, domain := splitEmail(email)
	fmt.Printf("Username: %s, Domain: %s\n", username, domain)

}
