package main

import "fmt"

const Pi = 3.14

func constants() {
	const World = "世界"
	fmt.Print("[constants.go]")
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
