package day3

import "testing"

var validCases = []struct {
	a, b, c int
	want    bool
}{
	{5, 10, 25, false},
	{10, 11, 12, true},
	{10, 10, 10, true},
	{100, 50, 60, true},
}

func TestValid(t *testing.T) {
	for i, c := range validCases {
		got := isValid(c.a, c.b, c.c)
		if got != c.want {
			t.Fatalf("Case %d: Wanted %v, got %v\n", i+1, c.want, got)
		}
	}
}

var numValidCases = []struct {
	input string
	want  int
}{
	{"5 10 25\n10 11 12\n10 10 10\n100 50 60", 3},
}

func TestNumValid(t *testing.T) {
	for i, c := range numValidCases {
		got, _ := NumberOfValid(c.input)
		if got != c.want {
			t.Fatalf("Case %d: Wanted %v, got %v\n", i+1, c.want, got)
		}
	}
}
