// package main
package piscine

// import "fmt"

func Atoi(s string) int {

	ln := 0
	num := 0
	sl := []rune(s)
	isneg := false
	for index, b := range sl {
		if index == 0 {
			if b == '-' {
				isneg = true
			}
			if b != '+' && b != '-' {
				if b < '0' || b > '9' {
					return 0
				}
			}
		}
		if index > 0 {
			if b < '0' || b > '9' {
				return 0
			}
		}

		ln++
	}

	for i := 0; i < ln; i++ {

		dig := 0

		for j := '0'; j < rune(s[i]); j++ {
			dig++
		}
		for k := 0; k < ln-i-1; k++ {
			dig *= 10
		}
		num += dig
	}
	if isneg == true {
		num *= -1
	}
	return num
}

// func main() {
// 	s := "+1234" // must return 0
// 	n := Atoi(s)
// 	fmt.Println(n)
// }
