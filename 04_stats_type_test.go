package main

import "testing"

func TestStatsType(t *testing.T) {
	s := Stats{Cnt{"bing", 28}, Cnt{"flickr", 100}}

	s[0].Total = 0 // dummy operation to avoid compile error of unused var
}
