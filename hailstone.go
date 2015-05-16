package main

import "fmt"
import "time"

//memoization ayy lmao
var hailstone_len_map map[int]int = make(map[int]int)

func hailstone_len(n int) int {
	length := 0
	if n == 1 {
		return 1
	}
	i, ok := hailstone_len_map[n]
	if !ok {
		if n%2 == 0 {
			length = 1 + hailstone_len(n/2)
		} else {
			length = 1 + hailstone_len(3*n+1)
		}
		hailstone_len_map[n] = length
		return length
	}
	return i
}

func max_hailstone_len(limit int) int {
	start := time.Now()
	defer fmt.Println(time.Since(start))
	var max int
	var max_i int
	var curr_len int
	for i := 1; i < limit; i++ {
		curr_len = hailstone_len(i)
		if curr_len > max {
			max = curr_len
			max_i = i
		}
	}
	return max_i
}

func main() {
	fmt.Println(max_hailstone_len(1000000))
}
