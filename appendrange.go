package piscine

// package main

// import "fmt"

func AppendRange(min, max int) []int {

	var answer []int
	if min >= max {
		return answer
	}

	for i := min; i < max; i++ {
		answer = append(answer, i)
	}
	return answer

}

// func main() {
// 	fmt.Println(AppendRange(11, 10))
// }
