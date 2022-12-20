package main

import (
	"fmt"
	"regexp"
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
	// echo("+:", byte(43), "-:", byte(45), "8:", byte(56))
	echo(myAtoi("  00000-42a1234")) // -2147
}

func myAtoi2(s string) int {
	reg := regexp.MustCompile(`^\s*[\+\-]?\d+`)
	str := reg.FindString(s)
	reg1 := regexp.MustCompile(`[\+\-]?\d+`)
	str1 := reg1.FindString(str)
	x, _ := strconv.Atoi(str1)
	if x < -2147483648 {
		return -2147483648
	} else if x > 2147483647 {
		return 2147483647
	} else {
		return x
	}
}

func myAtoi(s string) int {
	var n = len(s)
	var readNumber = true
	var start, i int
	for ; i < n; i++ {
		if s[i] == byte(32) && readNumber {
			continue
		}
		if ((s[i] == byte(43) || s[i] == byte(45)) || (byte(48) <= s[i] && s[i] <= byte(57))) && readNumber {
			readNumber = false
			start = i
			continue
		}
		if s[i] < byte(48) || s[i] > byte(57) {
			break
		}
	}
	x, _ := strconv.Atoi(s[start:i])
	if x < -2147483648 {
		return -2147483648
	} else if x > 2147483647 {
		return 2147483647
	} else {
		return int(x)
	}
}
