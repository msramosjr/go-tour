package main

import (
	"fmt"
	"math"
)

func type_conversions() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println("[type-conversions.go]", x, y, z)
}
