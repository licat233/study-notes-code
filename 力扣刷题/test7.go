package main

import (
	"bytes"
	"container/list"
	"fmt"
	"math"
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
	var data int = -153400
	var res int
	res = reverse(data)
	fmt.Println(res) //9
}

func reverse(x int) int {
	var n int
	for x != 0 {
		n = n*10 + x%10
		x /= 10
	}
	if n < -2147483648 || n > 2147483647 {
		return 0
	}
	return n
}

func reverse2(x int) int {
	var build bytes.Buffer
	var s = fmt.Sprint(x)
	var n = len(s)
	var i int
	if x < 0 {
		build.WriteString("-")
		i = 1
	}
	var stack = list.New()
	for i < n {
		stack.PushBack(s[i])
		i++
	}
	for stack.Len() != 0 {
		last := stack.Back()
		build.WriteByte(last.Value.(byte))
		stack.Remove(last)
	}
	x, _ = strconv.Atoi(build.String())
	if x < -2147483648 || x > 2147483647 {
		return 0
	}
	return x
}

func reverse1(x int) int {
	var isFu = false
	if x < 0 {
		// x = int(math.Abs(float64(x)))
		x = x * -1
		isFu = true
	}
	s := fmt.Sprint(x)
	n := len(s)
	var c = make([]byte, n)
	for i := 0; i < n; i++ {
		c[i] = s[i]
	}
	isEven := n%2 != 0
	var mid int
	if isEven {
		mid = (n - 1) / 2
		reverseHandler(c, mid-1, mid+1)
	} else {
		mid = n/2 - 1
		reverseHandler(c, mid, mid+1)
	}
	s = string(c)
	x, _ = strconv.Atoi(s)
	if isFu {
		if int(math.Pow(2, 31)) < x {
			return 0
		}
		x = x * -1
		return x
	}
	if int(math.Pow(2, 31)-1) < x {
		return 0
	}
	return x
}

func reverseHandler(c []byte, left, right int) {
	if left < 0 || right >= len(c) {
		return
	}
	c[left], c[right] = c[right], c[left]
	reverseHandler(c, left-1, right+1)
}
