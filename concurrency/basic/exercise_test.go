package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

var testString string

func Test_updateMessage(t *testing.T) {

	testString = "testUpdateMessages"
	wg.Add(1)
	updateMessages(testString)
	wg.Wait()
	if testString != msg {
		t.Errorf("expected same string, but not")
	}
}

func Test_printMessage(t *testing.T) {

	testString = "testUpdateMessages"
	wg.Add(1)
	updateMessages(testString)
	wg.Wait()

	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	printMessage()
	_ = w.Close()
	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut
	if !strings.Contains(output, testString) {
		t.Errorf("expected same string :%s, but not:%s", output, testString)
	}

}

func Test_exercise(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()

	os.Stdout = w

	exercise()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut
	if !strings.Contains(output, "Hello, universe!") {
		t.Errorf("expected same string :%s, but not:%s", output, "Hello, universe!")
	}
	if !strings.Contains(output, "Hello, cosmos!") {
		t.Errorf("expected same string :%s, but not:%s", output, "Hello, cosmos!")
	}
	if !strings.Contains(output, "Hello,world!") {
		t.Errorf("expected same string :%s, but not:%s", output, "Hello,world!")
	}
}
