package piscine

import "github.com/01-edu/z01"

func PrintStr(str string) {
	slice := []byte(str)
	ln := len(str)
	for i := 0 ; i < ln ; i++ {
		z01.PrintRune(rune(slice[i]))
	}
}