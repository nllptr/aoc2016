package main

import (
	"fmt"

	"flag"

	"github.com/nllptr/aoc2016"
	"github.com/nllptr/aoc2016/day1"
	"github.com/nllptr/aoc2016/day2"
	"github.com/nllptr/aoc2016/day3"
	"github.com/nllptr/aoc2016/day4"
	"github.com/nllptr/aoc2016/day5"
	"github.com/nllptr/aoc2016/day6"
)

func main() {
	sPtr := flag.Int("s", 0, "The solution to run. Defaults to running all solutions.")
	flag.Parse()

	if *sPtr == 0 || *sPtr == 1 {
		commands := aoc2016.GetInput("input/day1.txt")
		fmt.Println("\n--- Day 1 ---")
		fmt.Println("Number of blocks:", day1.Distance(commands))
		fmt.Println("Number of blocks to first revisited location:", day1.FirstRevisitedDistance(commands))
	}

	if *sPtr == 0 || *sPtr == 2 {
		keypadLines := aoc2016.GetInput("input/day2.txt")
		fmt.Println("\n--- Day 2 ---")
		fmt.Println("Code:", day2.FindCodeKeypad1(keypadLines))
		fmt.Println("Code:", day2.FindCodeKeypad2(keypadLines))
	}

	if *sPtr == 0 || *sPtr == 3 {
		triangles := aoc2016.GetInput("input/day3.txt")
		fmt.Println("\n--- Day 3 ---")
		numberValid, _ := day3.NumberOfValid(triangles)
		fmt.Println("Number of valid triangles:", numberValid)
		numberValidVertical, _ := day3.NumberOfValidVertical(triangles)
		fmt.Println("Number of valid triangles (input read vertically):", numberValidVertical)
	}

	if *sPtr == 0 || *sPtr == 4 {
		rooms := aoc2016.GetInput("input/day4.txt")
		fmt.Println("\n--- Day 4 ---")
		fmt.Println("Sum of SecotrIDs of valid rooms:", day4.SumSectorIDs(rooms))
		fmt.Println("Sum of SecotrIDs of valid rooms:", day4.DecryptAndFindNorthPoleObjects(rooms))
	}

	if *sPtr == 0 || *sPtr == 5 {
		doorID := aoc2016.GetInput("input/day5.txt")
		fmt.Println("\n--- Day 5 ---")
		day5.CrackPassword(doorID)
		fmt.Println(" Password 1 cracked!")
		day5.CrackPassword2(doorID)
		fmt.Println(" Password 2 cracked!")
	}

	if *sPtr == 0 || *sPtr == 6 {
		errorLines := aoc2016.GetInput("input/day6.txt")
		fmt.Println("\n--- Day 6 ---")
		fmt.Println("Error corrected message:", day6.Correct(errorLines))
		fmt.Println("Error corrected message:", day6.LeastCorrect(errorLines))
	}
}
