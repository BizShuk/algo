package util

import (
	"bufio"
	"fmt"
	"os"
)

// [Pattern]: [Read stdio]
func ReadStdin() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
	return text
}

func ReadStdinOneline() string {
	var input string
	fmt.Scanln(input) // Just for one line
	return input
}

func ReadStdinOneline2() (input string) {
	fmt.Sscanln("%v", input)
	return input
}
