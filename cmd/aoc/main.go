package main

import (
	"fmt"

	"github.com/nllptr/aoc2016"
	"github.com/nllptr/aoc2016/day1"
	"github.com/nllptr/aoc2016/day2"
	"github.com/nllptr/aoc2016/day3"
	"github.com/nllptr/aoc2016/day4"
	"github.com/nllptr/aoc2016/day5"
)

func main() {
	commands := aoc2016.GetInput("input/day1.txt")
	fmt.Println("\n--- Day 1 ---")
	fmt.Println("Number of blocks:", day1.Distance(commands))
	fmt.Println("Number of blocks to first revisited location:", day1.FirstRevisitedDistance(commands))

	keypadLines := aoc2016.GetInput("input/day2.txt")
	fmt.Println("\n--- Day 2 ---")
	fmt.Println("Code:", day2.FindCodeKeypad1(keypadLines))
	fmt.Println("Code:", day2.FindCodeKeypad2(keypadLines))

	triangles := aoc2016.GetInput("input/day3.txt")
	fmt.Println("\n--- Day 3 ---")
	numberValid, _ := day3.NumberOfValid(triangles)
	fmt.Println("Number of valid triangles:", numberValid)
	numberValidVertical, _ := day3.NumberOfValidVertical(triangles)
	fmt.Println("Number of valid triangles (input read vertically):", numberValidVertical)

	rooms := aoc2016.GetInput("input/day4.txt")
	fmt.Println("\n--- Day 4 ---")
	fmt.Println("Sum of SecotrIDs of valid rooms:", day4.SumSectorIDs(rooms))
	fmt.Println("Sum of SecotrIDs of valid rooms:", day4.DecryptAndFindNorthPoleObjects(rooms))

	doorID := aoc2016.GetInput("input/day5.txt")
	fmt.Println("\n--- Day 5 ---")
	day5.CrackPassword(doorID)
	fmt.Println(" Password 1 cracked!")
	day5.CrackPassword2(doorID)
	fmt.Println(" Password 2 cracked!")
}
