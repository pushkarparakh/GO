package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const secondsInMinute = 60
	const minutesInHour = 60
	const secondsInHour = 3600

	//you cannot declare a constant that can only be computed at run-time like you can in JavaScript.
	fmt.Println("number of seconds in an hour:", secondsInHour)
	fmt.Println("===================================================================================================")

	const name = "Saul Goodman"
	const openRate = 30.54

	msg := fmt.Sprintf("Hi %s", name) + fmt.Sprintf(", your open rate is %.1f percent\n", openRate)

	fmt.Print(msg)
	fmt.Println("===================================================================================================")

	const name1 = "🐻"
	fmt.Printf("constant 'name' byte length: %d\n", len(name1))
	fmt.Printf("constant 'name' rune length: %d\n", utf8.RuneCountInString(name1))
	fmt.Println("=====================================")
	fmt.Printf("Hi %s, so good to have you back in the arcanum\n", name1)

	fmt.Println("===================================================================================================")
	fname := "Dalinar"
	lname := "Kholin"
	age := 45
	messageRate := 0.5
	isSubscribed := false
	message := "Sometimes a hypocrite is nothing more than a man in the process of changing."

	userLog := fmt.Sprintf("Name: %s %s, Age : %d, Rate: %g, Is Subscribed: %t, Message: %s", fname, lname, age, messageRate, isSubscribed, message)

	fmt.Println(userLog)
}
