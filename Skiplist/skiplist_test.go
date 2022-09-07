package skiplist_test

import (
	skiplist "SkipList/Skiplist"
	"testing"
)

func TestMake(t *testing.T) {
	var skip = skiplist.Make[int, string](6)

	var want = "abc"
	skip.Insert(5, want)
	var got = skip.Find(5)

	if want != *got {
		t.Errorf("want %s, got %s", want, *got)
	}
}

func TestSkiplist_Insert(t *testing.T) {
	var skip = skiplist.Make[string, float64](4)

	pi, g, e := 3.14, 9.81, 2.72
	skip.Insert("pi", pi)
	skip.Insert("g", g)
	skip.Insert("e", e)

	var piFound = *skip.Find("pi")
	if piFound != pi {
		t.Errorf("want %f, got %f", pi, piFound)
	}

	var gFound = *skip.Find("g")
	if gFound != g {
		t.Errorf("want %f, got %f", g, gFound)
	}

	var eFound = *skip.Find("e")
	if eFound != e {
		t.Errorf("want %f, got %f", e, eFound)
	}
}
