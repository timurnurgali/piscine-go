package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	for i := 0; i < len(n)-1; i++ {
		z01.PrintRune(rune(n[i]))
	}
}
