package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func drainAndClose(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)
			//规范：需要清空request body中的内容，并且关闭
			//另外无论是copy还是read，都是消耗来源中的内容，即读过之后，内容源中不再存在内容
			_, _ = io.Copy(io.Discard, r.Body)
			_ = r.Body.Close()
		},
	)
}

func TestSimpleMux(t *testing.T) {
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})
	serveMux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprint(w, "Hello friend.")
	})
	serveMux.HandleFunc("/hello/there/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprint(w, "Why, hello there.")
	})
	mux := drainAndClose(serveMux)

	testCases := []struct {
		path     string
		response string
		code     int
	}{
		{"http://test/", "", http.StatusNoContent},
		{"http://test/hello", "Hello friend.", http.StatusOK},
		{"http://test/hello/there/", "Why, hello there.", http.StatusOK},
		//请求路径会先进行绝对匹配，如果没有匹配到，会尝试在原本路径后面加/再次进行匹配
		{"http://test/hello/there", "<a href=\"/hello/there/\">Moved Permanently</a>.\n\n", http.StatusMovedPermanently},
		{"http://test/hello/there/you", "Why, hello there.", http.StatusOK},

		{"http://test/hello/and/goodbye", "", http.StatusNoContent},
		{"http://test/something/esle/entirely", "", http.StatusNoContent},
		{"http://test/hello/you", "", http.StatusNoContent},
	}

	for i, c := range testCases {
		r := httptest.NewRequest(http.MethodGet, c.path, nil)
		w := httptest.NewRecorder()
		mux.ServeHTTP(w, r)
		resp := w.Result()

		if actual := resp.StatusCode; c.code != actual {
			t.Errorf("%d: expected %d; actual %d", i, c.code, actual)
		}

		b, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatal(err)
		}
		_ = resp.Body.Close()

		if actual := string(b); c.response != actual {
			t.Errorf("%d: expected response %q; actual %q", i, c.response, actual)
		}
	}
}
