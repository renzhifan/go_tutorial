package main

import "fmt"

// 数组相关内容
func main() {
	// var a [3]int
	// var b [4]int
	// // a = b
	// fmt.Println(a)
	// fmt.Println(b)
	// 数组的初始化
	// 1. 定义时使用初始值列表的方式初始化
	// var cityArray = [4]string{"北京", "上海", "广州", "深圳"}
	// fmt.Println(cityArray)
	// fmt.Println(cityArray[0])
	// fmt.Println(cityArray[3])
	// // 2. 编译器推导数组的长度
	// var boolArray = [...]bool{true, false, false, true, false}
	// fmt.Println(boolArray)

	// // 3. 使用索引值方式初始化
	// var langArray = [...]string{1: "Golang", 3: "Python", 7: "Java"}
	// fmt.Println(langArray)
	// fmt.Printf("%T\n", langArray)

	// 数组的遍历
	// var cityArray = [4]string{"北京", "上海", "广州", "深圳"}
	// 1. for循环遍历
	// for i := 0; i < len(cityArray); i++ {
	// 	fmt.Println(cityArray[i])
	// }
	// 2. for range遍历
	// for _, value := range cityArray {
	// 	fmt.Println(value)
	// }

	// 二维数组
	// cityArray := [...][2]string{
	// 	{"北京", "西安"},
	// 	{"上海", "杭州"},
	// 	{"重庆", "成都"},
	// 	{"广州", "深圳"},
	// }
	// fmt.Println(cityArray)
	// fmt.Println(cityArray[2][0])
	// // 二维数组的遍历
	// for _, v1 := range cityArray {
	// 	fmt.Println(v1)
	// 	for _, v2 := range v1 {
	// 		fmt.Println(v2)
	// 	}
	// }

	// 数组是值类型
	x := [3][2]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	fmt.Println(x)
	f1(x)
	fmt.Println(x)
	y := x
	y[0][0] = 1000
	fmt.Println(x)

	//求数组[1, 3, 5, 7, 8]所有元素的和
	var intArr = [...]int{1, 3, 5, 7, 8}
	// 方法一
	sum := 0
	for i := 0; i < len(intArr); i++ {
		sum += intArr[i]
		if i == len(intArr)-1 {
			fmt.Println(sum)
		}
	}
	// 方法二
	sum1 := 0
	for _, value := range intArr {
		sum1 += value
	}
	fmt.Println(sum1)

	//找出数组中和为指定值的两个元素的下标，比如从数组[1, 3, 5, 7, 8]中找出和为8的两个元素的下标分别为(0,3)和(1,2)。
	res := 8
	for i := 0; i < len(intArr); i++ {
		for j := i + 1; j < len(intArr); j++ {
			fmt.Printf("下标:%d,%d;对应的值：%d,%d\n", i, j, intArr[i], intArr[j])
			if intArr[i]+intArr[j] == res {
				fmt.Printf("%d+%d=%d,下标为(%d,%d)\n", intArr[i], intArr[j], intArr[i]+intArr[j], i, j)
			}
		}
	}

}

func f1(a [3][2]int) {
	a[0][0] = 100
}
