package main

import (
	"fmt"
	"strings"
)

//匿名函数和闭包
//定义一个函数它的返回值是一个函数
//把函数作为返回值
func a(name string) func() {
	return func() {
		fmt.Println("hello", name)
	}
}

func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

func adder() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}
func deferAdd(a, b int) func() int {
	return func() int {
		return a + b
	}
}

func main() {
	// 此时返回的是匿名函数
	//addFunc := deferAdd(1, 2)
	// 这里才会真正执行加法操作
	//fmt.Println(addFunc())

	var f = adder()
	fmt.Println(f(10)) //10
	fmt.Println(f(10)) //20
	fmt.Println(f(20)) //40
	fmt.Println(f(30)) //70

	//f1 := adder()
	//fmt.Println(f1(40)) //40
	//fmt.Println(f1(50)) //90
	// 将匿名函数保存到变量
	//add := func(x, y int) {
	//	fmt.Println(x + y)
	//}
	//add(10, 20) // 通过变量调用匿名函数

	//自执行函数：匿名函数定义完加()直接执行
	//func(x, y int) {
	//	fmt.Println(x + y)
	//}(10, 20)

	//闭包 = 函数 + 外层变量的引用
	//x, y := calc(100)
	//ret1 := x(200) //base = 100 + 200
	//fmt.Println(ret1)
	//ret2 := y(200) //base = 300 - 200
	//fmt.Println(ret2)
}
