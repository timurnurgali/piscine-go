package piscine

// package main

func IsPrime(nb int) bool {
	isprime := true
	for i := 2; i < nb; i++ {
		if nb%i != 0 { //esli est ostatok
			continue
		} else {
			isprime = false
			break
		}
	}
	return isprime
}

// func main() {
// 	println(IsPrime(29))
// }
