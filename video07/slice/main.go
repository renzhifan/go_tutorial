package main

import (
	"fmt"
	"sort"
)

func main() {
	//切片的定义
	//var name []T

	// 声明切片类型
	//var a []string              //声明一个字符串切片
	//var b = []int{}             //声明一个整型切片并初始化
	//var c = []bool{false, true} //声明一个布尔切片并初始化
	//var d = []bool{false, true} //声明一个布尔切片并初始化
	//fmt.Println(a)              //[]
	//fmt.Println(b)              //[]
	//fmt.Println(c)              //[false true]
	//fmt.Println(a == nil)       //true
	//fmt.Println(b == nil)       //false
	//fmt.Println(c == nil)       //false
	//fmt.Println(c == d)   //切片是引用类型，不支持直接比较，只能和nil比较

	//a := [5]int{1, 2, 3, 4, 5}
	//s := a[1:3]        // s := a[low:high]
	//printSlice(s, "s") //name=s 值的类型=[]int  len=2 cap=4 带0x的指针=0xc0000b6008 不带0x的指针=c0000b6008 输出结构体=[2 3]
	//
	//s2 := s[3:4]         // 索引的上限是cap(s)而不是len(s)
	//printSlice(s2, "s2") //name=s2 值的类型=[]int  len=1 cap=1 带0x的指针=0xc00001a0e0 不带0x的指针=c00001a0e0 输出结构体=[5]
	//
	//t := a[1:3:5]
	//printSlice(t, "t") //name=t 值的类型=[]int len=2 cap=4 带0x的指针=0xc00001a0c8 不带0x的指针=c00001a0c8 输出结构体=[2 3]

	//使用make()函数构造切片
	//make([]T, size, cap) T:切片的元素类型,size:切片中元素的数量,cap:切片的容量
	//b := make([]int, 2, 10)
	//printSlice(b, "b") //name=b 值的类型=[]int len=2 cap=10 带0x的指针=0xc0000b4050 不带0x的指针=c0000b4050 输出结构体=[0 0]
	//
	//c := [8]int{0, 1, 2, 3, 4, 5, 6, 7}
	//s3 := c[:5]
	//printSlice(s3, "s3") //name=s3 值的类型=[]int len=5 cap=8 带0x的指针=0xc000020140 不带0x的指针=c000020140 输出结构体=[0 1 2 3 4]

	//切片不能直接比较
	//var s4 []int
	//s5 := []int{}
	//s6 := make([]int, 0)
	//isNil(s4, "s4")      //s4==nil
	//isNil(s5, "s5")      //s5!=nil
	//isNil(s6, "s6")      //s6!=nil
	//printSlice(s4, "s4") //name=s4 值的类型=[]int len=0 cap=0 带0x的指针=0x0 不带0x的指针=0 输出结构体=[]
	//
	//printSlice(s5, "s5") //name=s5 值的类型=[]int len=0 cap=0 带0x的指针=0x116de80 不带0x的指针=116de80 输出结构体=[]
	//
	//printSlice(s6, "s6") //name=s6 值的类型=[]int len=0 cap=0 带0x的指针=0x116de80 不带0x的指针=116de80 输出结构体=[]

	//切片的赋值拷贝
	//s7 := make([]int, 3) //[0 0 0]
	//s8 := s7             //将s7直接赋值给s8，s7和s8共用一个底层数组
	//s8[0] = 100
	//printSlice(s7, "s7") //name=s7 值的类型=[]int len=3 cap=3 带0x的指针=0xc0000160a8 不带0x的指针=c0000160a8 输出结构体=[100 0 0]
	//
	//printSlice(s8, "s8") //name=s8 值的类型=[]int len=3 cap=3 带0x的指针=0xc0000160a8 不带0x的指针=c0000160a8 输出结构体=[100 0 0]

	//append()方法为切片添加元素
	//var s9 []int
	//s9 = append(s9, 1)
	//fmt.Printf("s9:%v\n", s9) //s9:[1]
	//
	//s9 = append(s9, 2, 3, 4)
	//fmt.Printf("s9:%v\n", s9) //s9:[1 2 3 4]
	//
	//s10 := []int{5, 6, 7}
	//s9 = append(s9, s10...)
	//fmt.Printf("s9:%v\n", s9) //s9:[1 2 3 4 5 6 7]

	//append()添加元素和切片扩容
	//var numSlice []int
	//for i := 0; i < 10; i++ {
	//	numSlice = append(numSlice, i)
	//	fmt.Printf("%v  len:%d  cap:%d  ptr:%p\n", numSlice, len(numSlice), cap(numSlice), numSlice)
	//}

	//copy()复制切片 copy()函数可以迅速地将一个切片的数据复制到另外一个切片空间中
	//a := []int{1, 2, 3, 4, 5}
	//c := make([]int, 5, 5)
	//copy(c, a)     //使用copy()函数将切片a中的元素复制到切片c
	//fmt.Println(a) //[1 2 3 4 5]
	//fmt.Println(c) //[1 2 3 4 5]
	//c[0] = 1000
	//fmt.Println(a) //[1 2 3 4 5]
	//fmt.Println(c) //[1000 2 3 4 5]

	//从切片中删除元素
	//a := []int{30, 31, 32, 33, 34, 35, 36, 37}
	//// 要删除索引为2的元素
	//a = append(a[:2], a[3:]...)//>总结一下就是：要从切片a中删除索引为index的元素，操作方法是a = append(a[:index], a[index+1:]...)
	//fmt.Println(a) //[30 31 33 34 35 36 37]

	//练习题
	//var a = make([]string, 5, 10)
	//for i := 0; i < 10; i++ {
	//	a = append(a, fmt.Sprintf("%v", i))
	//}
	//printSlice(a, "a") //name=a 值的类型=[]string len=15 cap=20 带0x的指针=0xc00010a000 不带0x的指针=c00010a000 输出结构体=[     0 1 2 3 4 5 6 7 8 9]

	var a = [...]int{3, 7, 8, 9, 1}
	var s []int
	s = a[0:5]
	sort.Ints(s)
	printSlice(s, "s")
	fmt.Printf("a=%v\n", a)

}
//其实%p输出切片地址，很容易搞错，不同切片可能都一样，你得切片变量前面➕&才行，不然切片可能指向同一个内存中的数组，导致地址看起来是一样的
func printSlice(s []int, name string) {
	fmt.Printf("name=%s 值的类型=%T len=%d cap=%d 带0x的指针=%p 不带0x的指针=%#p 输出结构体=%v\n", name, s, len(s), cap(s), &s, &s, s)
}
func isNil(arg []int, name string) {
	if arg == nil {
		fmt.Printf("%s==nil\n", name)
	} else {
		fmt.Printf("%s!=nil\n", name)
	}
}
