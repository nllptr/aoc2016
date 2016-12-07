package day2

import "strings"

type keypad [][]byte

const ooo = '-'

var keyPad1 = keypad{
	{ooo, ooo, ooo, ooo, ooo},
	{ooo, '1', '2', '3', ooo},
	{ooo, '4', '5', '6', ooo},
	{ooo, '7', '8', '9', ooo},
	{ooo, ooo, ooo, ooo, ooo},
}

var keyPad2 = keypad{
	{ooo, ooo, ooo, ooo, ooo, ooo, ooo},
	{ooo, ooo, ooo, '1', ooo, ooo, ooo},
	{ooo, ooo, '2', '3', '4', ooo, ooo},
	{ooo, '5', '6', '7', '8', '9', ooo},
	{ooo, ooo, 'A', 'B', 'C', ooo, ooo},
	{ooo, ooo, ooo, 'D', ooo, ooo, ooo},
	{ooo, ooo, ooo, ooo, ooo, ooo, ooo},
}

func lookupKey(k keypad, b byte) (x, y int) {
	for y, row := range k {
		for x, cell := range row {
			if cell == b {
				return x, y
			}
		}
	}
	return -1, -1
}

func processKey(k keypad, curX, curY *int, b byte) byte {
	switch b {
	case 'U':
		if k[*curY-1][*curX] != ooo {
			*curY--
		}
	case 'D':
		if k[*curY+1][*curX] != ooo {
			*curY++
		}
	case 'L':
		if k[*curY][*curX-1] != ooo {
			*curX--
		}
	case 'R':
		if k[*curY][*curX+1] != ooo {
			*curX++
		}
	}
	return k[*curY][*curX]
}

func processLine(k keypad, curX, curY *int, line string) byte {
	var out byte
	for i := range line {
		out = processKey(k, curX, curY, line[i])
	}
	return out
}

// FindCodeKeypad1 returns the code for a "regular" keypad
func FindCodeKeypad1(input string) string {
	curX := new(int)
	curY := new(int)
	*curX, *curY = lookupKey(keyPad1, '5')
	var code string
	for _, line := range strings.Split(input, "\n") {
		code += string(processLine(keyPad1, curX, curY, line))
	}
	return code
}

// FindCodeKeypad2 returns the code for an "odd" keypad
func FindCodeKeypad2(input string) string {
	curX := new(int)
	curY := new(int)
	*curX, *curY = lookupKey(keyPad2, '5')
	var code string
	for _, line := range strings.Split(input, "\n") {
		code += string(processLine(keyPad2, curX, curY, line))
	}
	return code
}
