package day1

import "testing"

var moveCases = []struct {
	commands []string
	want     position
}{
	{
		[]string{"R2"},
		position{
			east,
			2,
			0,
		},
	},
	{
		[]string{"R2", "L3"},
		position{
			north,
			2,
			3,
		},
	},
	{
		[]string{"R2", "R2", "R2"},
		position{
			west,
			0,
			-2,
		},
	},
	{
		[]string{"R5", "L5", "R5", "R3"},
		position{
			south,
			10,
			2,
		},
	},
}

func TestMove(t *testing.T) {
	for _, c := range moveCases {
		var pos position
		pos = position{north, 0, 0}
		for _, cmd := range c.commands {
			err := move(&pos, cmd, nil)
			if err != nil {
				t.Fatal("TestMove: move returned an error: ", err)
			}
		}
		if pos.dir != c.want.dir {
			t.Fatalf("TestMode: Facing wrong direction. Want %v, got %v\n", c.want.dir, pos.dir)
		}
		if pos.x != c.want.x || pos.y != c.want.y {
			t.Fatalf("TestMove: Wrong coordinates. Want (%d, %d), got (%d, %d)\n", c.want.x, c.want.y, pos.x, pos.y)
		}
	}
}

var distanceCases = []struct {
	commands string
	want     int
}{
	{
		"R2,R2, R2",
		2,
	},
	{
		"R5, L5, R5, R3",
		12,
	},
}

func TestDistance(t *testing.T) {
	for _, c := range distanceCases {
		got := Distance(c.commands)
		if got != c.want {
			t.Fatalf("Got %d, wanted %d\n", got, c.want)
		}
	}
}

var revisitedCases = []struct {
	commands string
	want     int
}{
	{
		"R8, R4, R4, R8",
		4,
	},
}
