package main

import (
	"fmt"
	"math"
)

func prime_number(num int) bool {
	sq_root := int(math.Sqrt(float64(num)))
	for i := 2; i <= sq_root; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func append(base int, i int) {
	var n = base*10 + i
	if prime_number(n) {
		fmt.Printf("%d --> %d\n", base, n)
		prime_tree(n)
	}

}

func digit_suffix(base int) {
	for i := 1; i < 10; i++ {
		append(base, i)
	}
}

func prime_tree(i int) {
	digit_suffix(i)
}

func main() {
	prime_tree(17)
}
