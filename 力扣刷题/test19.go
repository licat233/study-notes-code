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
func performance(head *ListNode, n int) {
	defer func(start time.Time) {
		terminal := time.Since(start)
		fmt.Println(terminal)
	}(time.Now())
	echoList(removeNthFromEnd(head, n))
}
func main() {
	// data1 := &ListNode{
	// 	Val: 1,
	// 	Next: &ListNode{
	// 		Val: 2,
	// 		Next: &ListNode{
	// 			Val: 3,
	// 			Next: &ListNode{
	// 				Val: 4,
	// 				Next: &ListNode{
	// 					Val:  5,
	// 					Next: nil,
	// 				},
	// 			},
	// 		},
	// 	},
	// }
	// data2 := &ListNode{
	// 	Val:  1,
	// 	Next: nil,
	// }
	data3 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: nil,
		},
	}

	// performance(data1, 2) //[1,2,3,5]
	// echo()
	// performance(data2, 1) //[]
	// echo()
	// performance(data3, 1) //[1]
	// echo()
	performance(data3, 2) //[2]
	echo()
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var res, delMark = head, head
	i := -1
	for res = head; head != nil; head = head.Next {
		if i < n {
			i++
			continue
		}
		delMark = delMark.Next
	}
	if i < n {
		return res.Next
	}
	delMark.Next = delMark.Next.Next
	return res
}

func removeNthFromEnd3(head *ListNode, n int) (res *ListNode) {
	var nodes []*ListNode
	res = head
	for head != nil {
		nodes = append(nodes, head)
		head = head.Next
	}
	var sz = len(nodes)
	delIndex := sz - n
	if delIndex == 0 {
		res = res.Next
		return
	}
	delBefore := delIndex - 1
	nodes[delBefore].Next = nodes[delIndex].Next
	return
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd2(head *ListNode, n int) (res *ListNode) {
	var nodes []*ListNode
	res = head
	for head != nil {
		nodes = append(nodes, head)
		head = head.Next
	}
	var sz = len(nodes)
	i := sz - n - 1
	j := sz - n + 1
	if i >= 0 && j >= sz {
		nodes[i].Next = nil
		return
	} else if i >= 0 && j < sz {
		nodes[i].Next = nodes[j]
		return
	} else if i < 0 && j >= sz {
		return nil
	} else if i < 0 && j < sz {
		res = nodes[j]
		return
	}
	return
}

func removeNthFromEnd1(head *ListNode, n int) (res *ListNode) {
	var sz int
	res = head
	for sz = 0; head != nil; sz++ {
		head = head.Next
	}
	if sz == 1 {
		return nil
	}
	if n = sz - n - 1; n < 0 {
		res = res.Next
		return
	}
	for node, i := res, 0; node != nil; i++ {
		if i == n {
			node.Next = node.Next.Next
			return
		}
		node = node.Next
	}
	return
}

func echoList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val)
		head = head.Next
	}
	fmt.Println()
}

// func getNextListNode(head *ListNode, i, n int) *ListNode {
// 	if head == nil {
// 		return nil
// 	}
// 	i++
// 	//找到第n个节点了
// 	if i == n {
// 		//复制该节点之后的列表
// 		restListNode := head.Next
// 	}
// 	return getNextListNode(head, i)
// }
