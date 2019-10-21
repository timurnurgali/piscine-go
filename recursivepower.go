package piscine

// package main

func RecursivePower(nb int, power int) int {
	if power < 0 {
		return 0
	}
	if power == 0 {
		return 1
	}
	return nb * RecursivePower(nb, power-1)

}

// func main() {
// 	println(IterativePower(4,3))
// }
