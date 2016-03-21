package main

import "fmt"

func main() {
	numbers := [][]int{
		[]int{9, 7, 5, 4, 5, 6, 3},
		[]int{9, 2, 5},
		[]int{9, 7},
		[]int{9, 7, 5, 3, 100, 123, 543, 56, 7, 8, 5, 5, 4, 5, 6, 3},
	}
	for i := 0; i < len(numbers); i++ {
		sort(numbers[i], 0, len(numbers[i])-1)
	}

	fmt.Println(numbers)

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
