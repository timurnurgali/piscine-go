// package piscine APPEND ALLOWED

// package main
package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	var array [10]int
	if n == 0 {
		z01.PrintRune('0')
	}
	for n != 0 {
		array[n%10]++
		n /= 10
	}
	for i := 0; i < 10; i++ {
		if array[i] > 0 {
			z01.PrintRune('0' + rune(i))
		}
	}
}

// func main() {
// 	PrintNbrInOrder(348)
// }
