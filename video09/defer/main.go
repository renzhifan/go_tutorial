package main

import "fmt"

//defer:延迟执行
func f1() int {
	x := 5
	defer func() {
		fmt.Printf("f1:xxxx=%d\n", x)
		x++
		fmt.Printf("f1:xxx=%d\n", x)

	}()
	defer func() {
		x++
		fmt.Printf("f1:xx=%d\n", x)

	}()
	fmt.Printf("f1:x=%d\n", x)
	return x
}

func f2() (x int) {
	defer func() {
		fmt.Printf("f2:xx=%d\n", x)
		x++
		fmt.Printf("f2:xxx=%d\n", x)
	}()
	fmt.Printf("f2:x=%d\n", x)
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
		fmt.Printf("x=%d\n", x)
	}()
	fmt.Printf("x=%d\n", x)
	return x
}
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}

func main() {
	/*fmt.Println("start...")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("end...")
	结果
	start...
	end...
	3
	2
	1*/

	//fmt.Println(f1())
	//fmt.Println(f2())
	//fmt.Println(f3())
	//fmt.Println(f4())

	/*x := 1
	y := 2
	defer calc("AA", x, calc("A", x, y))
	x = 10
	defer calc("BB", x, calc("B", x, y))
	y = 20
	//结果
	A 1 2 3
	B 10 2 12
	BB 10 12 22
	AA 1 3 4*/

}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}
