package main

import (
	"testing"
	"time"
)

func TestMetaSearch(t *testing.T) {
	i, s, err := MetaImageSearch("car", 0)

	if err == nil {
		t.Errorf("timeout von 0 sollte Fehler schmeißen")
	}

	if len(i) != 0 {
		t.Errorf("list of images should be empty with a timeout of 0")
	}

	i, s, err = MetaImageSearch("car", 10*time.Second)

	if len(i) == 0 {
		t.Errorf("list of images should not be empty with a timeout of 10 seconds")
	}

	if len(s) != 3 {
		t.Errorf("stats should have 3 entries")
	}

	if s[0].Total == 0 {
		t.Errorf("%s should have returned at least one result", s[0].Source)
	}

	if s[1].Total == 0 {
		t.Errorf("%s should have returned at least one result", s[1].Source)
	}

	if s[2].Total == 0 {
		t.Errorf("%s should have returned at least one result", s[2].Source)
	}

}
