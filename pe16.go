package main

import "fmt"
import "math"

func power_digit_sum(x int, y int) int {
	var x_arr []int = toArr(x)
	var total []int = make([]int, len(x_arr))
	total = addArr(total, x_arr)
	for i := 2; i <= y; i++ {
		total = addArr(total, total)
	}
	return sum_over(total)
}

func toArr(x int) []int {
	var str string = string(x)
	var length int = len(str)
	arr := make([]int, length)
	for i := length - 1; i >= 0; i-- {
		arr[i] = int(str[i])
	}
	return arr
}

func addArr(x []int, y []int) []int {
	x_len := len(x)
	y_len := len(y)
	greater_len := int(math.Max(float64(x_len), float64(y_len)))
	lesser_len := int(math.Min(float64(x_len), float64(y_len)))
	sum := make([]int, greater_len)
	for i := 0; i < lesser_len; i++ {
		sum[i] = x[i] + y[i]
	}
	sum = append(sum[:lesser_len], x[lesser_len:]...)
	sum = carried(sum)
	return sum
}

func carried(x []int) []int {
	var car int
	length := len(x)
	for i := 0; i < length; i++ {
		if car == 1 {
			x[i] += 1
			car = 0
		}
		if x[i] >= 10 {
			x[i] = x[i] % 10
			car = 1
		}
	}
	if car == 1 {
		temp_len := len(x) + 1
		ret := make([]int, temp_len)
		copy(ret, x)
		ret[temp_len-1] = 1
		return ret
	} else {
		return x
	}
}

func sum_over(x []int) int {
	length := len(x)
	var sum int
	for i := 0; i < length; i++ {
		sum += x[i]
	}
	return sum
}

func main() {
	sum := power_digit_sum(2, 1000)
	fmt.Println(sum)
}
