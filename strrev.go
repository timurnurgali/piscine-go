package piscine

//import "github.com/01-edu/z01" NOT FINISHED YET

func StrRev(s string) string {
	sl := []byte(s)
	var newS []byte
	ln := 0
	for range s {
		ln += 1
		//z01.PrintRune(rune(sl[i]))
	}
	for i := ln - 1 ; i >= 0 ; i-- {
		newS += sl[i]
	}
	return string(newS)
}
