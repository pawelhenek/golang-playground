package main

import "fmt"

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTraits(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}

func main() {
	fmt.Println(fmt.Sprintf("1 + 3 + 6 + 8 = %d", Sum([]int{1, 3, 6, 8})))
	fmt.Println(fmt.Sprintf("1 + 3 + 6 + 8 = %d", SumAll([]int{1, 3, 6, 8})))
	fmt.Println(fmt.Sprintf("1 + 3 + 6 + 8 = %d", SumAllTraits([]int{1, 3, 6, 8})))
}
