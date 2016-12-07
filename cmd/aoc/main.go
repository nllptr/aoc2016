package main

import (
	"fmt"

	"github.com/nllptr/aoc2016"
	"github.com/nllptr/aoc2016/day1"
	"github.com/nllptr/aoc2016/day3"
)

func main() {
	commands := aoc2016.GetInput("input/day1.txt")
	fmt.Println("\n--- Day 1 ---")
	fmt.Println("Number of blocks:", day1.Distance(commands))
	fmt.Println("Number of blocks to first revisited location:", day1.FirstRevisitedDistance(commands))

	triangles := aoc2016.GetInput("input/day3.txt")
	fmt.Println("\n--- Day 3 ---")
	numberValid, _ := day3.NumberOfValid(triangles)
	fmt.Println("Number of valid triangles:", numberValid)
	numberValidVertical, _ := day3.NumberOfValidVertical(triangles)
}
