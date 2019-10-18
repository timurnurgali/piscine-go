package piscine

//import "github.com/01-edu/z01" NOT FINISHED YET

func StrRev(s string) string {

	slice := []byte(s)
	var newslice []byte
	var ln int

	for range s {
		ln += 1
	}

	for i := ln - 1; i >= 0; i-- {
		newslice = append(newslice, slice[i])
	}
	return string(newslice)
}
