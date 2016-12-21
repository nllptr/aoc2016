package day6

import (
	"strings"
)

type posMap map[rune]int

// Correct corrects the repeated code
func Correct(input string) string {
	lines := strings.Split(input, "\n")
	occurrances := make([]posMap, len(lines[0]))
	for i := range occurrances {
		occurrances[i] = make(posMap)
	}
	for _, line := range lines {
		line = strings.TrimSpace(line)
		for i, char := range line {
			occurrances[i][char]++
		}
	}
	output := ""
	for _, aMap := range occurrances {
		max := ' '
		for k, v := range aMap {
			if v > aMap[max] {
				max = k
			}
		}
		output += string(max)
	}
	return output
}

// LeastCorrect corrects the repeated code, using the least common character
func LeastCorrect(input string) string {
	lines := strings.Split(input, "\n")
	occurrances := make([]posMap, len(lines[0]))
	for i := range occurrances {
		occurrances[i] = make(posMap)
	}
	for _, line := range lines {
		line = strings.TrimSpace(line)
		for i, char := range line {
			occurrances[i][char]++
		}
	}
	output := ""
	for _, aMap := range occurrances {
		min := ' '
		for k, v := range aMap {
			if min == ' ' || v < aMap[min] {
				min = k
			}
		}
		output += string(min)
	}
	return output
}
