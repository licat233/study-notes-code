package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func echo(a ...interface{}) (int, error) {
	return fmt.Println(a...)
}
func performance(data string) {
	defer func(start time.Time) {
		terminal := time.Since(start)
		fmt.Println("me:", terminal)
		wg.Done()
	}(time.Now())
	for i := 0; i < 100000; i++ {
		letterCombinations(data)
	}
}

func performance2(data string) {
	defer func(start time.Time) {
		terminal := time.Since(start)
		fmt.Println("other:", terminal)
		wg.Done()
	}(time.Now())
	for i := 0; i < 100000; i++ {
		letterCombinations(data)
	}
}

func performance3(data string) {
	defer func(start time.Time) {
		terminal := time.Since(start)
		fmt.Println("me2:", terminal)
		wg.Done()
	}(time.Now())
	for i := 0; i < 100000; i++ {
		letterCombinations3(data)
	}
}

func main() {
	echo(letterCombinations("23"))
	return
	wg.Add(1)
	go performance("23")
	wg.Add(1)
	go performance2("23")
	wg.Add(1)
	go performance3("23")
	wg.Wait()
}

func letterCombinations(digits string) (res []string) {
	n := len(digits)
	//因为digits长度最大为 4，所以最佳方案是用枚举法来解：
	var letter1, letter2, letter3, letter4 rune
	var group1, group2, group3 string
	//如果长度有 1 ，则遍历第一个
	if n > 0 {
		//建立映射数据库，相比map[byte][]string，map空间开销会更小一些（map在这里创建会有一定的优化效果）
		m := map[byte]string{
			'2': "abc",
			'3': "def",
			'4': "ghi",
			'5': "jkl",
			'6': "mno",
			'7': "pqrs",
			'8': "tuv",
			'9': "wxyz",
		}
		for _, letter1 = range m[digits[0]] {
			group1 = string(letter1)
			//如果长度有2，则遍历第二个
			if n > 1 {
				for _, letter2 = range m[digits[1]] {
					group2 = group1 + string(letter2)
					//如果长度有3，则遍历第三个
					if n > 2 {
						for _, letter3 = range m[digits[2]] {
							group3 = group2 + string(letter3)
							//如果长度有4，则遍历第四个
							if n > 3 {
								for _, letter4 = range m[digits[3]] {
									//长度最多有4个，我们就不往下枚举了
									//此时就得到了一个字母组合
									res = append(res, group3+string(letter4))
								}
							} else { // 如果长度只有3个
								res = append(res, group3)
							}
						}
					} else { // 如果长度只有2个
						res = append(res, group2)
					}
				}
			} else { // 如果长度只有1个
				res = append(res, group1)
			}
		}
	} //如果长度为0个
	return
}

func letterCombinations2(digits string) []string {
	c := map[byte][]byte{
		'2': []byte("abc"),
		'3': []byte("def"),
		'4': []byte("ghi"),
		'5': []byte("jkl"),
		'6': []byte("mno"),
		'7': []byte("pqrs"),
		'8': []byte("tuv"),
		'9': []byte("wxyz"),
	}
	a1 := make([]string, 0)
	for _, v := range []byte(digits) {
		if d, ok := c[v]; ok {
			if len(a1) == 0 {
				for _, i := range d {
					a1 = append(a1, string(i))
				}
			} else {
				a2 := make([]string, 0, len(a1)*len(d))
				for _, i := range d {
					for _, a := range a1 {
						g := append([]byte(a), i)
						a2 = append(a2, string(g))
					}
				}
				a1 = a2
			}
		}
	}
	return a1
}

func letterCombinations3(digits string) (res []string) {
	n := len(digits)
	//特殊情况
	if n == 0 {
		return
	}
	//初始化
	//建立映射数据库
	m := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	for _, v := range m[digits[0]] {
		res = append(res, string(v))
	}
	for i := 1; i < n; i++ {
		var newRes []string
		//取出res中的每个元素出来进行组合
		for _, group := range res {
			//遍历数字映射的字母，每一个数字取出一个字母来进行组合
			for _, letter := range m[digits[i]] {
				newRes = append(newRes, group+string(letter)) //添加新的组合
			}
		}
		res = newRes
	}
	return
}
