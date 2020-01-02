package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var _ = strconv.Itoa // Ignore this comment. You can still use the package "strconv".

	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "

	scanner := bufio.NewScanner(os.Stdin)
	// Declare second integer, double, and String variables.
	var secondInt uint64
	var secondFloat float64
	var secondString string

	// Read and save an integer, double, and String to your variables.
	scanner.Scan()
	secondInt, _ = strconv.ParseUint(scanner.Text(), 10, 64)

	scanner.Scan()
	secondFloat, _ = strconv.ParseFloat(scanner.Text(), 32)

	scanner.Scan()
	secondString = scanner.Text()

	// Print the sum of both integer variables on a new line.
	fmt.Println(i + secondInt)

	// Print the sum of the double variables on a new line.
	fmt.Printf("%.1f\n", d+secondFloat)

	// Concatenate and print the String variables on a new line
	concatenated := s + secondString
	fmt.Println(concatenated)
	// The 's' variable above should be printed first.

}
