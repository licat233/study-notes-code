package main

import (
	"fmt"
	"sync"
	"time"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func generateListNode(data []int) *ListNode {
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

func echoNode(l *ListNode) {
	res := []int{}
	for l != nil {
		res = append(res, l.Val)
		l = l.Next
	}
	fmt.Println(res)
}

func main2() {
	defer func(start time.Time) {
		terminal := time.Since(start)
		fmt.Println(terminal)
	}(time.Now())
	data1 := []int{2, 4, 9}
	data2 := []int{5, 6, 4, 9}
	ld1 := generateListNode(data1)
	// echoNode(ld1)
	ld2 := generateListNode(data2)
	// echoNode(ld2)
	res := addTwoNumbers(ld1, ld2)
	echoNode(res) //[7,0,4,0,1]
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var res = &ListNode{}
	var lastNode = res
	var sum int
	for true {
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		if sum >= 10 {
			lastNode.Val = sum - 10
			sum = 1
		} else {
			lastNode.Val = sum
			if l1 == nil && l2 == nil {
				return res
			}
			sum = 0
		}

		lastNode.Next = &ListNode{}
		lastNode = lastNode.Next
	}
	return res
}

func main2s() {
	p := &sync.Pool{
		New: func() interface{} {
			return 0
		},
	}
	p.Put(2)
	p.Put(1)
	a := p.Get().(int)

	b := p.Get().(int)

	c := p.Get().(int)
	fmt.Println(a, b, c)
}
