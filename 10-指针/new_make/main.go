package main

import "fmt"

func main() {
	//var a *int //`var a *int`只是声明了一个指针变量a但是没有初始化，`指针作为引用类型需要初始化后才会拥有内存空间`，才可以给它赋值。
	//*a = 100
	//fmt.Println(a)
	//
	//var b map[string]int
	//b["沙河娜扎"] = 100
	//fmt.Println(b)

	//var a *int
	//a = new(int) //使用内置的new函数对a进行初始化之后就可以正常对其赋值了：
	//fmt.Println(*a) // 0
	//*a = 100
	//fmt.Println(*a) // 100

	var b map[string]int //var b map[string]int只是声明变量b是一个map类型的变量
	b = make(map[string]int, 10) //使用make函数进行初始化操作之后，才能对其进行键值对赋值
	b["沙河娜扎"] = 100
	fmt.Println(b)
}
