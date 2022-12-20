package main

import (
	"fmt"
	"time"
)

func echo(a ...interface{}) (int, error) {
	return fmt.Println(a...)
}

type Fn func(height []int) int

func benchmark(data []int, fn Fn) time.Duration {
	start := time.Now()
	fn(data)
	return time.Since(start)
}

func performance(data []int, fn Fn) {
	defer func(start time.Time) {
		terminal := time.Since(start)
		fmt.Println(terminal)
	}(time.Now())
	fn(data)
}

func test(data []int) {
	var result = map[string]int{"maxArea1": 0, "maxArea2": 0, "maxArea3": 0}
	for i := 0; i < 100000; i++ {
		d1 := benchmark(data, maxArea1)
		d2 := benchmark(data, maxArea2)
		d3 := benchmark(data, maxArea3)
		if d1 < d2 && d1 < d3 {
			result["maxArea1"]++
		} else if d2 < d1 && d2 < d3 {
			result["maxArea2"]++
		} else if d3 < d1 && d3 < d2 {
			result["maxArea3"]++
		}
	}
	echo(result)
}

func main() {
	var data2 = []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	performance(data2, maxArea1)
	performance(data2, maxArea2)
	performance(data2, maxArea3)
}

//2.4ms
func maxArea3(height []int) int {
	var max, x, y, res int
	for i, j := 0, len(height)-1; j != i; {
		if j > i {
			x = j - i
		} else {
			x = i - j
		}
		if height[j] > height[i] {
			y = height[i]
			i++
		} else {
			y = height[j]
			j--
		}
		if res = x * y; res > max {
			max = res
		}
	}
	return max
}

//9.1ms
func maxArea2(height []int) int {
	n := len(height) - 1
	var max, x, y, res int
	var MinNum = func(a, b int) int {
		if a < b {
			return a
		} else {
			return b
		}
	}
	for i := 0; i < n; i++ {
		for j := n; j > i; j-- {
			y = MinNum(height[j], height[i])
			x = j - i
			res = x * y
			if res = x * y; res > max {
				max = res
			}
		}
	}
	return max
}

//10.8ms
func maxArea1(height []int) int {
	n := len(height)
	var max, x, y, res int
	var findMax = func(i, j int) int {
		for k := j; k > i; k-- {
			if height[k] > height[j] {
				return k
			}
		}
		return -1
	}
	var MinNum = func(a, b int) int {
		if a < b {
			return a
		} else {
			return b
		}
	}
	for i := 0; i < n-1; i++ {
		for j := n - 1; j > i; {
			y = MinNum(height[j], height[i])
			x = j - i
			if res = x * y; res > max {
				max = res
			}
			if index := findMax(i, j); index != -1 {
				j = index
				continue
			} else {
				break
			}
		}
	}
	return max
}
