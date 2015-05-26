package main

import (
	"fmt"
)

var num_to_word = map[int]string{
	1:    "one",
	2:    "two",
	3:    "three",
	4:    "four",
	5:    "five",
	6:    "six",
	7:    "seven",
	8:    "eight",
	9:    "nine",
	10:   "ten",
	11:   "eleven",
	12:   "twelve",
	13:   "thirteen",
	14:   "fourteen",
	15:   "fifteen",
	16:   "sixteen",
	17:   "seventeen",
	18:   "eighteen",
	19:   "nineteen",
	20:   "twenty",
	30:   "thirty",
	40:   "forty",
	50:   "fifty",
	60:   "sixty",
	70:   "seventy",
	80:   "eighty",
	90:   "ninety",
	100:  "hundred",
	1000: "thousand",
}

func stringify(num []int) int {
	length := len(num)
	if length == 1 {
		return len(num_to_word[num[0]])
	} else if length == 2 {
		if num[0] == 0 && num[1] == 0 {
			return 0
		} else if num[1] == 0 {
			return len(num_to_word[num[0]*10])
		} else if num[0] == 1 {
			return len(num_to_word[num[0]*10+num[1]])
		} else {
			return len(num_to_word[num[0]*10] + num_to_word[num[1]])
		}
	} else if len(num) == 3 {
		x := len(num_to_word[num[0]] + num_to_word[100])
		y := stringify(num[1:])
		if y == 0 {
			return x
		} else {
			return x + len("and") + y
		}
	} else {
		return len("onethousand")
	}
}

func arrayify(num int) []int {
	numarr := make([]int, 0, 4)
	var f func(int)
	f = func(num2 int) {
		if num2 < 10 {
			numarr = append(numarr, num2)
		} else {
			dig := num2 % 10
			f(num2 / 10)
			numarr = append(numarr, dig)
		}
	}
	f(num)
	return numarr
}

func main() {
	var char_len int
	for i := 1; i <= 1000; i++ {
		a := arrayify(i)
		char_len += stringify(a)
	}
	fmt.Println(char_len)
}
