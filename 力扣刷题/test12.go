package main

import (
	"fmt"
	"time"
)

func echo(a ...interface{}) (int, error) {
	return fmt.Println(a...)
}

type Fn func(num int) string

func performance(data int, fn Fn) {
	defer func(start time.Time) {
		terminal := time.Since(start)
		fmt.Println(terminal)
	}(time.Now())
	echo(fn(data))
}

func main() {
	// echo(intToRoman(660))
	// performance(10, intToRoman)   //"X"
	// performance(3, intToRoman)    //"III"
	// performance(4, intToRoman)    //"IV"
	// performance(9, intToRoman)    //"IX"
	// performance(58, intToRoman)   //"LVIII"
	performance(1994, intToRoman) //"MCMXCIV"
}
func intToRoman(num int) string {
	var m = []string{"I", "V", "X", "L", "C", "D", "M", ""}
	var res string
	var n int
	for i := 1; num > 0; i++ {
		n = 2*i - 1
		switch num % 10 {
		case 1:
			res = m[n-1] + res
		case 2:
			res = m[n-1] + m[n-1] + res
		case 3:
			res = m[n-1] + m[n-1] + m[n-1] + res
		case 4:
			res = m[n-1] + m[n] + res
		case 5:
			res = m[n] + res
		case 6:
			res = m[n] + m[n-1] + res
		case 7:
			res = m[n] + m[n-1] + m[n-1] + res
		case 8:
			res = m[n] + m[n-1] + m[n-1] + m[n-1] + res
		case 9:
			res = m[n-1] + m[n+1] + res
		}
		num = num / 10
	}
	return res
}
