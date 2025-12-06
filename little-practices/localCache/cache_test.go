package localcache

import (
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	cache := NewCache(time.Second * 5)

	cache.Set("test", "testString", time.Hour*12)
	v, exist := cache.Get("test")
	if !exist {
		t.Fatal("cache doesnt contain this key value")
	}
	if v != "testString" {
		t.Fatal("the value does not same")
	}
}
