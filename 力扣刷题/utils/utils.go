package utils

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

func performance() {
	defer func(start time.Time) {
		terminal := time.Since(start)
		fmt.Println("me:", terminal)
	}(time.Now())
}
