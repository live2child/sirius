package main

import "fmt"

func fib(n uint) uint {
	if n < 2 {
		return n
	} else {
		return fib(n-1) + fib(n-2)
	}
}

func main() {
	fmt.Println(fib(3))
}
