package day6

import "testing"

var correctCases = []struct {
	input string
	want  string
}{
	{
		"eedadn\ndrvtee\neandsr\nraavrd\natevrs\ntsrnev\nsdttsa\nrasrtv\nnssdts\nntnada\nsvetve\ntesnvt\nvntsnd\nvrdear\ndvrsen\nenarar",
		"easter",
	},
}

func TestCorrect(t *testing.T) {
	for i, c := range correctCases {
		got := Correct(c.input)
		if got != c.want {
			t.Fatalf("Case %d: Wanted %v, got %v", i+1, c.want, got)
		}
	}
}

var leastCorrectCases = []struct {
	input string
	want  string
}{
	{
		"eedadn\ndrvtee\neandsr\nraavrd\natevrs\ntsrnev\nsdttsa\nrasrtv\nnssdts\nntnada\nsvetve\ntesnvt\nvntsnd\nvrdear\ndvrsen\nenarar",
		"advent",
	},
}

func TestLeastCorrect(t *testing.T) {
	for i, c := range leastCorrectCases {
		got := LeastCorrect(c.input)
		if got != c.want {
			t.Fatalf("Case %d: Wanted %v, got %v", i+1, c.want, got)
		}
	}
}
