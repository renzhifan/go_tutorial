package main

import (
	"errors"
	"fmt"
)

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

//	//函数可以作为参数：
func calc(x, y int, op func(int, int) int) int {
	return op(x, y)
}

//函数也可以作为返回值：

func do(s string) (func(int, int) int, error) {
	switch s {
	case "+":
		return add, nil
	case "-":
		return sub, nil
	default:
		err := errors.New("无法识别的操作符")
		return nil, err
	}
}
func main() {
	r1 := calc(100, 200, add)
	fmt.Println(r1) //300
	r2 := calc(100, 200, sub)
	fmt.Println(r2) //-100
	r3, r4 := do("+")
	fmt.Printf("r3.type=%T,r3=%d,r4.type=%T,r4=%v\n", r3, r3(1, 2), r4, r4) //r3.type=func(int, int) int,r3=3,r4.type=<nil>,r4=<nil>

}
