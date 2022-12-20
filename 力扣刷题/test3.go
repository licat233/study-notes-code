package main

import (
	"fmt"
	"time"
)

func main3() {
	defer func(start time.Time) {
		terminal := time.Since(start)
		fmt.Println(terminal)
	}(time.Now())
	data := "pazxpgcaxwazg"
	res := 0
	for i := 0; i < 100000; i++ {
		res = lengthOfLongestSubstring(data)
	}
	fmt.Println(res) // 6
}

//方案二：①一个map包含最长字符串；②一次遍历
func lengthOfLongestSubstring(s string) int {
	var sl = len(s)
	if sl == 0 {
		return 0
	}
	var ok bool
	var max, dl, z int
	var dm = make(map[rune]int)
	for k, v := range s {
		if z, ok = dm[v]; ok && k-z <= dl {
			dl = k - z
		} else {
			dl++
		}
		dm[v] = k
		if dl > max {
			max = dl
			continue
		}
		if sl-k-1+dl <= max {
			return max
		}
	}
	return max
}

//方案一：①包含哪些字母；②有多长；
// pazx
// azx
// zx
// xzc
// zcxwa
// cxwaz
// xwaz
// return
func lengthOfLongestSubstring1(s string) int {
	if len(s) == 0 {
		return 0
	}
	var max int
	var sl = len(s)
	var dl int
	var d map[byte]bool
	for k := 0; k < sl && max <= sl-k; k++ {
		d = make(map[byte]bool)
		d[s[k]] = false
		for i := k + 1; i < sl; i++ {
			if _, ok := d[s[i]]; ok {
				break
			}
			d[s[i]] = false
		}
		dl = len(d)
		if dl > max {
			max = dl
		}
	}
	return max
}

func echo(in ...interface{}) {
	for _, v := range in {
		fmt.Print(v)
	}
	fmt.Println()
}
