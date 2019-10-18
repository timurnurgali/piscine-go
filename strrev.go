// package main
package piscine

// import "fmt"
// import "github.com/01-edu/z01"

func StrRev(s string) string {

	bSlice := []byte(s)
	var ln int
	var rev []byte
	for range s {
		ln += 1
	}

	for i := ln - 1; i >= 0; i-- {
		rev += bSlice[i]
	}
	return string(rev)
}

// func main() {
// 	s := "Hello World!"
// 	s = StrRev(s)
// 	fmt.Println(s)
// }