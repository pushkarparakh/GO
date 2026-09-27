package main

import "fmt"

// We use structs in Go to represent structured data.
//Structs in Go are often used to represent data that you might use a dictionary or object for in other languages.
// Structs are a way to group together related data into a single type. Each piece of data in a struct is called a field, and each field has a name and a type.

// This creates a new struct type called car.
type car struct {
	brand   string
	model   string
	door    int
	year    int
	mileage int
	//nested structs are structs that are defined within other structs. They allow you to create more complex data structures by combining multiple structs together. In Go, you can define a struct within another struct by simply declaring it as a field of the outer struct.
	frontWheel wheel
	backWheel  wheel
}

type wheel struct {
	radius   int
	material string
}

type messageToSend struct {
	message   string
	sender    user
	recipient user
}

type user struct {
	name   string
	number int
}

var mToSend = messageToSend{
	message: "Hello, how are you?",
	sender: user{
		name:   "Alice",
		number: 1234567890,
	},
	recipient: user{
		name:   "Bob",
		number: 9876543210,
	},
}

func canSendMessage(mToSend messageToSend) bool {
	//This function checks if a message can be sent by verifying that the message, sender's name, and recipient's name are not empty. If any of these fields are empty, it returns false, indicating that the message cannot be sent. Otherwise, it returns true.
	if mToSend.message == "" || mToSend.sender.name == "" || mToSend.recipient.name == "" {
		return false
	}

	return true
}

func main() {
	//To create a car, use a struct literal:
	myCar := car{
		brand:   "Toyota",
		model:   "Camry",
		year:    2020,
		mileage: 15000,
	}
	//The fields of a struct can be accessed using the dot . operator.
	myCar.frontWheel.radius = 15

	fmt.Println(myCar)
}
