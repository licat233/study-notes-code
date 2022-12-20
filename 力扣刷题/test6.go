package main

import (
	"bytes"
	"fmt"
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
	data := "PAYPALISHIRING"
	row := 4
	var res string
	res = convert(data, row)
	fmt.Println(res) //PINALSIGYAHRPI
}

func convert(s string, numRows int) string {
	if numRows < 2 {
		return s
	}
	var n = len(s)
	if n <= 2 {
		return s
	}
	var row, col, i, x int
	var num = 2*numRows - 2
	var result bytes.Buffer
	//第一行
	for i = 0; i < n; i = i + num {
		result.WriteByte(s[i])
	}
	//中间部分，从上到下，从左到右，所以先遍历行，
	//中间部分行数
	var midRows = numRows - 2
	//中间部分列数
	var midCols = n/num + 1
	//遍历行数
	for row = 0; row < midRows; row++ {
		//遍历列数，这里是大列，表示一个V
		for col = 0; col < midCols; col++ {
			//先插入第一小列
			x = col*num + row + 1
			if x >= n {
				break
			}
			result.WriteByte(s[x])
			//再插入第二小列
			x = col*num + num - row - 1
			if x >= n {
				break
			}
			result.WriteByte(s[x])
		}
	}
	//最后一行
	for i = numRows - 1; i < n; i = i + num {
		result.WriteByte(s[i])
	}
	return result.String()
}

func convert1(s string, numRows int) string {
	if numRows < 2 {
		return s
	}
	var n = len(s)
	var mids = numRows - 2
	var result bytes.Buffer
	var num = (numRows + mids)
	for r := 0; r < numRows; r++ {
		for i := 0; i < n; i++ {
			if (i-r)%num == 0 {
				result.WriteByte(s[i])
				if r != 0 && r != numRows-1 {
					x := i + (numRows-r-1)*2
					if x < n {
						result.WriteByte(s[x])
					}
				}
			}
		}
	}
	return result.String()
}
