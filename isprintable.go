package piscine

func IsPrintable(str string) bool {
	var t bool = true
	var f bool = false
	ln := 0
	b := 0
	for range str {
		ln++
	}
	str_rune := []rune(str)
	for i := 0; i < ln; i++ {
		if str_rune[i] >= 32 && str_rune[i] <= 127 {
			b++
		}

	}
	if ln == b {
		return t
	} else {
		return f
	}

}
