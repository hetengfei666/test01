package main

// const (
// 	a = "abc"
// 	b = 12
// 	c = "abc"
// 	d = len(a)
// 	e = unsafe.Sizeof(a)
// )
const (
	a = iota
	b = "hha"
	c = 100
	d = iota
	e
)

func main() {
	// b = 2
	// c = "ab"
	println(a, b, c, d, e)
}
