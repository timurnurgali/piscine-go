package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {

	slStr := os.Args // []string{"main", "choumi", "is", "cat"}

	for i, b := range slStr {
		if i == 0 {
			continue
		}
		r := []rune(b)
		for _, c := range r {
			z01.PrintRune(c)
		}
		z01.PrintRune('\n')
	}
}
