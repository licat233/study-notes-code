package main

import (
	"fmt"
	"strconv"
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
	echo(isPalindrome4(10))
}

func isPalindrome4(x int) bool {
	s := strconv.Itoa(x)
	for s, left, right := strconv.Itoa(x), 0, len(s)-1; right-left >= 1; left, right = left+1, right-1 {
		if s[left] != s[right] {
			return false
		}
	}
	return true
}

func isPalindrome2(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}

	s := strconv.Itoa(x)
	n := len(s)
	if x < 1000 {
		if s[0] != s[len(s)-1] {
			return false
		}
		return true
	}

	var left, right int
	if n%2 != 0 {
		left = (n-1)/2 - 1
		right = left + 2
	} else {
		left = n/2 - 1
		right = left + 1
	}

	var left2, right2 int = 0, n - 1
	for i := 0; right+i < n && left-i >= 0; i++ {
		if s[left2+i] != s[right2-i] || s[left-i] != s[right+i] {
			return false
		}
	}
	return true
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	s := strconv.Itoa(x)
	n := len(s)
	isEven := n%2 != 0
	if isEven {
		return isEqual(s, (n-1)/2-1, (n-1)/2+1)
	}
	return isEqual(s, n/2-1, n/2)
}

func isEqual(s string, left, right int) bool {
	if left < 0 {
		return true
	}
	if s[left] == s[right] {
		return isEqual(s, left-1, right+1)
	}
	return false
}

func isPalindrome3(x int) bool {
	if x < 0 {
		return false
	}
	s := strconv.Itoa(x)
	n := len(s)
	var left, right int
	if n%2 != 0 {
		left = (n-1)/2 - 1
		right = left + 2
	} else {
		left = n/2 - 1
		right = left + 1
	}

	for i := 0; right+i < n; i++ {
		if s[left-i] != s[right+i] {
			return false
		}
	}
	return true
}
