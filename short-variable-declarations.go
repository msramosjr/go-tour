package main

import "fmt"

func short_variable_declarations() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println("[short-variable-declarations.go]", i, j, k, c, python, java)
}
