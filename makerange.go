package piscine

// package main

// import "fmt"

func MakeRange(min, max int) []int {

	var null []int
	if min >= max {
		return null
	}
	size := max - min
	answer := make([]int, size)
	for i := 0; i < size; i++ {
		answer[i] = min + i
	}
	return answer

}

// func main() {
// 	fmt.Println(MakeRange(11, 10))
// }
