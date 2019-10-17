package piscine

func StrLen(str string) int {
	n := 0
	for _, letter := range str {
		n += 1
	}
	return n
}
