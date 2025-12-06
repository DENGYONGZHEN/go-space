package main

import "fmt"

type number int

func (n number) print()   { fmt.Println(n) }
func (n *number) pprint() { fmt.Println(*n) }

// func main() {
// f, err := os.Open("file.txt")
// if err != nil {
// 	panic(err)
// }
// if f != nil {
// 	defer f.Close()
// }

//------------->
// var val [3]struct{}

//Go 1.22 修复了循环变量共享问题[1]，新行为：
// 每次迭代创建新变量：for 循环的每次迭代会为 i 生成一个新的实例。
// 闭包隔离性：闭包捕获的是当前迭代的 i 值，而非最终值。
// for i := range val {

//1.压栈：遇到defer，Go会将函数调用(含参数)压入栈中，但不会立即执行
//2.执行：函数返回时，按照先进后出的顺序执行栈中的defer调用
//3.变量求值：
// 1）闭包捕获的是变量的引用(内存的地址)
// 2）因为1.24在循环时，每个i值的地址都是独立的，所以和显示复制的闭包捕获是一样的
// 3）而函数参数传递是值传递，立即复制的

//1.闭包捕获
// defer func() {
// 	fmt.Println(i)
// }()

//2.显示复制的闭包捕获
// i := i
// defer func() {
// 	fmt.Println(i)
// }()

//3.函数参数传递
// defer func(i int) {
// 	fmt.Println(i)
// }(i)

// }
//------------->
// var n number

// defer n.print()  //0  引用的是值
// defer n.pprint() //3  引用的是指针(内存地址)
// defer func() {
// 	n.print() //3  闭包 引用地址
// }()
// defer func() {
// 	n.pprint() //3  闭包 引用地址
// }()
// n = 3
//------------->

// defer func() {
// 	fmt.Println("before return")
// }()

// //return 之后的defer函数不能被注册
// if true {
// 	fmt.Println("during return")
// 	return
// }

// defer func() {
// 	fmt.Println("after return")
// }()
//------------->
// fmt.Println(f())
// }

//***例子1
// func f() (r int) {
// 	t := 5
// 	defer func() {
// 		t = t + 5
// 	}()
// 	return t
// }

// 拆解后-->
// func f() (r int) {
// 	t := 5
// 	//1.赋值指令
// 	r = t
// 	//2.defer被插入到赋值与返回之间执行，这个例子中返回值r没被修改过
// 	defer func() {
// 		t = t + 5
// 	}()

// 	//3.空的return指令
// 	return
// }

//&&&&例子2

// func f() (r int) {
// 	defer func(r int) {
// 		r = r + 5
// 	}(r)
// 	return 1
// }

// 拆解后
// func f() (r int) {
// 	//1.赋值
// 	r = 1
// 	//2.这里改的r是之前传进去的r，不会改变要返回的那个r值
// 	func(r int) {
// 		r = r + 5
// 	}(r)

// 	//3.空的return
// 	return
// }
