package main

import "fmt"

//逃逸分析
// go build -gcflags -m ./main.go

//“逃逸”（escape）是指：变量原本应该分配在栈上，但因为某些原因，被编译器分配到了堆上。
//栈：速度快，函数返回时自动回收。
//堆：需要 GC 管理，分配和回收都比较慢。
//所以编译器在编译时会判断：“这个变量是否可能在函数外部被引用？”   如果是，就必须放到堆上（逃逸）。

func main() {}

// 指针，slice和map作为返回值
// 当带有指针的返回值被赋给外部变量或者作为参数传递给其他函数时，编译器无法确定该变量何时停止使用，
// 因此，为了确保安全性和正确性，它必须将该数据分配到堆上，并使其逃离当前函数作用域
func f1() (*int, []int, map[int]int) {

	i := 0
	list := []int{1, 2, 3, 4}
	mp := map[int]int{1: 1, 2: 2}
	return &i, list, mp
}

// 向chan中发送数据的指针或者包含指针的值
// 编译器此时不知道什么时候会被接收，因此只能放入堆中
func f2() {
	i := 2
	ch := make(chan *int, 2)
	ch <- &i
	<-ch
}

// 非直接的函数调用， 比如在闭包中引用包外的值，因为闭包执行的生命周期可能会超过函数周期，因此需要放入堆中
func f3() func() {
	i := 1
	return func() {
		fmt.Println(i)
	}
}

// 在 interface{} 中传递值
func baz() interface{} {
	x := 10
	return x
}

// 在slice或map中存储指针或者包含指针的值
// slice和map都需要动态分配内存来保存数据，当我们将一个指针或包含指针的值放入slice或map时，编译器无所确定该指针所引用的数据是否会在函数返回后仍然被使用
// 为了保证数据的有效性，编译器会将其分配到堆上，以便在函数返回后继续存在
func f4() {
	i := 1
	list := make([]*int, 10)
	list[0] = &i
}

// interface类型多态的应用，可能会导致逃逸
// 由于接口类型可以持有任意实现了该接口的类型，编译器在编译时无法确定具体的动态类型。
// 因此，为了保证程序正确性， 在运行时需要将接口对象分配到堆上
func f5() {
	var a animal = dog{}
	a.run()

	var al animal
	al = dog{}
	al.run()
}

type animal interface {
	run()
}

type dog struct{}

func (d dog) run() {}
