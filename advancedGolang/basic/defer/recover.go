package main

import "fmt"

//panic 会停掉当前正在执行的程序，而不只是当前线程。
//在这之前，它会有序地执行完当前线程defer列表里的语句，其他协程里定义的defer语句不做保证。
//所以在defer里定义一个recover语句，防止程序直接挂掉。就可以起到类似Java里try..catch的效果

//注意：recover()函数只在defer的函数中直接调用才有效

// func main() {
// 	defer fmt.Println("defer main")
// 	var user = os.Getenv("USER_")

// 	go func() {
// 		defer func() {
// 			fmt.Println("defer caller")
// 			if err := recover(); err != nil {
// 				fmt.Println("recover success. err: ", err)
// 			}
// 		}()

// 		func() {
// 			defer func() {
// 				fmt.Println("defer here")
// 			}()

// 			if user == "" {
// 				panic("should set user env.")
// 			}

// 			//此处不会执行
// 			fmt.Println("after panic")
// 		}()
// 	}()

// 	//时间太短的话，goroutine还没来得及执行，整个程序就退出了
// 	time.Sleep(10000)
// 	fmt.Println("end of main function")
// }

// recover 函数调用的位置
func main() {
	//在 defer函数中调用，有效
	defer f()
	//有效
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("recover")
		}
	}()
	//有效
	defer func() {
		recover()
	}()
	//无效
	recover()
	// 无效
	defer recover()
	//无效，多重defer嵌套
	defer func() {
		defer func() {
			recover()
		}()
	}()

	panic(404)
}

func f() {
	if e := recover(); e != nil {
		fmt.Println("recover")
		return
	}
}
