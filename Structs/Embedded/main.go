package main

import "fmt"

type car struct {
  brand string
  model string
}

type truck struct {
  // "car" is embedded, so the definition of a
  // "truck" now also additionally contains all
  // of the fields of the car struct
  car
  bedSize int
}
type sender struct {
	rateLimit int
}

type user struct {
	name   string
	number int
	sender sender
}

//Embedded vs. Nested
// Unlike nested structs, an embedded struct's fields are accessed at the top level like normal fields.
// Like nested structs, you assign the promoted fields with the embedded struct in a composite literal.



//Struct Methods
type rect struct {
  width int
  height int
}

// area has a receiver of (r rect)
// rect is the struct
// r is the placeholder
func (r rect) area() int {
  return r.width * r.height
}
//While Go is not object-oriented, it does support methods that can be defined on structs. 
//Methods are just functions that have a receiver. A receiver is a special parameter that syntactically goes before the name of the function.

var r = rect{
  width: 5,
  height: 10,
}

type authenticationInfo struct {
	username string
	password string
}

// create the method below
func (getBasicAuthInfo authenticationInfo) getBasicAuth() string {
  return getBasicAuthInfo.username + ":" + getBasicAuthInfo.password
}

func main() {	
fmt.Println(r.area())
// prints 50
}