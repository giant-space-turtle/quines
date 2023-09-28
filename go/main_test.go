package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestQuine(t *testing.T) {
	want, err := os.ReadFile("./main.go")
	if err != nil {
		t.Fail()
		return
	}
	got := pipeStdout()
	if string(want) != got {
		t.Fatal("Not a quine")
	}
}

func pipeStdout() string {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	outC := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	main()
	w.Close()
	os.Stdout = oldStdout
	return <-outC
}
