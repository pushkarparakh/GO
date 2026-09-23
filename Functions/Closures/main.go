package main

// A closure is a function that references variables from outside its own function body.
// The function may access and assign to the referenced variables.
func adder() func(int) int {
	s := 0
	return func(x int) int {
		s += x
		return s
	}
}

func main() {
	agg := adder()
	println(agg(1))
	println(agg(2))
	println(agg(3))
}
