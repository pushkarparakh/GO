package main

import (
	"fmt"
	"reflect"
)

type contact struct {
	sendingLimit int32
	userID       string
	age          int32
}

// the order of fields in a struct can have a big impact on memory usage. This is the same struct as above, but better designed:
type contact1 struct {
	sendingLimit int32
	age          int32
	userID       string
}

type perms struct {
	canSend         bool
	canReceive      bool
	permissionLevel int
	canManage       bool
}
type perms1 struct {
	permissionLevel int
	canSend         bool
	canReceive      bool
	canManage       bool
}

// anonymous empty struct type

// named empty struct type
type emptyStruct struct{}

var empty = emptyStruct{}

func main() {
	//However, if you have a specific reason to be concerned about memory usage, aligning the fields by size (largest to smallest) can help.
	//You can also use the reflect package to debug the memory layout of a struct:
	typ := reflect.TypeOf(contact{})
	fmt.Printf("Size of contact struct %d bytes", typ.Size())
	println()

	typ1 := reflect.TypeOf(contact1{})
	fmt.Printf("Size of contact1 struct %d bytes.", typ1.Size())
	println()

	typ2 := reflect.TypeOf(perms{})
	fmt.Printf("Size of perms struct %d bytes.", typ2.Size())
	println()

	typ3 := reflect.TypeOf(perms1{})
	fmt.Printf("Size of perms1 struct %d bytes.", typ3.Size())
	println()

	emptyTyp := reflect.TypeOf(emptyStruct{})
	fmt.Printf("Size of empty struct %d bytes.", emptyTyp.Size())
	println()

}
