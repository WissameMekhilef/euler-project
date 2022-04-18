package ex_3

import "fmt"

// The prime factors of 13195 are 5, 7, 13 and 29.
// What is the largest prime factor of the number 600851475143 ?

func Run() {
	// var n int64 = 13195
	var n int64 = 600851475143
	factors := primeFactorsOf(n)
	// fmt.Printf("Prime factors of %v are %v\n", n, factors, len(factors))
	fmt.Println(factors[len(factors)-1])
	fmt.Println(len(factors))
}

func primeFactorsOf(n int64) []int64 {
	primeFactors := make([]int64, 0)
	var c int64 = 2
	for n > 1 {
		if n%c == 0 {
			primeFactors = append(primeFactors, c)
			n /= c
		} else {
			c++
		}
	}
	return primeFactors
}
