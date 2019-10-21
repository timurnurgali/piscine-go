package piscine

func RecursiveFactorial(nb int) int {
	if nb == 0 || nb == 1 {
		return 1
	}
	if nb > 1 && nb < 13 {
		return nb * RecursiveFactorial(nb-1)
	}
	return 0
}

// func main() {
// 	println(RecursiveFactorial(-8))
// }
