// package main
package piscine

// import "fmt"
// import "github.com/01-edu/z01"

func StrRev(s string) string {

	temp := rune(s)
	bSlice := []rune(s)
	var ln int

	for range s {
		ln += 1
	}

	for i := 0; i < ln; i++ {
		bSlice[i] = temp[ln-i-1]
	}
	return string(bSlice)
}

// func main() {
// 	s := "Hello World!"
// 	s = StrRev(s)
// 	fmt.Println(s)
// }
