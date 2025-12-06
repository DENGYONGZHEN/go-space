package chapter4

import (
	"bufio"
	"net"
	"reflect"
	"testing"
)

const payload = "The bigger the interface, the weaker the abstraction"

func TestScanner(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:")
	if err != nil {
		t.Fatal(err)
	}

	go func() {
		conn, err := listener.Accept()
		if err != nil {
			t.Error(err)
			return
		}
		defer conn.Close()
		_, err = conn.Write([]byte(payload))
		if err != nil {
			t.Error(err)
		}
	}()

	conn, err := net.Dial("tcp", listener.Addr().String())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	scanner.Split(bufio.ScanWords)

	var words []string

	for scanner.Scan() { // Every call to Scan result in multiple calls to the network connection's Read method until
		//the scanner finds its delimeter or reads an error from the connection
		words = append(words, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		t.Error(err)
	}

	expected := []string{"The", "bigger", "the", "interface,", "the", "weaker", "the", "abstraction."}
	if reflect.DeepEqual(words, expected) {
		t.Fatal("inaccurate scanned word list")
	}
	t.Logf("Scanned words: %#v", words)
}
