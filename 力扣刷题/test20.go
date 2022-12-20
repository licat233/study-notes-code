package main

import (
	"fmt"
	"strings"
	"time"
)

func echo(a ...interface{}) (int, error) {
	return fmt.Println(a...)
}

func performance(s string) {
	defer func(start time.Time) {
		terminal := time.Since(start)
		fmt.Println(terminal)
	}(time.Now())
	echo(isValid(s))
}

func main() {
	// performance("()")
	// performance("()[]{}")
	// performance("(]{{((")
	// performance("([)]")
	// performance("{[]}")
	// performance("({[)")
	performance("(([]){})")
	// echo(byte('('), byte(')'))
	// echo(byte('{'), byte('}'))
	// echo(byte('['), byte(']'))
	// var s = "(([]){})()"
	// reg := regexp.MustCompile("(\\(\\))|(\\{\\})|(\\[\\])+")
	// sb := []byte(s)
	// space := []byte{}
	// sb = reg.ReplaceAll(sb, space)
	// echo(sb)
}
func isValid(str string) bool {
	n := len(str)
	if n%2 != 0 {
		return false
	}
	for i, s := 0, []byte(str); i+2 <= n; {
		if d := s[i+1] - s[i]; d == 2 || d == 1 {
			s = append(s[:i], s[i+2:]...)
			if n = len(s); i != 0 {
				i--
			}
		} else {
			i++
		}
	}
	return n == 0
}

//思路，找到最小单位的括号，然后将其删除
func isValid2(s string) bool {
	var n = len(s)
	if n%2 != 0 {
		return false
	}
	count := n / 2
	for i := 0; i < count; i++ {
		s = strings.ReplaceAll(s, "()", "")
		s = strings.ReplaceAll(s, "[]", "")
		s = strings.ReplaceAll(s, "{}", "")
	}
	return n == 0
}
