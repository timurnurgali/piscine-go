package piscine

func IsNumeric(str string) bool {
	slice := []rune(str)
	for _, r := range slice {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}
