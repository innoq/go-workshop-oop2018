package main

import "fmt"

func ExampleStatsInc() {
	s := Stats{}

	s.Inc("bing")
	s.Inc("bing")
	s.Inc("bing")
	s.Inc("flickr")

	fmt.Printf("%s", s)

	// Output:
	// bing: 3
	// flickr: 1

}
