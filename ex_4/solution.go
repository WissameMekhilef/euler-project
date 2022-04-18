package ex_4

import "fmt"

// A palindromic number reads the same both ways.
// The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
// Find the largest palindrome made from the product of two 3-digit numbers.

func Run() {
	var largestPalindrome int
	for i := 0; i <= 999; i++ {
		for j := 0; j <= 999; j++ {
			m := i * j
			if isPalindrome(m) && largestPalindrome < m {
				largestPalindrome = m
			}
		}
	}
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
