package main

import "fmt"

func (stats Stats) String() string {
	s := ""

	for _, c := range stats {
		s = fmt.Sprintf("%s%s: %d\n", s, c.Source, c.Total)
	}

	return s
}
