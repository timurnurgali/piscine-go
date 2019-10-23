package piscine

func Join(strs []string, sep string) string {
	str := ""
	flag := true
	for _, s := range strs {
		if flag {
			flag = false
		} else {
			str += sep
		}
		str += s
	}
	return str
}
