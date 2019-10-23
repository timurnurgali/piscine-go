package piscine

func IsUpper(str string) bool {
	slice := []rune(str)
	for _, r := range slice {
		if r < 'A' || r > 'Z' {
			return false
		}
	}
	return true
}
