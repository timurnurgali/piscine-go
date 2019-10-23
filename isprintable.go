package piscine

func IsPrintable(str string) bool {
	slice := []rune(str)
	for _, r := range slice {
		if r == '\t' || r == '\n' || r == '\v' || r == '\f' || r == '\r' {
			return false
		}
	}
	return true
}
