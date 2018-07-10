package main

import "fmt"

func main() {
	var p int
	var n1 int = 1
	var n2 int = 1
	fmt.Println(0)
	fmt.Println(1)
	fmt.Println(1)
	var n int = 110
	for i := 0; i < n; i++ {
		p = n1 + n2
		n1 = n2
		n2 = p
		fmt.Println(p)
	}
}
