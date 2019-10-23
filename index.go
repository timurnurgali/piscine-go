// package main
package piscine

func Index(s string, toFind string) int {

	ln := 0
	for range toFind {
		ln++
	}

	for index := range s {
		if s[index] == toFind[0] {
			if s[index:index+ln] == toFind[:] {
				return index
			}
		}
	}
	return -1
}

// func main() {
// 		println(Index("Alem School", "cho"))
// 	}
