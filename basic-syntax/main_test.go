package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestMainPrint(t *testing.T) {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main() // call the function that you are testing

	// read from pipe and copy to a buffer
	_ = w.Close()
	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)

	// now check if output is correct
	if buf.String() != "Hello, World!\n" {
		t.Errorf("Expected 'Hello, World!', got '%s'", buf.String())
	}

	os.Stdout = rescueStdout // restore original stdout
}
