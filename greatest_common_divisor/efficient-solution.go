package main

import "fmt"

// Euclidean Algorithm
// a' = a%b; gcd(a, b) == gcd(a', b) == gcd(b, a')
// q = a/b
// a = a' + b*q
func euclidGCD(a, b int) (best int) {
	if b == 0 {
		return a
	}
	rem_a := a % b
	return euclidGCD(b, rem_a)
}

func main() {
	fmt.Println(euclidGCD(13452, 4123))
	fmt.Println(euclidGCD(357, 234))
	fmt.Println(euclidGCD(1234534524232, 5234412323))
}
