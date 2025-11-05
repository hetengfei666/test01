package main

import "fmt"

func copmax(a, b int) int {
	var result int
	if a > b {
		result = a
	} else {
		result = b
	}
	return result
}
func main() {
	a, b := 2, 3
	ret := copmax(a, b)
	fmt.Printf("%d\n", ret)
	fmt.Print(a)
}
