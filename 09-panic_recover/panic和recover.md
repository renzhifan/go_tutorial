## panic/recover

Go语言中目前（Go1.12）是没有异常机制，但是使用`panic/recover`模式来处理错误。 panic可以在任何地方引发，**但recover只有在defer调用的函数中有效**。 首先来看一个例子：

```go
func funcA() {
fmt.Println("func A")
}

func funcB() {
panic("panic in B")
}

func funcC() {
fmt.Println("func C")
}
func main() {
funcA()
funcB()
funcC()
}
```

输出

```shell
func A
panic: panic in B

goroutine 1 [running]:
main.funcB(...)
        .../panic_recover/main.go:17
main.main()
        .../panic_recover/main.go:26 +0x66


```

程序运行期间funcB中引发了panic导致程序崩溃，异常退出了。 这个时候我们就可以通过recover将程序恢复回来，继续往后执行。

```go
package main

import "fmt"

//panic和recover
func a() {
	fmt.Println("func a")
}

func b() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	panic("panic in b")
}

func c() {
	fmt.Println("func c")
}

func main() {
	a()
	b()
	c()
}

```

输出

```shell
func a
panic in b
func c

```


