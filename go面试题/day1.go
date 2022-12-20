package main

import (
	"fmt"
)

func main() {
}

func test5() {
	sn1 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}

	sn2 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}

	sn3 := struct {
		name string
		age  int
	}{age: 11, name: "qq"}
	fmt.Println(sn3)

	//结构体只能比较是否相等，但是不能比较大小
	if sn1 == sn2 {
		fmt.Println("sn1 == sn2")
	}

	//相同类型的结构体才能够进行比较，结构体是否相同不但与属性类型有关，还与属性顺序相关
	// if sn1 == sn3 {
	// 	fmt.Println("sn1 == sn3")
	// }

	sm1 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}
	fmt.Println(sm1)

	sm2 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}
	fmt.Println(sm2)

	//常见的有 bool、数值型、字符、指针、数组等可以比较
	//切片、map、函数等不能比较
	// if sm1 == sm2 {
	// 	fmt.Println("sm1 == sm2")
	// }
}

func test4b() {
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5}
	// s1 = append(s1, s2) //错误，不能使用slice作为append的第二个argument
	s1 = append(s1, s2...) //正确使用方式
	fmt.Println(s1)
}

func test4a() {
	list := new([]int) //new返回指针类型
	// list = append(list, 1) //会报错
	fmt.Println(list)
}

func test3c() {
	/**
	 *new() 与 make() 的区别:
	 *new(T) 会为 T 类型的新值分配已置零的内存空间，并返回地址（指针），即类型为 *T的值。换句话说就是，返回一个指针，该指针指向新分配的、类型为 T 的零值。适用于值类型，如数组、结构体等。
	 *make(T,args) 返回初始化之后的 T 类型的值，这个值并不是 T 类型的零值，也不是指针 *T，是经过初始化之后的 T 的引用。make() 只适用于 slice、map 和 channel.
	 *简单一句话：否是分配内存，但new返回指针类型，make返回类型本身
	 */
}
func test3b() {
	s := make([]int, 0)
	s = append(s, 1, 2, 3, 4)
	fmt.Println(s) //[1 2 3 4]
}
func test3a() {
	s := make([]int, 5) //int的默认值为 0，[0,0,0,0,0]
	s = append(s, 1, 2, 3)
	fmt.Println(s) //[0 0 0 0 0 1 2 3]
}

func test2() {
	slice := []int{0, 1, 2, 3}
	m := make(map[int]*int)

	for key, val := range slice {
		fmt.Println(&val) //每次循环的val地址都是一样的
		m[key] = &val     //所以赋值给m[key]的地址值都是一样的
	}

	for k, v := range m {
		fmt.Println(k, "->", *v)
		/**
		0 -> 3
		1 -> 3
		2 -> 3
		3 -> 3
		*/
	}
}

func test1() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()
	panic("触发异常")
	//defer 的执行顺序是后进先出。当出现 panic 语句的时候，会先按照 defer 的后进先出的顺序执行，最后才会执行panic
	/**
	  打印后
	  打印中
	  打印前
	  panic: 触发异常
	*/
}
