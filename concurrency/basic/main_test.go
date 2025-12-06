package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_printSomething(t *testing.T) {

	//将原始的标准输出（通常是终端）保存到 stdOut 变量中
	stdOut := os.Stdout
	//os.Pipe() 创建一个管道，返回读取端 r 和写入端 w。
	//将标准输出重定向到管道的写入端 w，这样所有写入 os.Stdout 的内容都会被重定向到管道中。
	r, w, _ := os.Pipe()
	os.Stdout = w

	var wg sync.WaitGroup
	wg.Add(1)

	go printSomething("epislon", &wg)
	wg.Wait()

	//关闭管道的写入端 w，表示不会再写入数据。
	_ = w.Close()

	//io.ReadAll(r) 从管道的读取端 r 读取所有数据，返回字节切片 result。
	result, _ := io.ReadAll(r)
	output := string(result)

	//将标准输出恢复为原来的值（通常是终端）。
	os.Stdout = stdOut

	if !strings.Contains(output, "epislon") {
		t.Errorf("Expected to find epsilon, but it is not there")
	}
}
