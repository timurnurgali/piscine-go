package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
	}
	check(n)
}

func check(nbr int) {
	digit := '0'
	if nbr == 0 {
		z01.PrintRune(digit)
		return
	}
	for i := 1; i <= nbr%10; i++ {
		digit++
	}
	for i := -1; i >= nbr%10; i-- {
		digit++
	}
	if nbr/10 != 0 {
		check(nbr / 10)
	}
	z01.PrintRune(digit)
	return
}
