package day5

import "testing"

var hashCases = []struct {
	door  string
	index int
	want  string
}{
	{"abc", 3231929, "1"},
	{"abc", 5017308, "8"},
	{"abc", 5278568, "f"},
}

func TestHash(t *testing.T) {
	for i, c := range hashCases {
		got := hash(c.door, c.index)
		if got != c.want {
			t.Fatalf("Case %d: Wanted %v, got %v", i+1, c.want, got)
		}
	}
}
