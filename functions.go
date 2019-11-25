package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func functions() {
	fmt.Println("[functions.go]", add(42, 13))
}
