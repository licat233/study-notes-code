package main

import (
	"fmt"
	"time"
)

func echo(a ...interface{}) (int, error) {
	return fmt.Println(a...)
}

type Fn func(strs []string) string

func performance(data []string, fn Fn) {
	defer func(start time.Time) {
		terminal := time.Since(start)
		fmt.Println(terminal)
	}(time.Now())
	echo(fn(data))
}

func main() {
	// performance([]string{"flower", "flow", "flight"}, longestCommonPrefix)
	performance([]string{"a", "abc"}, longestCommonPrefix)
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 || strs[0] == "" {
		return strs[0]
	}
	var k int
	for n, j := len(strs)-1, 1; k < len(strs[0]) && k < len(strs[j]) && strs[0][k] == strs[j][k]; j++ {
		if j == n {
			j = 0
			k++
		}
	}
	return strs[0][:k]
}
