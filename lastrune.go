// package main
package piscine

// import "github.com/01-edu/z01"

func LastRune(s string) rune {

	srun := []rune(s)
	ln := 0
	for range s {
		ln++
	}
	return srun[ln-1]
}

// func main() {
// 	z01.PrintRune(LastRune("Hellowwwwx"))
// }
