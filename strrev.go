package piscine

//import "github.com/01-edu/z01"

func StrRev(s string) string {
	sl := []byte(s)
	var newS string
	ln := 0
	for range s {
		ln += 1
		//z01.PrintRune(rune(sl[i]))
	}
	for _, letter := range sl {
		newS += sl[letter]
	}
	return newS
}
