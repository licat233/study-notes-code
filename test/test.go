package main

import (
	"fmt"
	"regexp"
)

func main2() {
	var s = "IVXLCDM"
	var res string = "字符"
	res = string([]byte{s[0], s[2]}) + res
	fmt.Println(res)
}

func main3() {
	PhoneNumber := "123a456789"
	if ok, _ := regexp.MatchString("^\\d{6,15}$", PhoneNumber); !ok {
		fmt.Println("匹配失败")
		return
	} else {
		fmt.Println("匹配成功")
	}
}

func main4() {
	a := 10
	b := 3
	var c int = a - a%b
	fmt.Println(c)
}

func main() {
	var s []int = nil
	s = append(s, 1, 2, 3)
	fmt.Println(s)
}
