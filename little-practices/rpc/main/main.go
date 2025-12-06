package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"rpc"
	"rpc/codec"
	"time"
)

func main() {
	log.SetFlags(0)
	addr := make(chan string)
	go startServer(addr)

	// in fact ,following code is like a simple rpc client
	conn, _ := net.Dial("tcp", <-addr)
	defer func() { _ = conn.Close() }()

	time.Sleep(time.Second)

	//send option
	_ = json.NewEncoder(conn).Encode(rpc.DefaultOption)
	cc := codec.NewGobCodec(conn)

	//send request & receive response
	for i := range 5 {
		h := &codec.Header{
			ServiceMethod: "Foo.Sum",
			Seq:           uint64(i),
		}
		_ = cc.Write(h, fmt.Sprintf("rpc req %d", h.Seq))
		_ = cc.ReadHeader(h)

		var reply string
		_ = cc.ReadBody(&reply)
		log.Println("reply:", reply)
	}

}

func startServer(addr chan string) {
	//pick a free port
	lis, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatal("network error:", err)
	}
	log.Println("start rpc server on", lis.Addr())
	addr <- lis.Addr().String()
	rpc.Accept(lis)
}
