package main

import (
	"fmt"
	"math/rand"
	"time"
)

func packages() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("[packages.go] My favorite number is", rand.Intn(10))
}
