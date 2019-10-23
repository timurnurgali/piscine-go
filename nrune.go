// package main
package piscine

// import "github.com/01-edu/z01"

func NRune(s string, n int) rune {

	if n > 0 {
		var convert rune
		srun := []rune(s)
		b := n - 1
		convert = rune(srun[b])
		return convert
	}
	return 0
}

// func main() {
// 	z01.PrintRune(NRune("Tadasdasdsad",3)) //u index 3
// }
