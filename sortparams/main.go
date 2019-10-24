package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {

	slice := os.Args[1:] //"a" "b" "1" "A"
	var array [256]int
	var runes [256]rune
	for _, b := range slice {
		for _, c := range b {
			array[[]byte(b)[0]]++
			runes[[]byte(b)[0]] = c
		}
	}

	for i, tick := range array {

		if tick > 0 {
			z01.PrintRune(runes[i])
		}
	}
}
