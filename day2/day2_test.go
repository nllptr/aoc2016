package day2

import "testing"

var lookupKeyCases = []struct {
	key   byte
	pad   keypad
	wantX int
	wantY int
}{
	{'5', keyPad1, 2, 2},
	{'7', keyPad1, 1, 3},
}

func TestLookupKey(t *testing.T) {
	for i, c := range lookupKeyCases {
		gotX, gotY := lookupKey(c.pad, c.key)
		if gotX != c.wantX && gotY != c.wantY {
			t.Fatalf("Case %d: Wanted (%d, %d), got (%d, %d)", i+1, c.wantX, c.wantY, gotX, gotY)
		}
	}
}

var processKeyCases = []struct {
	pad      keypad
	startKey byte
	command  byte
	want     byte
}{
	{keyPad1, '5', 'U', '2'},
	{keyPad1, '5', 'D', '8'},
	{keyPad1, '5', 'L', '4'},
	{keyPad1, '5', 'R', '6'},
	{keyPad1, '2', 'U', '2'},
	{keyPad1, '8', 'D', '8'},
	{keyPad1, '4', 'L', '4'},
	{keyPad1, '6', 'R', '6'},
}

func TestProcessKey(t *testing.T) {
	curX := new(int)
	curY := new(int)
	for i, c := range processKeyCases {
		*curX, *curY = lookupKey(c.pad, c.startKey)
		got := processKey(c.pad, curX, curY, c.command)
		if got != c.want {
			t.Fatalf("Case %d: Wanted %v, got %v", i+1, string(c.want), string(got))
		}
	}
}

var processLineCases = []struct {
	pad      keypad
	line     string
	startKey byte
	want     byte
}{
	{keyPad1, "UURRLRDDD", '5', '9'},
	{keyPad1, "LLLUUURRRDDDULDLURRD", '9', '9'},
}

func TestProcessLine(t *testing.T) {
	curX := new(int)
	curY := new(int)
	for i, c := range processLineCases {
		*curX, *curY = lookupKey(c.pad, c.startKey)
		got := processLine(c.pad, curX, curY, c.line)
		if got != c.want {
			t.Fatalf("Case %d: Wanted %v, got %v", i+1, string(c.want), string(got))
		}
	}
}
