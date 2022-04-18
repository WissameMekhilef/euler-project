package ex_1

import (
	"fmt"
)

const (
	a int32 = 3
	b int32 = 5
)

func Run() {
	x := findAllMultipleUntil(10000)
	fmt.Println(x)
}

func findAllMultipleUntil(limit int32) int {

	var result int
	for i := 0; i < int(limit); i++ {
		if i%int(a) == 0 || i%int(b) == 0 {
			result += i
		}
	}
	return result
}
