package main

import (
	"fmt"
	"time"
)

func echo(a ...interface{}) (int, error) {
	return fmt.Println(a...)
}

func main() {
	defer func(start time.Time) {
		terminal := time.Since(start)
		fmt.Println(terminal)
	}(time.Now())

	// echo(isMatch("abcdefg", ".*") == true)                         //true
	echo(isMatch("baabbbaccbccacacc", "c*..b*a*a.*a..*c") == true) //true
	// echo(isMatch("abbbcd", "ab*bbbcd") == true)                    //true
	// echo(isMatch("a", ".*.") == true)                              //true
	// echo(isMatch("bbbba", ".*a*a") == true)                        //true
	// echo(isMatch("a", ".*..a*") == false)                          //false
	// echo(isMatch("a", "ab*") == true)                              //true
	// echo(isMatch("ab", ".*b") == true)                             //true
	// echo(isMatch("ab", ".*.*") == true)                            //true
	// echo(isMatch("aa", "a") == false)                              //false
	// echo(isMatch("aa", "a*a") == true)                             //true
	// echo(isMatch("ab", ".*..") == true)                            //true
	// echo(isMatch("ab", ".*.*..") == true)                          //true
	// echo(isMatch("aab", "c*a*b") == true)                          //true
	// echo(isMatch("mississippi", "mis*is*p*.") == false)            //false
}

func isMatch(s string, p string) bool {
	//万能匹配表达式
	if p == ".*" || p == ".*." || p == "..*" {
		return true
	}
	//将p打散，[a]、[.]、[a*]、[.*]为一个单位
	pc := toSlice(p)
	var sclen, pclen = len(s), len(pc)
	var dp func(i, j int) bool
	dp = func(i, j int) bool {
		if i >= sclen && j >= pclen {
			return true
		}
		if j >= pclen {
			return false
		}
		if i >= sclen {
			for ; j < len(pc); j++ {
				if len(pc[j]) == 1 {
					return false
				}
			}
			return true
		}

		if len(pc[j]) == 2 {
			if pc[j][0] == s[i] || pc[j][0] == '.' { //case1
				return dp(i+1, j) || dp(i, j+1) //s移动，或者p移动
			} else { //case2
				return dp(i, j+1) //只移动p
			}
		}

		if pc[j][0] == s[i] || pc[j][0] == '.' { //case3
			return dp(i+1, j+1) //s和p同时移动
		}
		//case4
		return false
	}
	return dp(0, 0)
}

func toSlice(p string) (pc []string) {
	for i := 0; i < len(p); i++ {
		if p[i] != '*' {
			if i+1 < len(p) {
				if p[i+1] != '*' {
					pc = append(pc, string(p[i]))
				} else {
					pc = append(pc, p[i:i+2])
				}
			} else {
				pc = append(pc, string(p[i]))
			}
		}
	}
	return
}
