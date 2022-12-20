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
func performance(list1 *ListNode, list2 *ListNode) {
	defer func(start time.Time) {
		terminal := time.Since(start)
		fmt.Println(terminal)
	}(time.Now())
	echoList(mergeTwoLists(list1, list2))
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

func main() {
	performance(genNode([]int{1, 2, 4}), genNode([]int{1, 3, 4}))
	performance(genNode([]int{}), genNode([]int{}))
	performance(genNode([]int{}), genNode([]int{0}))
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}
	var res *ListNode
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

func mergeTwoLists2(list1 *ListNode, list2 *ListNode) *ListNode {
	var res = &ListNode{}
	node := res
	for ; list1 != nil && list2 != nil; node = node.Next {
		if list1.Val < list2.Val {
			node.Next = &ListNode{
				Val: list1.Val,
			}
			list1 = list1.Next
		} else {
			node.Next = &ListNode{
				Val: list2.Val,
			}
			list2 = list2.Next
		}
	}
	if list1 != nil {
		node.Next = list1
	} else if list2 != nil {
		node.Next = list2
	}
	return res.Next
}
