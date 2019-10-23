package piscine

func IsLower(str string) bool {
	slice := []rune(str)
	for _, r := range slice {
		if r < 'a' || r > 'z' {
			return false
		}
	}
	return true
}
