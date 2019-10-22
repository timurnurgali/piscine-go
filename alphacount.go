package piscine

// package main

// import "github.com/01-edu/z01"
// import "fmt"

func AlphaCount(str string) int {

	count := 0

	slice := []byte(str)
	for _, l := range slice {
		if l >= 'A' && l <= 'Z' || l >= 'a' && l <= 'z' {
			count++
		}
	}
	return count
}

// func main() {
// 	str := "Hello 78 World!    44y55 /"
// 	nb := AlphaCount(str)
// 	fmt.Println(nb)

// }
