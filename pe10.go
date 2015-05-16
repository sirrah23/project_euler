package main

/*
import (
	"fmt"
)
*/

func sieve_gen(primes int) []int {
	size := primes
	sieve := make([]int, size)
	for i := 2; i < size; i++ {
		pos := i
		j := 2
		newpos := pos * j
		for newpos < size {
			sieve[newpos] = 1
			newpos = pos * j
			j += 1
		}
	}
	return sieve
}

/*
func main() {
	var sum int
	s := sieve_gen(2000000)
	for i, val := range s {
		if val != 1 {
			sum += i
		}
	}
	sum -= 1
	fmt.Println(sum)
}
*/
