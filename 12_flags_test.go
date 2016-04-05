package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestFlags(t *testing.T) {
	cmd := exec.Command("bash", "-xc", "go build -o 11.out && ./11.out -h")
	out, _ := cmd.CombinedOutput()

	if strings.Index(string(out), "-timeout duration") < 0 {
		t.Errorf("error: main.go does not implement timeout flag of type duration")
	}
}
