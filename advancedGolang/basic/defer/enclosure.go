package main

import "fmt"

//闭包捕获的是内存地址

func Accumulator() func(int) int {
	var x int
	return func(delta int) int {
		fmt.Printf("(%+v,%+v)-", &x, x)
		x += delta
		return x
	}
}

//闭包引用了x变量，a,b可以看作两个不同的实例，实例之间互不影响。而实例内部，X变量是同一个地址

// func main() {

// 	var a = Accumulator()
// 	fmt.Printf("%d\n", a(1))
// 	fmt.Printf("%d\n", a(10))
// 	fmt.Printf("%d\n", a(100))

// 	fmt.Println("----------------------")

// 	var b = Accumulator()
// 	fmt.Printf("%d\n", b(1))
// 	fmt.Printf("%d\n", b(10))
// 	fmt.Printf("%d\n", b(100))
// }
