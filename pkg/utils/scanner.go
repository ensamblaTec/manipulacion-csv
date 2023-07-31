package utils

import (
	"bufio"
	"os"
)

func ReadInput(input *string) {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()

	*input = scanner.Text()
}
