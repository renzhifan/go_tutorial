package main

import "fmt"

// 结构体指针

type person struct {
	name, city string
	age        int8
}

func main() {
	var p2 = new(person)
	fmt.Printf("p2.type=%T\n", p2) //p2.type=*main.person

	// (*p2).name = "小王子"
	// (*p2).city = "北京"
	// (*p2).age = 18
	p2.name = "小王子"
	p2.city = "北京"
	p2.age = 18
	fmt.Printf("p2.value=%#v\n", p2) //p2.value=&main.person{name:"小王子", city:"北京", age:18}

	// 取结构体的地址进行实例化
	p3 := &person{}
	fmt.Printf("p3.type=%T\n", p3) //p3.type=*main.person

	fmt.Printf("p3.value=%#v\n", p3) //p3.value=&main.person{name:"", city:"", age:0}

	p3.name = "娜扎"
	p3.city = "北京"
	p3.age = 18
	fmt.Printf("p3.value=%#v\n", p3)//p3.value=&main.person{name:"娜扎", city:"北京", age:18}

}
