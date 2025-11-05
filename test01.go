package main

import "fmt"

var a, b, c = 1, 2, 3
var (
	d int
	f string
)

func main() {
	d = 0
	f = "	qima"
	e := "adsggh"
	fmt.Println("%v %v %v %v %q\n", a, b, c, d, f)
	fmt.Print(e, a, b, "\n")

	fmt.Println(e, a, b)
	// fmt.Printf("%v %v %v %v %q\n", a, b, c, d, f)

}
