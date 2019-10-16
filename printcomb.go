package piscine

import "github.com/01-edu/z01"

func PrintComb() {

	for i := '0'; i <= '7'; i++ {
		z01.PrintRune(i)
		for j := '1'; '8' >= j > i; j++ {
			z01.PrintRune(j)
			for k := '2'; '9' >= k > j; k++ {
				z01.PrintRune(k + ',' + ' ')
			}
		}
	}
	z01.PrintRune(8)
	z01.PrintRune(8)
	z01.PrintRune(10)
}
