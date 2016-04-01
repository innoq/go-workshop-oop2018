package main

import (
	"os/exec"
	"testing"
)

// just a test that is formatted incorrectly
func TestFormatting(t *testing.T) {
	cmd := exec.Command("gofmt", "-l", "-e", ".")
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Errorf("Fehler beim ausführen von `gofmt`: %s", err.Error())
	}

	if len(out) > 0 {
		t.Errorf("folgende Dateien sind nicht korrekt formatiert: %s", out)
	}
}
