package main

import "fmt"

// 指针

/*
func main() {
	a := 10 // int类型

	b := &a                      // *int类型（int指针）
	fmt.Printf("%v %p\n", a, &a) // 10 0xc00001a098
	fmt.Printf("%v %p\n", b, b)  // 0xc00001a098 0xc00001a098
	// 变量b是一个int类型的指针（*int），它存储的是变量a的内存地址
	c := *b        // 根据内存地址去取值
	fmt.Println(c) // 10
}
*/

//变量、指针地址、指针变量、取地址、取值的相互关系和特性如下：
//- 对变量进行取地址（&）操作，可以获得这个变量的指针变量。
//- 指针变量的值是指针地址。
//- 对指针变量进行取值（*）操作，可以获得指针变量指向的原变量的值。

func modify1(x int) {
	x = 100
}

func modify2(y *int) {
	fmt.Printf("y.type=%T y=%p y=%d\n", y, y, *y)
	*y = 100
}

func main() {
	a := 1
	modify1(a)
	fmt.Println(a) // 1
	modify2(&a)
	fmt.Println(a)
}
