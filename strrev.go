// package main
package piscine

// import "fmt"
// import "github.com/01-edu/z01"

func StrRev(s string) string {

	bSlice := []byte(s)
	var ln int
	temp := s

	for range s {
		ln += 1
	}

	for i := 0; i < ln; i++ {
		s[i] = temp[ln-i-1]
	}
	return s
}

// func main() {
// 	s := "Hello World!"
// 	s = StrRev(s)
// 	fmt.Println(s)
// }
