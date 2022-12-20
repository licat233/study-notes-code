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
func echoList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val)
		head = head.Next
	}
	fmt.Println()
}

func performance(data *ListNode, k int) {
	defer func(start time.Time) {
		terminal := time.Since(start)
		fmt.Println("me:", terminal)
	}(time.Now())
	echoList(reverseKGroup(data, k))
}

func genData() *ListNode {
	return &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
}
func genList(list []int) *ListNode {
	n := len(list)
	if n == 0 {
		return nil
	}
	var res = &ListNode{
		Val:  list[0],
		Next: nil,
	}
	node := res
	for i := 1; i < n; i++ {
		node.Next = &ListNode{
			Val:  list[i],
			Next: nil,
		}
		node = node.Next
	}
	return res
}

func main() {
	echoList(genData())
	performance(genData(), 2)            // [2,1,4,3,5]
	performance(genData(), 3)            // [3,2,1,4,5]
	performance(genList([]int{1}), 1)    // [1]
	performance(genList([]int{1, 2}), 2) // [1,2]
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}
	var allNode = []*ListNode{}
	for ; head != nil; head = head.Next {
		allNode = append(allNode, head)
	}
	var count = len(allNode)
	if count < k {
		return allNode[0]
	}
	endNode := allNode[0]
	swapNum := count - count%k
	for b, i := 0, 0; i < swapNum; {
		//截取位置:【i,i+k]
		endNode.Next = allNode[i+k-1]
		for b = i + k - 1; b > i; b-- {
			allNode[b].Next = allNode[b-1]
		}
		endNode, i = allNode[i], i+k
	}
	if swapNum == count {
		endNode.Next = nil
	} else {
		endNode.Next = allNode[swapNum]
	}
	return allNode[k-1]
}

func reverseKGroup3(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}
	var allNode = []*ListNode{}
	for ; head != nil; head = head.Next {
		allNode = append(allNode, head)
	}
	var count = len(allNode)
	if count < k {
		return allNode[0]
	}
	var endNode *ListNode
	for i := 0; i <= count; i++ {
		if i != 0 && i%k == 0 {
			//截取位置:【i-k,i]
			a, b := i-k, i-1
			if endNode != nil {
				endNode.Next = allNode[b]
			}
			for ; b > a; b-- {
				allNode[b].Next = allNode[b-1]
			}
			endNode = allNode[a]
		}
	}
	if y := count % k; y != 0 {
		y = count - y
		endNode.Next = allNode[y]
	} else {
		endNode.Next = nil
	}
	return allNode[k-1]
}

func reverseKGroup2(head *ListNode, k int) *ListNode {
	var getGroupList = func(list *ListNode) ([]*ListNode, bool) {
		var group = []*ListNode{}
		for n, node := 0, list; node != nil; node = node.Next {
			group = append(group, node)
			if n++; n >= k {
				return group, true
			}
		}
		return nil, false
	}

	group, ok := getGroupList(head)
	if !ok {
		return head
	}
	n := k - 1
	first, last, list := group[n], group[0], group[n].Next
	for i := n; i > 0; i-- {
		group[i].Next = group[i-1]
	}
	last.Next = list
	var reverseHandler func(*ListNode, *ListNode)
	reverseHandler = func(before, list *ListNode) {
		group, ok := getGroupList(list)
		if !ok {
			return
		}
		n := k - 1
		first, last, list := group[n], group[0], group[n].Next
		for i := n; i > 0; i-- {
			group[i].Next = group[i-1]
		}
		last.Next = list
		before.Next = first
		reverseHandler(last, list)
		return
	}
	reverseHandler(last, list)
	return first
}

func reverseKGroup1(head *ListNode, k int) *ListNode {
	var getGroupList = func(list *ListNode) ([]*ListNode, bool) {
		var group = []*ListNode{}
		for n, node := 0, list; node != nil; node = node.Next {
			group = append(group, node)
			if n++; n >= k {
				return group, true
			}
		}
		return nil, false
	}

	group, ok := getGroupList(head)
	if !ok {
		return head
	}
	n := k - 1
	first, last, list := group[n], group[0], group[n].Next
	for i := n; i > 0; i-- {
		group[i].Next = group[i-1]
	}
	last.Next = list
	var reverseHandler func(*ListNode, *ListNode)
	reverseHandler = func(before, list *ListNode) {
		group, ok := getGroupList(list)
		if !ok {
			return
		}
		n := k - 1
		first, last, list := group[n], group[0], group[n].Next
		for i := n; i > 0; i-- {
			group[i].Next = group[i-1]
		}
		last.Next = list
		before.Next = first
		reverseHandler(last, list)
		return
	}
	reverseHandler(last, list)
	return first
}
