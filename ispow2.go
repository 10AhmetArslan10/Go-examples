package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	arr := os.Args[1:]
	if len(arr) != 1 {
		return
	}
	sayi,_ := strconv.Atoi(arr[0])
	if Ispowerof2(sayi) {
		Printstr("true\n")
	} else {
		Printstr("false\n")
	}
}

func Printstr(s string) {
	for _, i := range s {
		z01.PrintRune(i)
	}
}

func Ispowerof2(n int) bool {
	if n == 1 {
		return true
	} else if n%2 != 0 {
		return false
	}
	return Ispowerof2(n / 2)
}
