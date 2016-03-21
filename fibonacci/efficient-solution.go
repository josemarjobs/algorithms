package main

import "fmt"

func fib(n int) int {
	f := []int{0, 1}
	for i := 2; i <= n; i++ {
		f = append(f, f[i-1]+f[i-2])
	}
	return f[n]
}
func main() {
	for i := 0; i < 50; i++ {
		fmt.Println(i, ":", fib(i), " - ", T(i))
	}
}

// Number os steps
func T(n int) int {
	return 2*n + 2
}

// 0 1 1 2 3 5 8 13
