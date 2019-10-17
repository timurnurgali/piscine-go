package piscine

import "github.com/01-edu/z01"

func PrintStr(str string) {
	sl := []byte(str)
	for _, letter := range sl {
		z01.PrintRune(rune(letter))
	}
}
