package main

import "fmt"

func fib(n uint) uint {
	if n == 0 {
		fmt.Println("0 반환합니다")
		return 0
	} else if n == 1 {
		fmt.Println("1 반환합니다.")
		return 1
	} else {
		fmt.Println("엘스", n)
		return fib(n-1) + fib(n-2)
	}
}

func main() {
	for i := 0; i < 11; i++ {
		n := uint(i)
		fmt.Println(n, "일때 값", fib(n))
		fmt.Println("------------------------------------------")
	}
}
