package main

import (
	"fmt"
)

// map(映射)
func main() {
	//scoreMap := make(map[string]int, 8)
	//scoreMap["张三"] = 90
	//scoreMap["小明"] = 100
	//fmt.Println(scoreMap) //map[小明:100 张三:90]
	//
	//fmt.Println(scoreMap["小明"]) //100
	//
	//fmt.Printf("type of a:%T\n", scoreMap) //type of a:map[string]int
	//
	//userInfo := map[string]string{
	//	"username": "沙河小王子",
	//	"password": "123456",
	//}
	//fmt.Println(userInfo) //map[password:123456 username:沙河小王子]
	//
	//scoreMap := make(map[string]int)
	//scoreMap["张三"] = 90
	//scoreMap["小明"] = 100
	//scoreMap["娜扎"] = 60
	////Go语言中使用for range遍历map。
	//for k, v := range scoreMap {
	//	fmt.Println(k, v)
	//}
	////使用delete()函数删除键值对
	//delete(scoreMap, "小明") //将小明:100从map中删除
	//for k, v := range scoreMap {
	//	fmt.Println(k, v)
	//}

	// 统计一个字符串中每个单词出现的次数
	// "how do you do"中每个单词出现的次数
	// 0. 定义一个map[string]int
	//var s = "how do you do"
	//var wordCount = make(map[string]int, 10)
	// 1. 字符串中都有哪些单词
	//words := strings.Split(s, " ")
	// 2. 遍历单词做统计
	//for _, word := range words {
	//	v, ok := wordCount[word]
	//	if ok {
	//		// map中有这个单词的统计记录
	//		wordCount[word] = v + 1
	//	} else {
	//		// map中没有这个单词的统计记录
	//		wordCount[word] = 1
	//	}
	//}
	//for k, v := range wordCount {
	//	fmt.Println(k, v)
	//}

	//解决数据共享问题
	//slice1 := []int{1, 2, 3, 4, 5}
	//slice2 := slice1[1:3]
	//printSlice(slice1, "slice1")
	//printSlice(slice2, "slice2")
	//slice2 = append(slice1, 6)
	//printSlice(slice2, "slice2")
	//slice2 = append(slice2[:1], slice2[2:]...)
	////slice1[1] = 2
	////slice2[1] = 6
	//printSlice(slice1, "slice1")
	//printSlice(slice2, "slice2")

	//思考题
	//[append() 函数并不会改变原来的切片，而是会生成一个容量更大的切片，然后把原有的元素和新元素一并拷贝到新切片中。]
	type Map map[string][]int
	m := make(Map)
	s := []int{1, 2}
	s = append(s, 3)
	//fmt.Printf("%+v\n", s) //[1 2 3]
	printSlice(s, "s")

	m["key"] = s
	s = append(s[:1], s[2:]...)
	//fmt.Printf("%+v\n", s) //[1 3]
	printSlice(s, "s")

	//fmt.Printf("%+v\n", m["key"]) //[1 3 3]
	printSlice(m["key"], "m[\"key\"]")
	//append() 函数并不会改变原来的切片，而是会生成一个容量更大的切片，然后把原有的元素和新元素一并拷贝到新切片中。

}
//其实%p输出切片地址，很容易搞错，不同切片可能都一样，你得切片变量前面➕&才行，不然切片可能指向同一个内存中的数组，导致地址看起来是一样的
func printSlice(s []int, name string) {
	fmt.Printf("name=%s 值的类型=%T len=%d cap=%d 带0x的指针=%p 不带0x的指针=%#p 输出结构体=%v\n", name, s, len(s), cap(s), &s, &s, s)
}
