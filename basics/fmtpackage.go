package main

import "fmt"

func main() {

	myInt := 42

	fmt.Println(myInt)
	fmt.Printf("%T\n", myInt)
	fmt.Printf("%b\n", myInt)  // binary
	fmt.Printf("%x\n", myInt)  // hex
	fmt.Printf("%#x\n", myInt) // hex with 0x

	// String Print
	formattedBinary := fmt.Sprintf("%b\n", myInt) // binary
	fmt.Println(formattedBinary)
}
