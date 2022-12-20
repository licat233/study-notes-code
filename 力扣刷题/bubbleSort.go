package main

import "fmt"

func main() {
	ori := []int{5, 9, 4, 7, 3, 6, 2, 8, 1, 0}
	bubbleSort(ori)
	fmt.Println(ori)
}

//冒泡排序
func bubbleSort(arr []int) {
	n := len(arr)
	//需要n次冒泡
	for i := n; i > 1; i-- {
		//开始冒泡,j从尾部开始
		for j := 1; j < i; j++ {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
}
