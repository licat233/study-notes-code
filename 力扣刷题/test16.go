package main

import (
	"fmt"
	"sort"
	"time"
)

func echo(a ...interface{}) (int, error) {
	return fmt.Println(a...)
}

type Fn func(nums []int, target int) int

func performance(nums []int, target int, fn Fn) {
	defer func(start time.Time) {
		terminal := time.Since(start)
		fmt.Println(terminal)
	}(time.Now())
	echo(fn(nums, target))
}

func main() {
	data := []int{1, 2, 4, 8, 16, 32, 64, 128}
	target := 82
	performance(data, target, threeSumClosest) //82
}

func threeSumClosest(nums []int, target int) (res int) {
	sort.Ints(nums)
	n := len(nums)
	res = nums[0] + nums[1] + nums[2]
	var minDiff = res - target
	if minDiff == 0 || n == 3 {
		return
	} else if minDiff < 0 {
		minDiff = -1 * minDiff
	}
	var a, b, c, left, diff, sum int
	for maxA := n - 3; a <= maxA; a++ {
		left = nums[a]
		if left+nums[a+1]+nums[a+2]-target >= minDiff {
			return
		}
		//过滤a指针重复值
		if a != 0 && left == nums[a-1] {
			continue
		}
		for b, c = a+1, n-1; b < c; {
			if sum = left + nums[b] + nums[c]; sum < target { //如果sum比target小，则需要变大，移动 b 指针，c不变
				//过滤b指针重复值
				for b++; b > a+1 && nums[b] == nums[b-1] && b < c; b++ {
				}
			} else if sum > target { //如果sum 比 target大，则需要变小，移动 c 指针
				//过滤c指针重复值
				for c--; c < n-1 && nums[c] == nums[c+1] && c > b; c-- {
				}
			} else { //如果相等，则直接return
				res = sum
				return
			}
			if diff = sum - target; diff < 0 {
				//取绝对值
				diff = -1 * diff
			}
			if diff < minDiff {
				res = sum
				minDiff = diff
			}
		}
	}
	return
}
