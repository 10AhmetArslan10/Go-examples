package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) > 0 {
		last := []rune(args[len(args)-1])
		for _, char := range last {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}
}
