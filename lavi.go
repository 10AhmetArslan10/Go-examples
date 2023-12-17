package main

import "github.com/01-edu/z01"

func main() {
	for _, i := range "HELLO" {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}
