package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {

	slice := os.Args[1:] //"a" "b" "1" "A"
	var array [256]int

	for _, b := range slice {
		array[[]byte(b)[0]]++
	}
	for i, tick := range array {
		if tick > 0 {
			z01.PrintRune(rune(i))
		}
	}
}
