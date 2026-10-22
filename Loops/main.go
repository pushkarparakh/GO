package main

import "fmt"

func bulkSend(numMessages int) float64 {
	totalCost := 0.0

	for i := 0; i < numMessages; i++ {
		// Each message costs 1.0 plus an additional fee based on its index
		fee := float64(i) * 0.01
		totalCost += 1.0 + fee
	}

	return totalCost
}

func maxMessages(thresh int) int {
	totalCost := 0
	numMessages := 0
	// no condition for loop
	for {
		// Calculate the cost of the next message
		nextMessageCost := 100 + numMessages // 100 pennies + fee based on index
		if totalCost+nextMessageCost > thresh {
			break // Stop if adding the next message exceeds the budget
		}
		totalCost += nextMessageCost // Add the cost of the next message
		numMessages++                // Increment the number of messages sent
	}

	return numMessages // Return the maximum number of messages that can be sent
}

func fizzbuzz() {

	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Println("fizzbuzz")
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}
}

// continue and break
func printPrimes(max int) {
	for i := 2; i <= max; i++ {
		isPrime := true
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Println(i)
		}
	}
}

func test(max int) {
	fmt.Printf("Primes up to %v:\n", max)
	printPrimes(max)
	fmt.Println("===============================================================")
}

func countConnections(groupSize int) int {
	connections := 0
	for i := 1; i < groupSize; i++ {
		connections += i
	}
	return connections
}

func main() {
	fizzbuzz()
	test(10)
	test(20)
	test(30)
	fmt.Printf("Connections in a group of 5 people: %v\n", countConnections(5))
	fmt.Printf("Connections in a group of 10 people: %v\n", countConnections(10))

}
