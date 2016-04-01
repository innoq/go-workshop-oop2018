package main

import "testing"

func TestImageType(t *testing.T) {
	c := Cnt{Source: "bing", Total: 28}

	c.Source = "bing" // dummy operation to avoid compile error of unused var
}
