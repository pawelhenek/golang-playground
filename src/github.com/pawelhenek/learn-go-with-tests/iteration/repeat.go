package main

import "fmt"

func Repeat(char string, no int) string {
	var repeated string
	for i := 0; i < no; i++ {
		repeated += char
	}
	return repeated
}

func main() {
	fmt.Println(Repeat("abc", 7))
}
