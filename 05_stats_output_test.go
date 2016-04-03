package main

import "fmt"

func ExampleStats_String() {
	s := Stats{
		{"bing", 20},
		{"flickr", 100},
	}.String()

	fmt.Print(s)

	// Output:
	// bing: 20
	// flickr: 100
}
