// package main
package piscine

// import "fmt"

func BasicAtoi2(s string) int {

	ln := 0
	num := 0

	for _, strl := range s {
		if rune(strl) < '0' || rune(strl) > '9' {
			return 0
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
	return num
}

// func main() {
// 	s := "0012h3" // must return 0
// 	n := BasicAtoi(s)
// 	fmt.Println(n)
// }
