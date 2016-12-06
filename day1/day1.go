package day1

import (
	"fmt"
	"strconv"
	"strings"
)

type position struct {
	dir  direction
	x, y int
}

type direction int

type memoryMap map[int]map[int]bool

const (
	north = iota
	east
	south
	west
)

func move(pos *position, cmd string, memory memoryMap) error {
	var turn direction
	switch cmd[:1] {
	case "R":
		turn = 1
	case "L":
		turn = -1
	default:
		return fmt.Errorf("First letter of command must be R or L")
	}
	steps, err := strconv.Atoi(cmd[1:])
	if err != nil {
		return fmt.Errorf("Cannot convert steps from string to int (cmd=%s)", cmd)
	}
	pos.dir += turn
	if pos.dir < 0 {
		pos.dir += 4
	}
	pos.dir = pos.dir % 4
	var dim *int
	var sign int
	switch pos.dir {
	case north:
		dim = &pos.y
		sign = 1
	case east:
		dim = &pos.x
		sign = 1
	case south:
		dim = &pos.y
		sign = -1
	case west:
		dim = &pos.x
		sign = -1
	}
	for i := 0; i < steps; i++ {
		if memory != nil {
			if ymap, ok := memory[pos.x]; ok {
				if _, ok := ymap[pos.y]; ok {
					return nil // Stop walking
				}
			} else {
				memory[pos.x] = make(map[int]bool)
			}
			memory[pos.x][pos.y] = true
		}
		*dim += sign
	}
	return nil
}

// Distance returns the Taxicab-geometric distance from the starting point to the
// point where the commands lead.
func Distance(commands string) int {
	pos := position{north, 0, 0}
	for _, c := range strings.Split(commands, ",") {
		move(&pos, strings.TrimSpace(c), nil)
	}
	return abs(pos.x) + abs(pos.y)
}

func FirstRevisitedDistance(commands string) int {
	pos := position{north, 0, 0}
	visited := make(memoryMap)
	//visited[pos.x] = make(map[int]bool)
	//visited[pos.x][pos.y] = true
	for _, c := range strings.Split(commands, ",") {
		move(&pos, strings.TrimSpace(c), visited)
	}
	return abs(pos.x) + abs(pos.y)
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
