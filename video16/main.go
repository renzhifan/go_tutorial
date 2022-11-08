package main

import (
	"encoding/json"
	"fmt"
)

// 结构体字段的可见性与JSON序列化

// 如果一个Go语言包中定义的标识符是首字母大写的，那么就是对外可见的。
// 如果一个结构体中的字段名首字母是大写的，那么该字符就是对外可见的。

type student struct {
	ID   int
	Name string
}

// student的构造函数
func newStudent(id int, name string) student {
	return student{
		ID:   id,
		Name: name,
	}
}

type class struct {
	Title    string    `json:"title"`
	Students []student `json:"student_list" db:"student" xml:"ss"`
}

func main() {
	// 创建一个班级变量c1
	c1 := class{
		Title:    "火箭101",
		Students: []student{},
	}
	fmt.Printf("Students.point=%p\n", c1.Students)
	// 往班级c1中添加学生
	for i := 0; i < 10; i++ {
		// 造十个学生
		tmpStu := newStudent(i, fmt.Sprintf("stu%02d", i))
		c1.Students = append(c1.Students, tmpStu)
	}
	fmt.Printf("Students.point=%p\n", c1.Students)
	//fmt.Printf("%#v\n", c1)
	// JSON序列化：Go语言中的数据 -> JSON格式的字符串
	data, err := json.Marshal(c1)
	if err != nil {
		fmt.Println("json marshal failed, err:", err)
		return
	}
	//fmt.Printf("%T\n", data)
	fmt.Printf("data=%s\n", data)
	// JSON反序列化：JSON格式的字符串 -> Go语言中的数据
	//jsonStr := `{"title":"火箭101","Students":[{"ID":0,"Name":"小王子"},{"ID":1,"Name":"娜扎"},{"ID":2,"Name":"stu02"}]}`
	var c2 class
	err = json.Unmarshal(data, &c2)
	if err != nil {
		fmt.Println("json unmarshal failed, err:", err)
		return
	}
	fmt.Printf("%#v\n", c2)

	//a := [8]int{0, 1, 2, 3, 4, 5, 6, 7}
	//fmt.Printf("a=%p\n", &a)
	//s1 := a[:5]
	//fmt.Printf("s1=%p\n", &s1)
	//s2 := a[3:6]
	//fmt.Printf("s2.point=%p s2.point=%v\n", &s2)

	//s1 := make([]int, 3) //[0 0 0]
	//s2 := s1             //将s1直接赋值给s2，s1和s2共用一个底层数组
	//fmt.Printf("s1=%p\n", s1)
	//fmt.Printf("s2=%p\n", s2)
	//s2[0] = 100
	//
	//fmt.Println(s1) //[100 0 0]
	//fmt.Println(s2) //[100 0 0]

}
