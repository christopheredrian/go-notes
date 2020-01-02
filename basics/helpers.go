package helpers

import (
	"bufio"
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
