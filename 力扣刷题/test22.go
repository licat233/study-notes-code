package main

import (
	"fmt"
	"time"
)

func echo(a ...interface{}) (int, error) {
	return fmt.Println(a...)
}

func performance(n int) {
	defer func(start time.Time) {
		terminal := time.Since(start)
		fmt.Println(terminal)
	}(time.Now())
	echo(generateParenthesis(n))
}

func main() {
	performance(3)
	performance(1)
}

func generateParenthesis(n int) (res []string) {
	base := map[string]bool{"()": true}
	for i := 1; i < n; i++ {
		newBase := map[string]bool{}
		for k := range base {
			for index := range k {
				newBase[k[:index]+"()"+k[index:]] = true
			}
		}
		base = newBase
	}
	for k := range base {
		res = append(res, k)
	}
	return
}
