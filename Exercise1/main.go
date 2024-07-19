package main

import (
	"fmt"
)

func main() {
	var nums = []int{1, 2, 3, 4, 5}
	var sum int
	for _, num := range nums {
		sum += num
	}
	fmt.Println(sum)
}