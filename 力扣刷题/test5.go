package main

import (
	"fmt"
	"time"
)

var echo = func(in ...interface{}) {
	for _, v := range in {
		switch v.(type) {
		case byte:
			newV, _ := v.(byte)
			fmt.Printf("%s ", string(newV))
		case rune:
			newV, _ := v.(rune)
			fmt.Printf("%s ", string(newV))
		default:
			fmt.Print(v, " ")
		}
	}
	fmt.Println()
}

func main() {
	defer func(start time.Time) {
		terminal := time.Since(start)
		fmt.Println(terminal)
	}(time.Now())

	var res string
	res = longestPalindrome("aab")
	fmt.Println("结果:", res) // aaaa
	// fmt.Println()
	// fmt.Println()
	// res = longestPalindrome("oaaoaao")
	// fmt.Println("结果:", res)
}

func longestPalindrome(s string) string {
	var n = len(s)
	if n == 1 {
		return s
	}
	if n == 2 {
		if s[0] == s[1] {
			return s
		}
	}
	var start, end, left, right, i int
	for j := 0; j < n; j++ {

	}
	for pivot := 0; pivot < n; pivot++ {
		for i = pivot; i+1 < n; i++ {
			if s[i] != s[i+1] {
				break
			}
		}
		if i-pivot > n/2 {
			return over(s, pivot, i)
		}
		if i-pivot > end-start {
			start, end = pivot, i
		}

		if (n-1-pivot)*2 <= end-start {
			if start == end {
				return s[:1]
			}
			return s[start : end+1]
		}
		left, right = pivot-1, pivot+1
		if left >= 0 && right < n {
			if a, b := maxLengthIndex(s, left, right); b-a > end-start {
				if b-a == n-1 {
					return over(s, a, b)
				}
				start, end = a, b
			}
			if s[right] == s[pivot] {
				left = pivot
				if a, b := maxLengthIndex(s, left, right); b-a > end-start {
					if b-a == n-1 {
						return over(s, a, b)
					}
					start, end = a, b
				}
			}
			if s[left] == s[pivot] && pivot == 1 {
				right = pivot
				if a, b := maxLengthIndex(s, left, right); b-a > end-start {
					if b-a == n-1 {
						return over(s, a, b)
					}
					start, end = a, b
				}
			}
		}
	}
	return over(s, start, end)
}

func over(s string, start, end int) string {
	if start == end {
		return s[:1]
	}
	return s[start : end+1]
}

func maxLengthIndex(s string, left, right int) (start, end int) {
	var n = len(s)
	for left >= 0 && right < n {
		if s[left] == s[right] {
			if right-left > end-start {
				start, end = left, right
			}
		} else {
			return
		}
		left--
		right++
	}
	return
}
