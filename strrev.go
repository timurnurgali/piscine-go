package piscine

//import "github.com/01-edu/z01"

func StrRev(s string) string {
	sl := []byte(s)
	var stringg string
	for i := len(s) - 1; i >= 0; i-- {
		stringg += sl[i]
		//z01.PrintRune(rune(sl[i]))
	}
	return stringg
}
