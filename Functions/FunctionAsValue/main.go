package main

import "fmt"

// reformat takes a string and a FUNCTION that formats the string in some way.
func reformat(message string, formatter func(string) string) string {
	formatted := message
	for i := 0; i < 3; i++ {
		formatted = formatter(formatted)
	}
	return "TEXTIO: " + formatted
}

// main demonstrates how to use a FUNCTION as a VALUE by passing it as an argument to another FUNCTION.
func main() {
	addPeriod := func(message string) string {
		return message + "."
	}

	fmt.Println(reformat("General Kenobi", addPeriod))
}

// aggregate takes three integers and a FUNCTION that performs an arithmetic operation on two integers.
func aggregate(a, b, c int, arithmetic func(int, int) int) int {
	firstResult := arithmetic(a, b)
	secondResult := arithmetic(firstResult, c)
	return secondResult
}
