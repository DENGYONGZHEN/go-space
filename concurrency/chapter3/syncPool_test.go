package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"sync"
	"testing"
	"time"
)

//go test -benchtime=10s -bench=. syncPool_test.go

// 模拟网络连接耗时：
// 假装“连接到外部服务”需要 1 秒。
func connectToService() interface{} {
	time.Sleep(1 * time.Second)
	return struct{}{}
}

func warmServiceConnCache() *sync.Pool {
	p := &sync.Pool{
		New: connectToService,
	}
	for i := 0; i < 10; i++ {
		p.Put(p.New())
	}
	return p
}

func startNetworkDaemon() *sync.WaitGroup {
	//sync.WaitGroup 在这里用于确保 服务启动完成再跑测试。
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		connPool := warmServiceConnCache()
		server, err := net.Listen("tcp", "localhost:8080")
		if err != nil {
			log.Fatalf("cannot listen: %v", err)
		}
		defer server.Close()
		wg.Done()

		for {
			conn, err := server.Accept()
			if err != nil {
				log.Printf("cannot accept connection: %v", err)
				continue
			}
			// connectToService()
			svcConn := connPool.Get()
			fmt.Fprintln(conn, "")
			connPool.Put(svcConn)
			conn.Close()
		}
	}()
	return &wg
}

func init() {
	daemonStarted := startNetworkDaemon()
	daemonStarted.Wait()
}

func BenchmarkNetworkRequest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		conn, err := net.Dial("tcp", "localhost:8080")
		if err != nil {
			b.Fatalf("cannot dial host: %v", err)
		}
		if _, err := io.ReadAll(conn); err != nil {
			b.Fatalf("cannot read: %v", err)
		}
		conn.Close()
	}
}
