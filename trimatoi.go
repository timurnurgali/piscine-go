// package main
package piscine

// import "github.com/01-edu/z01"

func TrimAtoi(s string) int {
	sl := []rune(s)
	var numRune []rune
	neg := 1
	negdisabled := true
	lnRune := 0 //length of numRune
	lnStr := 0  //length of string
	finalNum := 0
	for range sl {
		lnStr++
	}

	for i := 0; i < lnStr; i++ {
		if negdisabled {
			if sl[i] == '-' {
				neg = -1
			}
		}

		if sl[i] >= '0' && sl[i] <= '9' {
			negdisabled = false
			numRune = append(numRune, sl[i])
		}
	}

	for range numRune {
		lnRune++
	}
	multiplier := 1
	for i := lnRune - 1; i >= 0; i-- {
		digit := 0
		for j := '0'; j < numRune[i]; j++ {
			digit++
		}
		digit *= multiplier
		multiplier *= 10
		finalNum += digit
	}
	return finalNum * neg

}

// func main() {
// 	ex := TrimAtoi("dsax4-5rd8")
// 	println(ex)
// } //FUCKING YES!!!
