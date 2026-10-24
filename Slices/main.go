package main

import (
	"errors"
)

// arrays in GO
var myint [10]int
var primes = [10]int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}

func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
	// Create the string array
	strings := [3]string{primary, secondary, tertiary}

	// Create the integer array
	integers := [3]int{len(primary), len(primary) + len(secondary), len(primary) + len(secondary) + len(tertiary)}

	return strings, integers
}

// slices in GO
// A slice is a dynamically-sized, flexible view of the elements of an array.
var mySlice = primes[1:4]

//slicing an array

const (
	planFree = "free"
	planPro  = "pro"
)

func getMessageWithRetriesForPlan(plan string, messages [3]string) ([]string, error) {
	if plan == planPro {
		return messages[:], nil
	}
	if plan == planFree {
		return messages[0:2], nil
	}
	return nil, errors.New("unsupported plan")
}

// Make function
// func make([]T, len, cap) []T
var mySliceWithCap = make([]int, 5, 10)

// the capacity argument is usually omitted and defaults to the length
var mySliceDefaultCap = make([]int, 5)

// If we want to create a slice with a specific set of values, we can use a slice literal:
var mmySlice = []string{"I", "love", "go"}

// Variadic fucntion
// Variadic functions can be called with any number of trailing arguments. For example, fmt.Println is a common variadic function.
func sum(nums ...int) int {
	s := 0
	for i := 0; i < len(nums); i++ {
		s += nums[i]
	}
	return s
}
