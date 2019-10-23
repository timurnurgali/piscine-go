// package main
package piscine

// import "github.com/01-edu/z01"

func NRune(s string, n int) rune {
	if n <= 0 {
		return '\n'
	}
	srun := []rune(s)
	return srun[n-1]
}

// func main() {
// 	z01.PrintRune(NRune("Tadasdasdsad",0)) //u index 3
// }
