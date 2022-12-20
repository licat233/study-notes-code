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
func performance(data *ListNode) {
	defer func(start time.Time) {
		terminal := time.Since(start)
		fmt.Println("me:", terminal)
	}(time.Now())
	// for i := 0; i < 100000; i++ {
	swapPairs(data)
	// }
}

func main() {
	data := &ListNode{}
	performance(data)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var before, after = head, head.Next
	if after == nil {
		return before
	}

	var temp = after.Next
	after.Next = before
	before.Next = temp
	swapTwo(before, temp)
	return after
}

func swapTwo(head, list *ListNode) {
	if list == nil {
		return
	}
	var before, after = list, list.Next
	if after == nil {
		return
	}
	var temp = after.Next
	if temp == nil {
		after.Next = before
		before.Next = nil
		head.Next = after
		return
	}
	after.Next = before
	before.Next = temp
	head.Next = after
	swapTwo(before, temp)
	return
}
