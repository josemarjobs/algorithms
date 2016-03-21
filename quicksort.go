package main

import "fmt"

func main() {
	numbers := []int{9, 7, 5, 4, 5, 6, 3}
	sort(numbers, 0, len(numbers)-1)
	fmt.Println(numbers)
	fmt.Println(".....................")

	numbers = []int{9, 2, 5}
	sort(numbers, 0, len(numbers)-1)
	fmt.Println(numbers)
	fmt.Println(".....................")

	numbers = []int{9, 7}
	sort(numbers, 0, len(numbers)-1)
	fmt.Println(numbers)
	fmt.Println(".....................")
}

func sort(numbers []int, left, right int) {
	pivot := numbers[(left+right)/2]
	i, j := left, right
	for i <= j {
		for numbers[i] < pivot {
			i++
		}
		for numbers[j] > pivot {
			j--
		}
		if i <= j {
			numbers[i], numbers[j] = numbers[j], numbers[i]
			i++
			j--
		}
	}
	if left < j {
		sort(numbers, left, j)
	}
	if i < right {
		sort(numbers, i, right)
	}
}
