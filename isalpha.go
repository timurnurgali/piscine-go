package piscine

func IsAlpha(str string) bool {
	slice := []rune(str)
	for _, r := range slice {
		if !(r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r >= '0' && r <= '9') {
			return false
		}
	}
	return true
}
