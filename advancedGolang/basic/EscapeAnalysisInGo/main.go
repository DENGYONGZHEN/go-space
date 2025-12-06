package main

import "fmt"

// 逃逸分析是编译器在编译时进行的一种静态分析，
// 用于确定变量的生命周期是否会超出其定义的作用域（即是否会“逃逸”到堆上）。
// 如果变量逃逸到堆上，就需要进行堆内存分配，而不是栈内存分配。
// 逃逸分析的目的是优化内存分配，减少垃圾回收的压力。

// 栈分配：如果变量不会逃逸，编译器会将其分配在栈上，栈分配的内存会在函数返回时自动释放。

// 堆分配：如果变量逃逸，编译器会将其分配在堆上，堆分配的内存需要由垃圾回收器（GC）来管理。

//go build -gcflags '-m -l' main.go

// 步骤 1：生成汇编代码
// # 输出 main.go 的汇编代码到控制台
// go build -gcflags=-S main.go

// # 将汇编代码保存到文件
// go build -gcflags=-S main.go 2> main.s

// 步骤 2：查看详细编译日志
// go build -x -work main.go

func foo() *int {
	t := 3
	return &t
}

func main() {

	x := foo()
	fmt.Println(*x)
}
