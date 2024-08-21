package main

import "fmt"

func Add(p1, p2 int) int {
	return p1 + p2
}

func main() {
	fmt.Println(fmt.Sprintf("1 + 2 = %d", Add(1, 2)))
	fmt.Println(fmt.Sprintf("3 + 4 = %d", Add(3, 4)))
}
