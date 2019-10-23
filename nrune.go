// package main
package piscine

// import "github.com/01-edu/z01"

func NRune(s string, n int) rune {
	for index, r := range []rune(s) {
		if index+1 == n {
			return r
		}
	}
	return '\x00'

// func main() {
// 	z01.PrintRune(NRune("Tadasdasdsad",0))
// }