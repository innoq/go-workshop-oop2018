package main

import (
	"io/ioutil"
	"os"
	"testing"
)

// just a test that is formatted incorrectly
func TestHelloAustralia(t *testing.T) {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	helloWorld()

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	expected := "hello world\n"
	if string(out) != expected {
		t.Errorf("Fehler: %q != %q", out, expected)
	}
}
