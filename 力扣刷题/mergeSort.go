package main

import "fmt"

func main() {
	var arr = []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	res := mergeSort(arr)
	fmt.Println(res)
}

//并归排序算法
func mergeSort(arr []int) []int {
	n := len(arr)
	if n < 2 {
		return arr
	}
	mid := n / 2
	left := arr[0:mid]
	right := arr[mid:]
	return mergeTwoInts(mergeSort(left), mergeSort(right))
}

//合并排序两个升序切片
func mergeTwoInts(before, after []int) []int {
	var res []int
	i, in, j, jn := 0, len(before), 0, len(after)
	for i < in && j < jn {
		if before[i] < after[j] {
			res = append(res, before[i])
			i++
		} else {
			res = append(res, after[j])
			j++
		}
	}
	if i != in {
		res = append(res, before[i:]...)
	}
	if j != jn {
		res = append(res, after[j:]...)
	}
	return res
}
