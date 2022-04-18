package ex_2

import "fmt"

const limit = 4000000

func Run() {
	cache := make([]int64, limit+1)

	var result int64
	var i int64
	for {
		fibo := fibo(i, cache)
		if fibo > limit {
			break
		}
		result += fibo

		i++
	}
	fmt.Println(result)
}

func fibo(n int64, cache []int64) int64 {
	if n < 2 {
		return n
	}
	if cache[n] != 0 {
		return cache[n]
	}
	res := fibo(n-1, cache) + fibo(n-2, cache)
	cache[n] = res
	return res
}
