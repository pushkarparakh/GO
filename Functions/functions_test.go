package main

import (
	"fmt"
	"testing"
)

func TestGetMonthlyPrice(t *testing.T) {
	type testCase struct {
		tier     string
		expected int
	}

	runCases := []testCase{
		{tier: "basic", expected: 10000},
		{tier: "premium", expected: 15000},
		{tier: "enterprise", expected: 50000},
	}

	submitCases := append(runCases, []testCase{
		{tier: "invalid", expected: 0},
		{tier: "", expected: 0},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	skipped := len(submitCases) - len(testCases)

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		output := getMonthlyPrice(test.tier)
		if output != test.expected {
			failCount++
			t.Errorf(`---------------------------------
Inputs:     (%v)
Expecting:  %v
Actual:     %v
Fail
`, test.tier, test.expected, output)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     (%v)
Expecting:  %v
Actual:     %v
Pass
`, test.tier, test.expected, output)
		}
	}

	fmt.Println("---------------------------------")
	if skipped > 0 {
		fmt.Printf("%d passed, %d failed, %d skipped\n", passCount, failCount, skipped)
	} else {
		fmt.Printf("%d passed, %d failed\n", passCount, failCount)
	}
}

var withSubmit = true

func TestBilling(t *testing.T) {
	type testCase struct {
		name         string
		costPerSend  int
		lastMonth    int
		thisMonth    int
		expectedBill int
		expectedDiff int
	}

	testCases := []testCase{
		{name: "increase", costPerSend: 2, lastMonth: 100, thisMonth: 150, expectedBill: 300, expectedDiff: 100},
		{name: "decrease", costPerSend: 5, lastMonth: 40, thisMonth: 10, expectedBill: 50, expectedDiff: -150},
		{name: "no change", costPerSend: 3, lastMonth: 20, thisMonth: 20, expectedBill: 60, expectedDiff: 0},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			bill := getBillForMonth(test.costPerSend, test.thisMonth)
			if bill != test.expectedBill {
				t.Fatalf("getBillForMonth(%d, %d) = %d, want %d", test.costPerSend, test.thisMonth, bill, test.expectedBill)
			}

			increase := monthlyBillIncrease(test.costPerSend, test.lastMonth, test.thisMonth)
			if increase != test.expectedDiff {
				t.Fatalf("monthlyBillIncrease(%d, %d, %d) = %d, want %d", test.costPerSend, test.lastMonth, test.thisMonth, increase, test.expectedDiff)
			}
		})
	}
}