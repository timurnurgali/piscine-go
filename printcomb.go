package piscine

import "github.com/01-edu/z01"

func PrintComb() {

	for i := '0'; i <= '7'; i++ {
		for j := '1'; j > i && j <= '8'; j++ {
			for k := '2'; k > j && k <= '9'; k++ {
				z01.PrintRune(i)
				z01.PrintRune(j)
				z01.PrintRune(k)
				z01.PrintRune(44)
				z01.PrintRune(32)
			}
		}
	}
	z01.PrintRune(8)
	z01.PrintRune(8)
	z01.PrintRune(10)
}
