package piscine

func IterativeFactorial(nb int) int {
	if nb >= 0 {
		res := 1
		for i := 2; i <= nb; i++ {
			res *= i
		}
		if res >= 1 {
			return res
		}
	}
	return 0

}

// func main() {
// 	println(IterativeFactorial(25))
// }
