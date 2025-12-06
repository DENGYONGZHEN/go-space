package chapter8

import (
	"net/http"
	"testing"
	"time"
)

func TestHeadTime(t *testing.T) {
	resp, err := http.Head("https://www.time.gov/")
	if err != nil {
		t.Fatal(err)
	}

	//如果不关闭 response.Body，Go 的 HTTP 客户端 无法重用连接，并且会 泄露文件描述符，最终可能导致后续请求失败。
	// 当你调用 client.Do(...) 后，会打开一个文件描述符（例如 socket）。若不调用 resp.Body.Close()，这个资源不会被释放，会一直占用。

	//defer resp.Body.Close()

	//io.Copy(io.Discard, resp.Body)  // 可选：读取并丢弃，确保重用连接
	_ = resp.Body.Close() // 无论是否读取，都要关闭
	now := time.Now().Round(time.Second)
	date := resp.Header.Get("Date")
	if date == "" {
		t.Fatal("no Date header received from time.gov")
	}
	dt, err := time.Parse(time.RFC1123, date)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("time.gov: %s (skew %s)", dt, now.Sub(dt))
}
