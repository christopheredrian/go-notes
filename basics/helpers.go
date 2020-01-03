package helpers

import (
	"bufio"
	"fmt"
	"os"
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
