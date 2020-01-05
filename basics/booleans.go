package main

import "fmt"

var x bool // zero value is false

func main() {
	fmt.Println(x)
	x = true

	a := 7
	b := 42
	fmt.Println(a > b)
}
