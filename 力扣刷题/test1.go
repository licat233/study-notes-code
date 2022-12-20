package main

import (
	"fmt"
	"strings"
	"time"
)

func main1() {
	// res := isMatchTest("mississippi", ".*ip*i")
	// fmt.Println("测试结果:", res)

	defer func(start time.Time) {
		terminal := time.Since(start)
		fmt.Println(terminal)
	}(time.Now())

	res := []int{}
	for i := 0; i < 100000; i++ {
		res = twoSum2([]int{2, 7, 11, 15, 1, 6, 4, 9, 7, 5, 3, 0, 5, 11, 2, 3, 8, 9, 12, 26, 30}, 24)
	}
	fmt.Println(res)
}

func twoSum2(nums []int, target int) []int {
	if len(nums) > 30 {
		var m = make(map[int]int, len(nums))
		for k := 0; k < len(nums); k++ {
			v := nums[k]
			if p, ok := m[target-v]; ok {
				return []int{p, k}
			}
			m[v] = k
		}
	} else {
		for k := 0; k < len(nums); k++ {
			for j := k + 1; j < len(nums); j++ {
				if target-nums[k] == nums[j] {
					return []int{k, j}
				}
			}
		}
	}
	return nil
}

func twoSum(nums []int, target int) []int {
	var m = make(map[int]int, len(nums))
	for k, v := range nums {
		if p, ok := m[target-v]; ok {
			return []int{p, k}
		}
		m[v] = k
	}
	return nil
}

func isMatchTest(s string, p string) bool {
	if len(s) == 0 {
		return false
	}
	if len(p) == 0 {
		return false
	}
	if strings.Index(p, "*") == -1 && strings.Index(p, ".") == -1 {
		if s != p {
			return false
		}
		return true
	}
	x := []byte("*")[0]
	d := []byte(".")[0]
	type T struct {
		Before byte
		Enum   int
		Value  string
	}
	//aa 用 0 表示，a*用1表示，.* 用2表示, . 用 3 表示
	sl := []T{}
	// ps := p
	//取p的两位
	for i := 0; i < len(p); i++ {
		//是 *
		if p[i] == x {
			//前面是 .
			if p[i-1] == d {
				sl = append(sl, T{
					Before: p[i-1],
					Enum:   2,
					Value:  string(p[i-1]),
				})
				continue
			}
			//前面是 字母
			sl = append(sl, T{
				Before: p[i-1],
				Enum:   1,
				Value:  string(p[i-1]),
			})
			continue
		}
		//是.
		if p[i] == d {
			//如果是最后一个
			if i == len(p)-1 {
				sl = append(sl, T{
					Before: p[i],
					Enum:   3,
					Value:  string(p[i]),
				})
				continue
			}
			//如果后面有 *
			if p[i+1] == d {
				continue
			}
			sl = append(sl, T{
				Before: p[i],
				Enum:   3,
				Value:  string(p[i]),
			})
			continue
		}
		//是字母，如果不存在
		if strings.Index(s, string(p[i])) == -1 {
			return false
		}
		//是字母，且是最后一个
		if i == len(p)-1 {
			sl = append(sl, T{
				Before: p[i],
				Enum:   0,
				Value:  string(p[i]),
			})
			continue
		}
		//是字母，但后面一个是 *，不记录
		if p[i+1] == x {
			continue
		}
		//是字母
		sl = append(sl, T{
			Before: p[i],
			Enum:   0,
			Value:  string(p[i]),
		})

	}

	//为了防止死循环的产生
	var count int
	// fmt.Println(sl)
	for i := 0; i < len(s); i++ {
		//如果正则匹配完了，字符串还有剩，则返回false
		if len(sl) == 0 {
			return false
		}
		var k int
		for k < len(sl) && count < 100 {
			count++
			v := sl[k]
			fmt.Printf("匹配【%s】:%v \n", string(s[i]), sl)
			fmt.Printf("字母比较: %s ?? %s \n", string(s[i]), v.Value)
			//sl中存在相同字母
			if v.Before == s[i] {
				//且类型为 0 --  aa，则且结束循环，匹配下一个字母
				if v.Enum == 0 {
					fmt.Println("匹配到了", string(s[i]))
					//防止k等于len时，切割panic
					if k >= len(sl) {
						break
					}
					//排除掉已经匹配过的正则
					sl = sl[k+1:]
					//k值归0
					k = 0
					break
				}
				//且类型为 1 -- a*，则结束循环
				if v.Enum == 1 {
					fmt.Println("匹配到了", string(s[i]))
					// a* 可有可无，不用管
					break
				}
			} else { //如果不相等，则取出第一个sl中的元素看是否为.或者.*
				if v.Enum == 2 {
					//判断是不是最后一个 .*
					if k+1 == len(sl) {
						return true
					}
					//判断 .*后面有没有a
					for n := k + 1; n <= len(sl); n++ {
						if sl[n].Enum == 0 {

						}
					}
					//当s匹配完了，如果p还有剩
					if i == len(s)-1 && len(sl) > 1 {
						//检查是否存在 aa 或者 .
						for _, v1 := range sl {
							if v1.Enum == 0 || v1.Enum == 3 {
								return false
							}
						}
					}
					//万能匹配，跳出循环,
					break
				}
				if v.Enum == 3 {
					fmt.Println("匹配到了", string(s[i]))
					//防止k等于len时，切割panic
					if k >= len(sl) {
						break
					}
					//排除掉已经匹配过的正则，跳出循环
					sl = sl[k+1:]
					//k值归0
					k = 0
					break
				}
				// a* 表达式，为了防止二次匹配，将其剔除，并开始下一次循环
				if v.Enum == 1 {
					//防止k等于len时，切割panic
					if k >= len(sl) {
						break
					}
					fmt.Println(k)
					sl = sl[k+1:]
					//k值归0
					k = 0
					continue
				}
			}
			// 如果全部循环下来，都没有匹配到当前字母，则返回false
			if k == len(sl)-1 {
				return false
			}
			fmt.Println("没有匹配到", string(s[i]))
			k++
		}
	}

	//如果全部匹配通过，则返回true
	return true
}
