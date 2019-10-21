package piscine

func IterativeFactorial(nb int) int {
	if nb >= 0 && nb < 13 {
		res := 1
		for i := 2; i <= nb; i++ {
			res *= i
		}
		return res
	}
	return 0

}

// func main() {
// 	println(IterativeFactorial(25))
// }
