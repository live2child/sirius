package main

import "fmt"

func main() {

	var num int = 101
	var n, n1, n2 uint // F(n) = F(n-1)+(n-2)

	for i := 0; i < num; i++ {
		if i < 2 {
			fmt.Println(i)
			n1 = 0 // Ifn = 0
			n2 = 1 // Ifn = 1
		} else {
			n = n1 + n2
			fmt.Println(n)
			n1 = n2
			n2 = n
		}
	}
}
