package main

import "fmt"

var (
	a   = 4
	n   = 6
	ptr *int
)

func main() {
	ptr = &a
	fmt.Printf("a的值%d\n", *ptr)
}
