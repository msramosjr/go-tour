package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func multiple_results() {
	a, b := swap("hello", "world")
	fmt.Println("[multiple-results.go]", a, b)
}
