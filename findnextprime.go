package piscine

// package main

func FindNextPrime(nb int) int {
	if nb <= 2 {
		return 2
	}
	a := 0
	for i := nb; i > 2; i++ {
		for k := 2; k < i; k++ {

			if i%k == 0 {
				return FindNextPrime(nb + 1)
			}

		}
		a = i
		break
	}
	return a
}

// func main() {
// 	println(FindNextPrime(2))
// }
