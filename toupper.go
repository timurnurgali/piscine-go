package piscine

func ToUpper(s string) string {
	slice := []rune(s)
	b := 'A' - 'a'
	for i := range slice {
		if slice[i] >= 'a' && slice[i] <= 'z' {
			slice[i] += b
		}
	}
	return string(slice)
}
