package main

import (
	"fmt"
	"time"
)

const size = 100000

var s []int = sieve_gen(size)

func next_prime() func() int {
	x := 0
	return func() int {
		if x == 0 || x == 1 {
			x = 2
			return x
		} else {
			for i := x + 1; i < size; i++ {
				if s[i] == 0 {
					x = i
					return x
				}
			}
			return 0
		}
	}
}

func num_divisors(n int) int {
	var current_prime int
	var exp int
	prime := next_prime()
	current_prime = prime()
	divisors := 1
	copy_n := n

	for current_prime < n && current_prime != 0 {
		for {
			if copy_n%current_prime == 0 {
				exp += 1
				copy_n /= current_prime
			} else {
				divisors *= (exp + 1)
				current_prime = prime()
				exp = 0
				break
			}
		}
	}
	return divisors
}

func main() {
	start := time.Now()
	var div int
	tri_i := 1
	step := 2

	for div < 500 {
		tri_i += step
		step += 1
		div = num_divisors(tri_i)
	}
	fmt.Println(tri_i)
	fmt.Println("Took ", time.Since(start), " seconds")
}
