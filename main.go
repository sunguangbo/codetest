package main

import (
	"code_test/array"
	"fmt"
)

func main() {
	s := []int{8, 8}
	fmt.Println(array.MinSubArrayLen(s, 7))
}
