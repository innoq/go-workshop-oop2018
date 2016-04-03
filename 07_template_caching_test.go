package main

import "testing"

func TestTemplateCaching(t *testing.T) {
	if indexTemplate == nil {
		t.Errorf("die Variable `indexTemplate` ist nil")
	}
}
