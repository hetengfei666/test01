package main

import "fmt"

func swap(a, b string) (string, string) {
	var c string
	c = b
	b = a
	a = c
	return a, b
}
func main() {
	a := "htf"
	b := "gsy"
	result1, result2 := swap(a, b)
	fmt.Println(result1, result2)
}
