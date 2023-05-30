package main

import "fmt"

// 结构体构造函数:构造一个结构体实例的函数
// 结构体是值类型
type person struct {
	name, city string
	age        int8
}

// 构造函数
func newPerson(name, city string, age int8) *person {
	return &person{
		name: name,
		city: city,
		age:  age,
	}
}
func (p *person) SetAge(newAge int8) {
	p.age = newAge
}
func (p person) SetAge2(newAge int8) {
	p.age = newAge
}

//MyInt 将int定义为自定义MyInt类型
type MyInt int

//SayHello 为MyInt添加一个SayHello的方法
func (m MyInt) SayHello() {
	fmt.Println("Hello, 我是一个int。")
}

//结构体的匿名字段
type person2 struct { //>注意：这里匿名字段的说法并不代表没有字段名，而是默认会采用类型名作为字段名，结构体要求字段名称必须唯一，因此一个结构体中同种类型的匿名字段只能有一个。
	string
	int
}

//Address 地址结构体
type Address struct {
	Province   string
	City       string
	CreateTime string
}

//User 用户结构体
type User struct {
	Name    string
	Gender  string
	Address Address
}

//Email 邮箱结构体
type Email struct {
	Account    string
	CreateTime string
}

//User2 用户结构体
type User2 struct {
	Name    string
	Gender  string
	Address //匿名字段
}

//User3 用户结构体
type User3 struct {
	Name    string
	Gender  string
	Address //匿名字段
	Email   //匿名字段
}

//结构体的“继承”

type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%s会动！\n", a.name)
}

//Dog 狗
type Dog struct {
	Feet    int8
	*Animal //通过嵌套匿名结构体实现继承
}


func (d *Dog) say() {
	fmt.Printf("%s会汪汪汪~\n", d.name)
}

type Person struct {
	name   string
	age    int8
	dreams []string
}

/*func (p *Person) SetDreams(dreams []string) {
	p.dreams = dreams
}*/
func (p *Person) SetDreams(dreams []string) {
	p.dreams = make([]string, len(dreams))
	copy(p.dreams, dreams)
}
func main() {
	p1 := newPerson("小王子", "北京", int8(18))
	fmt.Printf("type:%T value:%#v\n", p1, p1) //type:*main.person value:&main.person{name:"小王子", city:"北京", age:18}

	//指针类型的接收者
	/*fmt.Println(p1.age) // 18
	p1.SetAge(30)
	fmt.Println(p1.age) // 30*/

	//值类型的接收者
	/*fmt.Println(p1.age) // 18
	p1.SetAge2(30)
	fmt.Println(p1.age) // 18*/

	//任意类型添加方法
	/*var m1 MyInt
	m1.SayHello() //Hello, 我是一个int。
	m1 = 100
	fmt.Printf("%#v  %T\n", m1, m1) //100  main.MyInt*/

	//结构体的匿名字段
	/*p2 := person2{
		"小王子",
		18,
	}
	fmt.Printf("%#v\n", p1) //&main.person{name:"小王子", city:"北京", age:18}

	fmt.Println(p2.string, p2.int) //小王子 18*/

	//嵌套结构体
	/*	user1 := User{
			Name:   "小王子",
			Gender: "男",
			Address: Address{
				Province: "山东",
				City:     "威海",
			},
		}
		fmt.Printf("user1=%#v\n", user1) //user1=main.User{Name:"小王子", Gender:"男", Address:main.Address{Province:"山东", City:"威海"}}
	*/
	//嵌套匿名字段
	/*var user2 User2
	user2.Name = "小王子"
	user2.Gender = "男"
	user2.Address.Province = "山东"    // 匿名字段默认使用类型名作为字段名
	user2.City = "威海"                // 匿名字段可以省略
	fmt.Printf("user2=%#v\n", user2) //user2=main.User{Name:"小王子", Gender:"男", Address:main.Address{Province:"山东", City:"威海"}}
	*/
	//嵌套结构体的字段名冲突
	/*var user3 User3
	user3.Name = "沙河娜扎"
	user3.Gender = "男"
	// user3.CreateTime = "2019" //ambiguous selector user3.CreateTime
	user3.Address.CreateTime = "2000" //指定Address结构体中的CreateTime
	user3.Email.CreateTime = "2000"   //指定Email结构体中的CreateTime*/

	//结构体的“继承”

	/*d1 := &Dog{
		Feet: 4,
		Animal: &Animal{ //注意嵌套的是结构体指针
			name: "乐乐",
		},
	}
	d1.say() //乐乐会汪汪汪~
	d1.move() //乐乐会动！*/


	p := Person{name: "小王子", age: 18}
	data := []string{"吃饭", "睡觉", "打豆豆"}
	p.SetDreams(data)

	// 你真的想要修改 p1.dreams 吗？
	data[1] = "不睡觉"
	fmt.Println(p.dreams)  // ?
}
