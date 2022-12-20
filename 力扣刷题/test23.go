package main

import (
	"fmt"
	"time"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func echo(a ...interface{}) (int, error) {
	return fmt.Println(a...)
}

func echoList(list *ListNode) {
	var res []int
	for node := list; node != nil; node = node.Next {
		res = append(res, node.Val)
	}
	echo(res)
}

func genNode(data []int) *ListNode {
	if len(data) == 0 {
		return nil
	}
	var node = &ListNode{}
	var count int
	for _, v := range data {
		count++

		newNode := ListNode{
			Val:  v,
			Next: nil,
		}
		var nd = node
		for nd.Next != nil {
			nd = nd.Next
		}
		if count == 1 {
			*nd = newNode
		} else {
			nd.Next = &newNode
		}
	}
	return node
}
func performance(lists []*ListNode) {
	defer func(start time.Time) {
		terminal := time.Since(start)
		fmt.Println(terminal)
	}(time.Now())
	echoList(mergeKLists(lists))
}
func main() {
	// performance([]*ListNode{genNode([]int{2}), genNode([]int{}), genNode([]int{-1})})
	// performance([]*ListNode{genNode([]int{}), genNode([]int{1})})
	// performance([]*ListNode{genNode([]int{})})
	performance([]*ListNode{genNode([]int{}), genNode([]int{})})
	performance([]*ListNode{genNode([]int{1, 4, 5}), genNode([]int{1, 3, 4}), genNode([]int{2, 6})})
}

func mergeKLists(lists []*ListNode) *ListNode {
	n := len(lists)
	if n == 0 {
		return nil
	} else if n == 1 {
		return lists[0]
	}
	mid := n / 2
	return mergeTwoLists(mergeKLists(lists[0:mid]), mergeKLists(lists[mid:]))
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	res := list1
	if list1.Val <= list2.Val {
		list1 = list1.Next
	} else {
		res = list2
		list2 = list2.Next
	}
	node := res
	for ; list1 != nil && list2 != nil; node = node.Next {
		if list1.Val <= list2.Val {
			node.Next = list1
			list1 = list1.Next
		} else {
			node.Next = list2
			list2 = list2.Next
		}
	}
	if list1 != nil {
		node.Next = list1
	} else if list2 != nil {
		node.Next = list2
	}
	return res
}
