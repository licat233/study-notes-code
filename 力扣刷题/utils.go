package main

import "fmt"

var echo = func(in ...interface{}) {
	for _, v := range in {
		fmt.Print(v)
	}
	fmt.Println()
}
