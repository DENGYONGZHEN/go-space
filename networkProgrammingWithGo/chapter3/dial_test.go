package chapter3

import (
	"io"
	"net"
	"testing"
)

func TestDial(t *testing.T) {

	listener, err := net.Listen("tcp", "127.0.0.1:")
	if err != nil {
		t.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		defer func() { done <- struct{}{} }()

		//accept incoming TCP  connection in a loop
		for {
			// 阻塞直到新连接到来      解除条件：客户端连接或监听器关闭
			conn, err := listener.Accept()
			if err != nil {
				t.Log(err)
				return
			}

			//spin off each connection into its own gorountine (handler)
			go func(c net.Conn) {
				defer func() {
					c.Close()
					done <- struct{}{}
				}()

				buf := make([]byte, 1024)
				for {
					// 阻塞直到收到数据或连接关闭
					//conn.Close()执行后，会受到一个io.EOF error
					n, err := c.Read(buf)
					if err != nil {
						if err != io.EOF {
							t.Error(err)
						}
						return
					}
					t.Logf("received: %q", buf[:n])
				}

			}(conn)
		}
	}()

	//客户端连接
	conn, err := net.Dial("tcp", listener.Addr().String())
	if err != nil {
		t.Fatal(err)
	}

	//具体信息查看TCP断开连接四次挥手
	conn.Close()
	<-done // 阻塞直到收到信号  来自connection handler
	listener.Close()
	<-done
}
