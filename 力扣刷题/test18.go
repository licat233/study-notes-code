package main

import (
	"fmt"
	"sort"
	"time"
)

func echo(a ...interface{}) (int, error) {
	return fmt.Println(a...)
}
func performance(nums []int, target int) {
	defer func(start time.Time) {
		terminal := time.Since(start)
		fmt.Println(terminal)
	}(time.Now())
	echo(fourSum(nums, target))
}

func main() {
	performance([]int{1, 0, -1, 0, -2, 2}, 0)               //[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
	performance([]int{-2, -1, -1, 1, 1, 2, 2}, 0)           //[[-2,-1,1,2],[-1,-1,1,1]]
	performance([]int{-1, 0, -5, -2, -2, -4, 0, 1, -2}, -9) //[[-5,-4,-1,1],[-5,-4,0,0],[-5,-2,-2,0],[-4,-2,-2,-1]]
}

func fourSum(nums []int, target int) (res [][]int) {
	sort.Ints(nums)
	n := len(nums)
	//特殊情况
	if n < 4 {
		return
	}
	maxCD := nums[n-2] + nums[n-1]
	maxBCD := nums[n-3] + nums[n-2] + nums[n-1]
	//采用四指针方式，a,b,c,d
	//a,b先不动，由c,d先动
	var a, b, c, d, av, bv, cv, dv, sum int
	var lasta, lastb, lastd = n - 4, n - 3, n - 1
	for a = 0; a <= lasta; a++ {
		//特殊情况
		// 提升速度的核心部分
		if nums[a]+nums[a+1]+nums[a+2]+nums[a+3] > target {
			return
		}
		av = nums[a]
		//过滤a指针重复值
		// 提升速度的核心部分
		if (a != 0 && av == nums[a-1]) || av+maxBCD < target {
			continue
		}
		for b = a + 1; b <= lastb; b++ {
			bv = nums[b]
			//过滤b指针重复值
			// 提升速度的核心部分
			if (b != a+1 && bv == nums[b-1]) || av+bv+maxCD < target {
				continue
			}
			for c, d = b+1, n-1; c < d; {
				cv, dv = nums[c], nums[d]
				if sum = av + bv + cv + dv; sum == target {
					res = append(res, []int{av, bv, cv, dv})
					//同时移动c和d
					for c++; c != b+1 && nums[c] == nums[c-1] && c < d; c++ {
					}
					for d--; d != lastd && nums[d] == nums[d+1] && d > c; d-- {
					}
				} else if sum > target {
					//移动d，并且过滤d指针重复值
					for d--; d != lastd && nums[d] == nums[d+1] && d > c; d-- {
					}
				} else if sum < target {
					//移动c，并且过滤c指针重复值
					for c++; c != b+1 && nums[c] == nums[c-1] && c < d; c++ {
					}
				}
			}
		}
	}
	return
}
