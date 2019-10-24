package main

import (
	"github.com/01-edu/z01"
	"os"
)

func strToNum(s string) int {
	ln := 0
	num := 0
	for range s {
		ln++
	}
	sl := []rune(s)
	multiplier := 1

	for i := ln - 1; i >= 0; i-- {
		x := 0
		for j := '0'; j < sl[i]; j++ {
			x++
		}
		num += x * multiplier
		multiplier *= 10
	}
	return num
}

func main() {

	args := os.Args[1:]
	isUpper := false

	for _, b := range args { // "12" "13"
		if b == "--upper" {
			isUpper = true
		}
		b2 := strToNum(b)
		if b2 > 0 && b2 < 27 {
			if isUpper {
				z01.PrintRune(rune(b2 + 64))
			} else {
				z01.PrintRune(rune(b2 + 96))
			}
		} else {
			z01.PrintRune(' ')
		}

	}
	z01.PrintRune('\n')

}
