package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	vowels = []string{"a", "e", "i", "o", "u"}
)

func main() {
	info, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	if info.Mode()&os.ModeCharDevice != 0 || info.Size() <= 0 {
		fmt.Println("The command is intended to work with pipes.")
		fmt.Println("Usage: echo 'herp derpus' | fewer")
		return
	}

	reader := bufio.NewReader(os.Stdin)
	var output []rune

	for {
		input, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		output = append(output, input)
	}

	for j := 0; j < len(output); j++ {
		if !(isVowel(output[j])) {
			fmt.Printf("%c", output[j])
		}

	}
}

func isVowel(val rune) bool {
	for _, item := range vowels {
		check := strings.ToLower(string(val))
		if item == check {
			return true
		}
	}
	return false
}
