/**
 * Primitive vs Composite Data Type
 */
package main

import "fmt"

var myInteger = 42
var myStr = "Mind is like a mad monkey"
var myStrWithQuotes = `"Min is like a mad monkey"`

type age int

var myAge age
var tempAge int

// zero values: 0, 0.0, "", nil

func main() {

	fmt.Println(myInteger)
	fmt.Printf("Type: %T\n", myInteger) // print the type

	fmt.Println(myStr)
	fmt.Printf("Type: %T\n", myStr)

	fmt.Println(myStrWithQuotes)

	tempAge = 0
	myAge = 24
	fmt.Println(myAge)

	// [notes 1] Conversion (not casting)
	tempAge = int(myAge)
	fmt.Println(tempAge)
}
