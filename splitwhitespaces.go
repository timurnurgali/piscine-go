package piscine

// package main

// import "fmt"

func SplitWhiteSpaces(str string) []string {

	wordCount := 1
	for _, b := range str {
		if b == ' ' || b == '\t' || b == '\n' {
			wordCount++
		}
	}

	answer := make([]string, wordCount) // 4 slots

	indexSlice := 0
	lastIndex := -1
	for i, b := range str {
		if b == ' ' || b == '\t' || b == '\n' {
			answer[indexSlice] = str[lastIndex+1 : i]
			lastIndex = i
			indexSlice++
		}
		if indexSlice == wordCount-1 {
			answer[indexSlice] = str[lastIndex+1:]
		}
	}
	return answer

}

// func main() {
// 	str := "Hello how are you ?"
// 	fmt.Println(SplitWhiteSpaces(str))
// }
