package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func named_results() {
	fmt.Print("[named-results.go] ")
	fmt.Println(split(17))
}
