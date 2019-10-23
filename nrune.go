package piscine

func NRune(s string, n int) rune {
	if n > 0 {
		srune := []rune(s)
		return srune[n-1]
	}
	return '\x00'