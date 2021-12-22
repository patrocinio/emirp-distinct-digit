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

func add(num int, digit int) {
	var i = num*10 + digit
	fmt.Printf("add i: %d\n", i)
	if prime_number(i) {
		fmt.Printf("==> i: %d\n", i)
	}
	run(i)
}

func run(num int) {
	fmt.Printf("run num: %d\n", num)
	//	add(num, 3)
	//	add(num, 7)
}

func main() {
	//	run(0)
}
