package main

import (
	"fmt"
	"unsafe"
)

// 结构体的初始化

type person struct {
	name, city string
	age        int8
}

type test struct {
	a int8
	b int8
	c int8
	d int8
}

func main() {
	var p1 person
	fmt.Printf("p1=%#v p1.point=%p\n", p1, &p1) //p1=main.person{name:"", city:"", age:0} p1.point=0xc00009a180


	// 1. 键值对初始化
	p4 := person{
		name: "小王子",
		age:  18,
	}
	fmt.Printf("p4.value=%#v p4.point=%p\n", p4, &p4) //p4.value=main.person{name:"小王子", city:"", age:18} p4.point=0xc00009a1e0


	p5 := &person{
		name: "小王子",
		city: "北京",
		age:  18,
	}
	fmt.Printf("p5.value=%#v\n", p5) //p5.value=&main.person{name:"小王子", city:"北京", age:18}

	// 2. 值的列表进行初始化
	p6 := person{
		"小王子",
		"北京",
		18,
	}
	fmt.Printf("p6.value=%#v\n", p6) //p6.value=main.person{name:"小王子", city:"北京", age:18}

	n := test{
		1, 2, 3, 4,
	}
	fmt.Printf("n.a %p\n", &n.a) //n.a 0xc00010c080
	fmt.Printf("n.b %p\n", &n.b) //n.b 0xc00010c081
	fmt.Printf("n.c %p\n", &n.c) //n.c 0xc00010c082
	fmt.Printf("n.d %p\n", &n.d) //n.d 0xc00010c083

	var v struct{}                //空结构体是不占用空间的。
	fmt.Println(unsafe.Sizeof(v)) // 0

	//面试题
	//stus在range遍历过程中 stu会对stus对应索引进行值拷贝
	//而在这里是将stu的地址传递给 m ， 因此 m 的值 指向同一片地址空间
	m := make(map[string]*student)
	stus := []student{
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
		{name: "大王八", age: 9000},
	}
	fmt.Printf("stus=%p\n", &stus)
	for _, stu := range stus {
		fmt.Printf("stu.point=%p stu.value=%v stu.name=%s\n", &stu, stu, stu.name)
		m[stu.name] = &stu
		fmt.Printf("&stu=%v\n", &stu)
		//stu.point=0xc0000a4018 stu.value={小王子 18} stu.name=小王子
		//&stu=&{小王子 18}
		//stu.point=0xc0000a4018 stu.value={娜扎 23} stu.name=娜扎
		//&stu=&{娜扎 23}
		//stu.point=0xc0000a4018 stu.value={大王八 9000} stu.name=大王八
		//&stu=&{大王八 9000}

	}
	fmt.Println(m) //map[大王八:0xc00000c030 娜扎:0xc00000c030 小王子:0xc00000c030]

	for k, v := range m {
		fmt.Println(k, "=>", v.name)
		//娜扎 => 大王八
		//大王八 => 大王八
		//小王子 => 大王八

	}
}

type student struct {
	name string
	age  int
}
