package main

import (
	"fmt"
	"time"
)

func echo(a ...interface{}) (int, error) {
	return fmt.Println(a...)
}

type Fn func(s string) int

func performance(data string, fn Fn) {
	defer func(start time.Time) {
		terminal := time.Since(start)
		fmt.Println(terminal)
	}(time.Now())
	echo(fn(data))
}

func main() {
	performance("MCMXCIV", romanToInt)  //1994
	performance("MCDLXXVI", romanToInt) //1476
	performance("MMCDXXV", romanToInt)  //2425
	performance("CMLII", romanToInt)    //952
	performance("MMMCDXII", romanToInt) //3412
	performance("MMMCDIII", romanToInt) //3403
	performance("MMMCXCI", romanToInt)  //3191
	performance("IV", romanToInt)       //4
}
func romanToInt(s string) int {
	var m = map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	var res, i, last, current int
	for i = len(s) - 1; i >= 0; i-- {
		if i < 1 {
			return res + m[string(s[i])]
		}
		last, current = m[string(s[i-1])], m[string(s[i])]
		if res += current; last == current {
			res += last
			i--
		} else if last < current {
			res -= last
			i--
		}
	}
	return res
}
func romanToInt4(s string) int {
	var m = map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	var res, i, last, current int
	for i = len(s) - 1; i >= 0; i-- {
		last, current = m[string(s[i-1])], m[string(s[i])]
		if res += current; last == current {
			res += last
			i--
		} else if last < current {
			res -= last
			i--
		}
	}
	if i == 0 {
		return res + m[string(s[i])]
	}
	return res
}
func romanToInt3(s string) int {
	var m = map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	var res, last, current int
	i := len(s) - 1
	for i > 0 {
		last, current = m[string(s[i-1])], m[string(s[i])]
		if last > current {
			res += current
			i--
			continue
		} else if last == current {
			res += last + current
		} else {
			res += current - last
		}
		i -= 2
	}
	if i == 0 {
		return res + m[string(s[i])]
	}
	return res
}
func romanToInt2(s string) int {
	var m = map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	var res, last, current int
	for i := len(s) - 1; i >= 0; {
		if i-1 < 0 {
			res += m[string(s[i])]
			break
		}
		last, current = m[string(s[i-1])], m[string(s[i])]
		if last > current { //LX
			res += current
			i--
			continue
		} else if last == current {
			res += last + current
		} else { //IX
			res += current - last
		}
		i -= 2
	}
	return res
}

func romanToIntOLD(s string) int {
	//创建一个罗马数字符库
	var m = map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000, "LV": -1, "LI": -1, "XV": -1, "CV": -1, "CI": -1, "DL": -1, "DX": -1, "DI": -1, "CL": -1, "MX": -1, "ML": -1, "MD": -1}
	var res int
	for i := len(s) - 1; i >= 0; {
		current := string(s[i])
		if i-1 < 0 { //已经是第一个了,XI
			res += m[current]
			break
		}
		next := string(s[i-1])
		if m[next] >= m[current] { //LXX
			if m[next+current] == -1 { //避免DL、CL、ML这种情况
				res += m[current]
				i--
				continue
			}
			res += m[current] + m[next]
			i -= 2
		} else { //IX
			res += m[current] - m[next]
			i -= 2
		}
	}
	return res
}
