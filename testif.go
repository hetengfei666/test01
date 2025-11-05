package main

import (
	"fmt"
)

//	func main() {
//		a := 6
//		if a == 5 {
//			fmt.Printf("今天星期五")
//		} else if a == 6 {
//			fmt.Printf("今天星期六")
//		} else {
//			fmt.Printf("不知道")
//		}
//	}
func main() {
	a := 6
	switch a {
	case 1:
		fmt.Printf("星期1")
	case 2:
		fmt.Printf("星期2")
	case 3:
		fmt.Printf("星期3")
	default:
		fmt.Printf("不知道")
	}
}
