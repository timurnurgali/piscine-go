package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {

	runes := []rune(os.Args[0])

	for _, b := range runes {
		z01.PrintRune(b)
	}
	z01.PrintRune('\n')
}
