package main

import (
	"fmt"
	"math"
)

func reverse_int(n int) int {
	new_int := 0
	for n > 0 {
		remainder := n % 10
		new_int *= 10
		new_int += remainder
		n /= 10
	}
	return new_int
}

func prime_number(num int) bool {
	sq_root := int(math.Sqrt(float64(num)))
	for i := 2; i <= sq_root; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func distinct_digits(i int) bool {
	var digits [10]bool

	digits[0] = true

	for i > 0 {
		var dig = i % 10
		if digits[dig] {
			return false
		}
		digits[dig] = true
		i = i / 10
	}
	return true
}

func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func main() {
	var limit = powInt(10, 9)
	for i := 1; i < limit; i++ {
		//		fmt.Printf("Checking %d\n", i)
		if distinct_digits(i) && prime_number(i) {
			var r = reverse_int(i)
			//			fmt.Printf("i: %d r: %d\n", i, r)
			if prime_number(r) {
				fmt.Printf("==> i: %d, r: %d\n", i, r)
			}
		}
	}
}
