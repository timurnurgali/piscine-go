package piscine

// package main

// import "fmt"

func ConcatParams(args []string) string {
	result := ""
	ln := 0
	for range args {
		ln++
	}

	for i := 0; i < ln; i++ {
		result += args[i]
		if i < ln-1 {
			result += "\n"
		}

	}
	return result

}

// func main() {
// 	test := []string{"Hello", "how", "are", "you?"}
// 	fmt.Println(ConcatParams(test))
// }
