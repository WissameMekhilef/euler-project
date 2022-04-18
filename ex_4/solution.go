package ex_4

import (
	"fmt"
	"sync"
)

// A palindromic number reads the same both ways.
// The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
// Find the largest palindrome made from the product of two 3-digit numbers.

func Run() {
	var largestPalindrome int
	// Naive version
	// for i := 0; i <= 999; i++ {
	// 	for j := 0; j <= 999; j++ {
	// 		m := i * j
	// 		if isPalindrome(m) && largestPalindrome < m {
	// 			largestPalindrome = m
	// 		}
	// 	}
	// }

	// go routine with channel
	var wg = sync.WaitGroup{}
	var lock = sync.RWMutex{}
	ch := make(chan int, 1000)
	wg.Add(2)
	go func(ch <-chan int) {
		for m := range ch {
			go func(m int) {
				lock.RLock()
				isGreaterPalindrome := isPalindrome(m) && largestPalindrome < m
				lock.RUnlock()
				if isGreaterPalindrome {
					lock.Lock()
					largestPalindrome = m
					lock.Unlock()
				}
			}(m)
		}
		wg.Done()
	}(ch)

	go func(ch chan<- int) {
		for i := 0; i <= 999; i++ {
			for j := 0; j <= 999; j++ {
				ch <- i * j
			}
		}
		close(ch)
		wg.Done()
	}(ch)
	wg.Wait()

	fmt.Println(largestPalindrome)
}

func isPalindrome(n int) bool {
	if n < 0 || (n%10 == 0 && n != 0) {
		return false
	}
	var revertedNumber int
	for n > revertedNumber {
		revertedNumber = revertedNumber*10 + n%10
		n /= 10
	}
	return n == revertedNumber || n == revertedNumber/10
}
