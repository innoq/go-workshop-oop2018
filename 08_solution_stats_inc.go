package main

func (s *Stats) Inc(source string) {
	for i, cnt := range *s {
		if cnt.Source == source {
			(*s)[i].Total = cnt.Total + 1
			return
		}
	}

	*s = append(*s, Cnt{source, 1})
}
