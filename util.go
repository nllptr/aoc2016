package aoc2016

import (
	"fmt"
	"io/ioutil"
	"os"
)

func GetInput(file string) string {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("Could not read %s. Make sure it exists.", file)
		os.Exit(1)
	}
	return string(b)
}
