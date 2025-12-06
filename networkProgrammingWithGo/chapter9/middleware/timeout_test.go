package middleware

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestTimeoutMiddleware(t *testing.T) {
	handler := http.TimeoutHandler(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNoContent)
			time.Sleep(time.Minute)
		}),
		time.Second,
		"Time out while reading response",
	)
	r := httptest.NewRequest(http.MethodGet, "http://test/", nil)
	w := httptest.NewRecorder()

	//http.TimeoutHandler返回的是一个实现http.Handler接口的对象，所以不能直接调用，需要对象调用方法
	handler.ServeHTTP(w, r)
	resp := w.Result()
	if resp.StatusCode != http.StatusServiceUnavailable {
		t.Fatalf("unexpected status code: %q", resp.Status)
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	_ = resp.Body.Close()
	if actual := string(b); actual != "Time out while reading response" {
		t.Logf("unexpected body: %q", actual)
	}
}
