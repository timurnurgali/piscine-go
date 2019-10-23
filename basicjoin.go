package piscine

func BasicJoin(strs []string) string {
	ret := ""
	for _, s := range strs {
		ret += s
	}
	return ret
}
