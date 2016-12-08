package day5

import "testing"

var hashCases = []struct {
	door  string
	index int
	want  byte
}{
	{"abc", 3231929, '1'},
	{"abc", 5017308, '8'},
	{"abc", 5278568, 'f'},
}

func TestHash(t *testing.T) {
	for i, c := range hashCases {
		got := hash(c.door, c.index)
		if got != c.want {
			t.Fatalf("Case %d: Wanted %v, got %v", i+1, c.want, got)
		}
	}
}

var hash2Cases = []struct {
	door     string
	index    int
	wantPos  int
	wantByte byte
}{
	{"abc", 3231929, 1, '5'},
	{"abc", 5017308, -1, '_'},
	{"abc", 5357525, 4, 'e'},
}

func TestHash2(t *testing.T) {
	pwd := []byte("________")
	for i, c := range hash2Cases {
		gotPos, gotByte := hash2(c.door, c.index, pwd)
		if gotPos != c.wantPos {
			t.Fatalf("Case %d: wantPos=%d, gotPos=%d", i+1, c.wantPos, gotPos)
		}
		if gotByte != c.wantByte {
			t.Fatalf("Case %d: wantByte=%d, gotByte=%d", i+1, c.wantByte, gotByte)
		}
	}
}
