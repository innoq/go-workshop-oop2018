package main

import (
	"io/ioutil"
	"os"
	"testing"
	"time"
)

// just a test that is formatted incorrectly
func TestHelloTony(t *testing.T) {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	c := make(chan string)

	helloTony(c)

	<-time.After(1 * time.Second)

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	expected := "hello tony\n"
	if string(out) != expected {
		t.Errorf("Fehler: %q != %q", out, expected)
	}
}
