package ex_4

import (
	"fmt"
	"sync"
)

// A palindromic number reads the same both ways.
// The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
// Find the largest palindrome made from the product of two 3-digit numbers.

type Ex4 struct{}

func NewEx4() *Ex4 {
	return &Ex4{}
}

func (sol *Ex4) Run() {
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
	var lock = sync.Mutex{}
	ch := make(chan int, 10000)
	wg.Add(1)
	go func(ch <-chan int) {
		for m := range ch {
			wg.Add(1)
			go func(m int) {
				if isPalindrome(m) {
					// Lock the ressource to see if the new number needs to replace the current largest
					lock.Lock()
					isGreater := m > largestPalindrome
					if isGreater {
						largestPalindrome = m
					}
					lock.Unlock()
				}
				wg.Done()
			}(m)
		}
		wg.Done()
	}(ch)
	wg.Add(1)
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
