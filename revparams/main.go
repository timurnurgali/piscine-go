package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {

	slStr := os.Args[1:]
	ln := 0
	for range slStr {
		ln++
	}

	for i := ln - 1; i >= 0; i-- {
		r := []rune(slStr[i])
		for _, c := range r {
			z01.PrintRune(c)
		}
		z01.PrintRune('\n')

	}
}
