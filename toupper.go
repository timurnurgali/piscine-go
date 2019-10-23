package piscine

func ToUpper(s string) string {
	a := []rune(s)
	b := 'A' - 'a'
	for i := range runes {
		if runes[i] >= 'a' && runes[i] <= 'z' {
			runes[i] += diff
		}
	}
	return string(runes)
}
