package day4

import (
	"strconv"
	"strings"
)

var alphabet = []rune{
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
}

func validate(name string) (bool, string, int) {
	split := strings.Split(name, "[")
	var namePart, sectorID string
	for i := len(split[0]) - 1; i >= 0; i-- {
		if split[0][i] == '-' {
			namePart = split[0][:i]
			sectorID = split[0][i+1:]
			break
		}
	}
	checksum := split[1][:len(split[1])-1]
	var histo = make(map[rune]int)
	for _, c := range namePart {
		if c != '-' {
			histo[c]++
		}
	}

	checkCalc := ""
	for i := 0; i < 5; i++ {
		var max rune
		for _, l := range alphabet {
			if histo[l] > histo[max] {
				max = l
			}
		}
		checkCalc += string(max)
		histo[max] = 0
	}

	sector, _ := strconv.Atoi(sectorID)
	return checkCalc == checksum, namePart, sector
}

func rotX(encrypted string, rot int) string {
	clear := ""
	rot = rot % len(alphabet)
	for _, c := range encrypted {
		if c == '-' {
			clear += string(' ')
		} else {
			d := c + rune(rot)
			if d > 'z' {
				d -= rune(len(alphabet))
			}
			clear += string(d)
		}
	}
	return clear
}

// SumSectorIDs sums up the sector IDs for all the valid rooms
func SumSectorIDs(input string) int {
	lines := strings.Split(input, "\n")
	var sum = 0
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if valid, _, sector := validate(trimmed); valid {
			sum += sector
		}

	}
	return sum
}

// DecryptAndFindNorthPoleObjects finds the sector where North Pole Objects are stored.
func DecryptAndFindNorthPoleObjects(input string) int {
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if valid, encrypted, sector := validate(trimmed); valid {
			if clear := rotX(encrypted, sector); strings.Contains(clear, "north") && strings.Contains(clear, "pole") && strings.Contains(clear, "object") {
				return sector
			}

		}

	}
	return -1
}
