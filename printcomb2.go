package piscine

import "github.com/01-edu/z01"

func PrintComb2() {

	for i := '0'; i <= '9'; i++ {
		for j := '0'; j <= '9'; j++ {
			for k := i; k <= '9'; k++ {
				for m := '0'; m <= '9'; m++ {
					if i == k && m <= j {
						continue
					}
					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(' ')
					z01.PrintRune(k)
					z01.PrintRune(m)
					if i != '9' || j != '8' || k != '9' || m != '9' {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
	}
	z01.PrintRune(10)
}
