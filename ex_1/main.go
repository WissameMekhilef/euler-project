package main

import "fmt"

func main() {
	x := getMultiplesOfUnder(3, 5, 10000)
	fmt.Println(x)
}

func getMultiplesOfUnder(a1 int, a2 int, b int) int {
	result := 0
	for i := 1; i <= b; i++ {
		if i%a1 == 0 || i%a2 == 0 {
			result += 1
		}

	}
	return result
}
