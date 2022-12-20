package main

import "fmt"

func main() {
	ori := []int{5, 9, 4, 7, 3, 6, 2, 8, 1, 0}
	QuickSort(ori)
	fmt.Println(ori)
}

func QuickSort(data []int) {
	quickSort(data, 0, len(data)-1)
}

/** quickSort 快速排序算法
 * description: 采用左右起止指针，递归方式
 * data type []int 排序的数据
 * startIndex type int 起始位置
 * endIndex type int 结束位置
 */
func quickSort(data []int, startIndex, endIndex int) {
	//left，right是该切片段的起止位置，如果相同，则表示将要操作的数据长度为 1 ，无需排序，直接结束
	if startIndex >= endIndex {
		return
	}
	//pivot取左策略，设为[startIndex,endIndex]数列的左边第一个元素，优点是无需计算
	pivot := data[startIndex]

	// 设置左右指针，左指针快，右指针慢
	left, right := startIndex, endIndex
	//开始遍历，左指针始终小于右指针
	for left < right {
		//右指针开始向左移动，如果指针元素大于等于 pivot，则继续左移动，但不能跑到left指针的左边去，直到遇到小于pivot的元素为止
		for data[right] >= pivot && right > left {
			right--
		}

		//左指针再开始向右移动，如果指针元素小于等于 pivot，则继续右移动，但不能跑到right指针的右边去，直到遇到大于pivot的元素为止，
		for data[left] <= pivot && left < right {
			left++
		}

		//当 left != right ：left所处位置的值大于 pivot，right所处位置的值小于 pivot
		//需要交换left与right指针的值
		data[left], data[right] = data[right], data[left]
	}
	//关键一步，将pivot位置data[startIndex]的值与指针相遇的位置data[left]值调换，等于根据pivot将数据一分为二，左右两部分
	data[startIndex], data[left] = data[left], data[startIndex]
	//将分好的两组数据继续排序，递归方式，注意，此时pivot值得位置已经排序好了，不会再变，所以不需要再继续操作，所以有 i-1 和 i+1
	quickSort(data, startIndex, left-1)
	quickSort(data, left+1, endIndex)
}
