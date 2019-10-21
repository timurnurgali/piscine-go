package piscine

// package main

func Sqrt(nb int) int {
	if nb == 0 {
		return 0
	}
	if nb == 1 {
		return 1
	}
	num := 0
	for i := 2; i <= nb; i++ {
		if nb%i == 0 && nb/i == i {
			num = i
			goto p
		}
	}
p:
	return num
}

// func main() {
// 	println(Sqrt(625))
// }
