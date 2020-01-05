package main

import (
	"fmt"
)

var x2 int
var y2 string
var z2 bool

var x3 = 42
var y3 = "James Bond"
var z3 = true

type age int

var y5 int

func main() {
	fmt.Println("Hello, playground")

	// Ex. 1
	x := 42           // asign x as an integer type with an initial value of 42
	y := "James Bond" // assign y as a string type with an initial value of "James Bond"
	z := true         // assign z as a boolean type with an initial value of true

	// single statement print
	fmt.Println(x, y, z)

	// print using multiple print lines
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	fmt.Println("----------")
	// Exercise 2.
	// Use var to DECLARE three package level variables
	// variables below are called zero values
	fmt.Println(x2, y2, z2)

	fmt.Println("----------")
	// Exercise 3.
	// declare package level scope
	s := fmt.Sprintf("%d %s %t", x3, y3, z3)
	s2 := fmt.Sprintf("%v\t%v\t%v", x3, y3, z3)
	fmt.Println(s)
	fmt.Println(s2)

	fmt.Println("----------")
	// Ex. 4 create own type

	var myAge age = 24

	fmt.Println(myAge)
	fmt.Printf("%T\t", myAge)
	myAge = 42
	fmt.Println(myAge)

	// Ex. 5
	fmt.Println("----------")
	y5 = int(myAge)
	fmt.Println(y5)
	fmt.Printf("%T\n", y5)
}
