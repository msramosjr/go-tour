package main

import "fmt"

func type_inference() {
	fmt.Print("[type_inference.go]")
	v := 42 // change me!
	fmt.Printf("v is of type %T\n", v)
	v2 := 42. // change me!
	fmt.Printf("v is of type %T\n", v2)
	v3 := 42 + 4i // change me!
	fmt.Printf("v is of type %T\n", v3)
}
