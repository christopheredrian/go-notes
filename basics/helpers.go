package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func getStdIn() []string {
	scanner := bufio.NewScanner(os.Stdin)

	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text()) // Println will add back the final '\n'
	}

	return lines

}

func formatAsFloat() {
	first := 5.0
	second := 3.0
	fmt.Printf("%.1f", first+second)
}

func convertToBytes() {
	myString := "Hello"

	convertedStr := []byte(myString)
	fmt.Println(convertedStr)

	fmt.Println(string(convertedStr))
}

func generateRandomNumber(maxNum) {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	randInt := r.Intn(maxNum - 1) // by default go uses a static seed
	return randInt
}

func main() {
	// convertToBytes()
	generateRandomNumber(maxNum)
}
