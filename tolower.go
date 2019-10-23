package piscine

func ToLower(s string) string {
	slice := []rune(s)
	b := 'A' - 'a'
	for i := range slice {
		if slice[i] >= 'A' && slice[i] <= 'Z' {
			slice[i] -= b
		}
	}
	return string(slice)
}
