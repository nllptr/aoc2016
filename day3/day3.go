package day3

import (
	"fmt"
	"strconv"
	"strings"
)

func isValid(a, b, c int) bool {
	return a+b > c && a+c > b && b+c > a
}

// NumberOfValid caluculates the number of valid triangles found in the input string.
func NumberOfValid(input string) (int, error) {
	numValid := 0
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		sides := strings.Fields(line)
		if len(sides) != 3 {
			return -1, fmt.Errorf("Badly formatted input: %v", line)
		}
		a, err := strconv.Atoi(sides[0])
		b, err := strconv.Atoi(sides[1])
		c, err := strconv.Atoi(sides[2])
		if err != nil {
			return -1, fmt.Errorf("Must be numerical: %v", line)
		}
		if isValid(a, b, c) {
			numValid++
		}
	}
	return numValid, nil
}
