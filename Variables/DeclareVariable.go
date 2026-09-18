//go:build ignore
// +build ignore

package main

import "fmt"

func declareSMSLimit() {
	var smsSendingLimit int
	smsSendingLimit = 1000
	fmt.Println("Your SMS sending limit is", smsSendingLimit)
}

func main() {
	declareSMSLimit()
}
