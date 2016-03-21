package main

import (
	"fmt"
	"os"
)

func gcd(a, b int) (best int) {
	// just to be safe
	if a > 10000000 || b > 10000000 {
		fmt.Println("Numbers too big")
		os.Exit(1)
	}

	for i := 1; i <= (a + b); i++ {
		if a%i == 0 && b%i == 0 {
			best = i
		}
	}
	return
}

func main() {
	fmt.Println(gcd(100, 4))
	fmt.Println(gcd(134352, 41223))
}
