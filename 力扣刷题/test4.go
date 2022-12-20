package main

import (
	"fmt"
	"time"
)

func main() {
	main1()
	// main2()
	// main3()
}
func sliceTest() {
	for i := 0; i < 10000; i++ {
	}
}
func main1() {
	// var sli = []int{1, 2, 3, 4, 5}
	// var arr = [...]int{1, 2, 3, 4, 5}
	// echo("slice:", sli[3:])
	// echo("array:", arr[3:])
	// return
	defer func(start time.Time) {
		terminal := time.Since(start)
		fmt.Println(terminal)
	}(time.Now())
	data1 := []int{0, 5, 9, 10, 15, 17, 19, 20, 26, 30, 34}
	data2 := []int{1, 6, 7, 9, 11, 24, 31, 33}
	var res float64
	for i := 0; i < 100000; i++ {
		res = findMedianSortedArrays5(data1, data2)
	}
	echo("main1:", res) // 2.5
}

// func findMedianSortedArrays7(nums1 []int, nums2 []int) float64 {
// 	var m, n = len(nums1), len(nums2)
// 	// if m == 0 {
// 	// 	return ArrayMiddleVal(nums1)
// 	// }
// 	// if n == 0 {
// 	// 	return ArrayMiddleVal(nums1)
// 	// }

// 	isEven := (m+n)%2 == 0

// 	mmid, mindex := ArrayMiddleVal(nums1)
// 	nmid, nindex := ArrayMiddleVal(nums2)

// 	if mmid > nmid {

// 	}

// 	var x = len(c)
// 	if x%2 != 0 {
// 		return float64(c[(x+1)/2-1])
// 	}
// 	x = x / 2
// 	return float64(c[x-1]+c[x]) / 2
// }

func ArrayMiddleVal(nums []int) (float64, int) {
	var x = len(nums)
	if x%2 != 0 {
		i := (x+1)/2 - 1
		return float64(nums[i]), i
	}
	x = x / 2
	return float64(nums[x-1]+nums[x]) / 2, x
}

func findMedianSortedArrays5(nums1 []int, nums2 []int) float64 {
	c := generateSortArray(nums1, nums2)
	var x = len(c)
	if x%2 != 0 {
		return float64(c[(x+1)/2-1])
	}
	x = x / 2
	return float64(c[x-1]+c[x]) / 2
}

func generateSortArray(nums1 []int, nums2 []int) []int {
	var m, n = len(nums1), len(nums2)
	if m == 0 {
		return nums2
	}
	if n == 0 {
		return nums1
	}
	var c []int
	var mi, ni, mv, nv int
	for true {
		if mi >= m {
			return append(c, nums2[ni:]...)
		}
		if ni >= n {
			return append(c, nums1[mi:]...)
		}
		mv = nums1[mi]
		nv = nums2[ni]
		if mv == nv {
			c = append(c, mv, nv)
			mi++
			ni++
		} else if mv < nv {
			c = append(c, mv)
			mi++
		} else {
			c = append(c, nv)
			ni++
		}
	}
	return c
}

func findMedianSortedArrays6(nums1 []int, nums2 []int) float64 {
	var c = append(append([]int{}, nums1...), nums2...)
	type sort func(data []int, left, right int)
	var quickSort sort
	quickSort = func(data []int, left, right int) {
		if left >= right {
			return
		}
		pivot := data[left]
		i, j := left, right
		for i != j {
			for data[j] >= pivot && j > i {
				j--
			}
			for data[i] <= pivot && i < j {
				i++
			}
			data[i], data[j] = data[j], data[i]
		}
		data[left], data[i] = data[i], data[left]
		quickSort(data, left, i-1)
		quickSort(data, i+1, right)
	}
	quickSort(c, 0, len(c)-1)
	var x int = len(c)
	if x%2 != 0 {
		return float64(c[(x+1)/2-1])
	}
	x = x / 2
	return float64(c[x-1]+c[x]) / 2
}

func findMedianSortedArrays4(nums1 []int, nums2 []int) float64 {
	var m, n = len(nums1), len(nums2)
	var c []int
	var mi, ni, mv, nv, m1, n1, last int
	if m != 0 {
		m1 = nums1[0]
		if n != 0 {
			n1 = nums2[0]
		} else {
			n1 = m1
		}
	} else {
		n1 = nums2[0]
		m1 = n1
	}
	last = m1 - 1
	if last >= n1 {
		last = n1 - 1
	}
	for true {
		if mi >= m && ni >= n {
			break
		}
		if mi < m {
			mv = nums1[mi]
		} else {
			if nv = nums2[ni]; nv != last {
				c = append(c, nv)
				last = nv
			}
			ni++
			continue
		}
		if ni < n {
			nv = nums2[ni]
		} else {
			if mv = nums1[mi]; mv != last {
				c = append(c, mv)
				last = mv
			}
			mi++
			continue
		}
		if mv <= nv {
			if mv != last {
				c = append(c, mv)
				last = mv
			}
			mi++
			if mv == nv {
				ni++
			}
			continue
		}

		if nv != last {
			c = append(c, nv)
			last = nv
		}
		ni++
	}
	echo(c)
	var x int = len(c)
	if x%2 != 0 {
		return float64(c[(x+1)/2-1])
	}
	x = x / 2
	return float64(c[x-1]+c[x]) / 2
}

func main2() {
	defer func(start time.Time) {
		terminal := time.Since(start)
		fmt.Println(terminal)
	}(time.Now())
	data1 := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 4, 4}
	data2 := []int{1, 3, 4, 4, 4, 4, 4, 4, 4, 4, 4}
	var res float64
	for i := 0; i < 100000; i++ {
		res = findMedianSortedArrays2(data1, data2)
	}
	echo("main2:", res) // 3.0
}

func main3() {
	defer func(start time.Time) {
		terminal := time.Since(start)
		fmt.Println(terminal)
	}(time.Now())
	data1 := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 4, 4}
	data2 := []int{1, 3, 4, 4, 4, 4, 4, 4, 4, 4, 4}
	var res float64
	for i := 0; i < 100000; i++ {
		res = findMedianSortedArrays3(data1, data2)
	}
	echo("main3:", res) // 3.0
}

func findMedianSortedArrays3(nums1 []int, nums2 []int) float64 {
	var m, n = len(nums1), len(nums2)
	var c []int
	var mi, ni, mv, nv, last int
	var first = true
	// if m != 0 {
	// 	m1 = nums1[0]
	// } else {
	// 	m1 = nums2[n-1] + 1
	// }
	// if n != 0 {
	// 	n1 = nums2[0]
	// } else {
	// 	n1 = nums1[m-1] + 1
	// }
	// last = m1 - 1
	// if last >= n1 {
	// 	last = nv - 1
	// }

	for true {
		if mi >= m && ni >= n {
			break
		}
		if mi < m {
			mv = nums1[mi]
		} else {
			if nv = nums2[ni]; nv != last || first {
				c = append(c, nv)
				last = nv
			}
			ni++
			continue
		}
		if ni < n {
			nv = nums2[ni]
		} else {
			if mv = nums1[ni]; mv != last || first {
				c = append(c, mv)
				last = mv
			}
			mi++
			continue
		}
		if mv <= nv {
			if mv != last || first {
				c = append(c, mv)
				last = mv
			}
			mi++
			if mv == nv {
				ni++
			}
			continue
		}

		if nv != last || first {
			c = append(c, nv)
			last = nv
		}
		if first {
			first = false
		}
		ni++
	}
	var cl int = len(c)
	if cl%2 != 0 {
		return float64(c[(cl+1)/2-1])
	}
	var x int = cl / 2
	return float64(c[x-1]+c[x]) / 2
}

func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	var m, n = len(nums1), len(nums2)
	var c []int
	var mi, ni, mv, nv int

	var mp = make(map[int]bool)

	for true {
		if mi >= m && ni >= n {
			break
		}
		if mi < m {
			mv = nums1[mi]
		} else {
			nv = nums2[ni]
			if _, ok := mp[nv]; !ok {
				c = append(c, nv)
				mp[nv] = false
			}
			ni++
			continue
		}

		if ni < n {
			nv = nums2[ni]
		} else {
			mv = nums1[mi]
			if _, ok := mp[mv]; !ok {
				c = append(c, mv)
				mp[mv] = false
			}
			mi++
			continue
		}
		if mv <= nv {
			if _, ok := mp[mv]; !ok {
				c = append(c, mv)
				mp[mv] = false
			}
			mi++
			if mv == nv {
				ni++
			}
			continue
		}

		if _, ok := mp[nv]; !ok {
			c = append(c, nv)
			mp[nv] = false
		}
		ni++
	}
	var cl int = len(c)
	if cl%2 != 0 {
		return float64(c[(cl+1)/2-1])
	}
	var x int = cl / 2
	return float64(c[x-1]+c[x]) / 2
}

var echo = func(in ...interface{}) {
	for _, v := range in {
		fmt.Print(v)
	}
	fmt.Println()
}
