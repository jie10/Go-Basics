package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestVarDeclaration(t *testing.T) {
	// Override os.Stdout so we can capture the output instead of printing it directly.
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// call main function here
	main()

	// Reset original stdout after test run
	os.Stdout = rescueStdout

	w.Close()
	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := strings.TrimSpace(buf.String()) // remove any trailing white space

	// Now we can check if the output is what we expect
	if !strings.Contains(output, "Value of i: 0") {
		t.Errorf("Expected output to contain 'Value of i: 0', but got: %s", output)
	}
	if !strings.Contains(output, "Value of j: 10") {
		t.Errorf("Expected output to contain 'Value of j: 10', but got: %s", output)
	}
	if !strings.Contains(output, "Value of k: true") {
		t.Errorf("Expected output to contain 'Value of k: true', but got: %s", output)
	}
}
