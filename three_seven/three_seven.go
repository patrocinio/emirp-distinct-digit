package main

import (
	"fmt"
	"math"
)

func prime_number(num uint64) bool {
	sq_root := uint64(math.Sqrt(float64(num)))
	for i := uint64(2); i <= sq_root; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func add(num uint64, digit uint64) {
	var i = num*10 + digit
	if i%10 != 3 && i%10 != 7 {
		return
	}
	//	fmt.Printf("add i: %d\n", i)
	if prime_number(i) {
		fmt.Printf("==> i: %d\n", i)
	}
	run(i)
}

func run(num uint64) {
	//	fmt.Printf("run num: %d\n", num)
	add(num, 3)
	add(num, 7)
}

func main() {
	run(0)
}
