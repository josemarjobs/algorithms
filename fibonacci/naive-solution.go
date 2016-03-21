package main

import (
	"fmt"
	"os"
)

// do not run this on numbers bigger the 60
func fib(n int) int {
	if n > 50 {
		fmt.Println("Number too big")
		os.Exit(1)
	}
	if n <= 1 {
		return n
	} else {
		return fib(n-1) + fib(n-2)
	}
}
func main() {
	for i := 0; i < 20; i++ {
		fmt.Println(i, ":", fib(i), " - ", T(i))
	}
}

// Number os steps
func T(n int) int {
	if n <= 1 {
		return 2
	} else {
		return (T(n-1) + T(n-2)) + 3
	}
}

// 0 1 1 2 3 5 8 13
