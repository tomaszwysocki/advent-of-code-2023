package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var input []string
var symbolRegex = regexp.MustCompile(`[^.0-9]`)

func main() {
	var lineNumber, sum int
	numberRegex := regexp.MustCompile(`\d+`)

	f1, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f1.Close()
	scanner1 := bufio.NewScanner(f1)

	f2, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f2.Close()
	scanner2 := bufio.NewScanner(f2)

	// Scan through the entire input file, capturing each line and storing
	// it in the 'input' slice.
	for scanner1.Scan() {
		line := scanner1.Text()
		input = append(input, line)
	}

	// Iterate through each line, identify numeric sequences, and check
	// if they qualify as part numbers. If they qualify, convert them to
	// an integer and accumulate them to the sum.
	for scanner2.Scan() {
		line := scanner2.Text()
		locs := numberRegex.FindAllStringIndex(line, -1)

		for _, loc := range locs {
			if checkForSymbols(lineNumber, loc) {
				partNumber, _ := strconv.Atoi(line[loc[0]:loc[1]])
				sum += partNumber
			}
		}
		lineNumber++
	}
	fmt.Println(sum)
}

// checkForSymbols takes a line number of a match and a two-element
// slice of integers defining the location of the match and returns
// true if it finds any symbols next to the match and false otherwise.
func checkForSymbols(lineNumber int, loc []int) bool {
	var rangeA, rangeB int
	line := input[lineNumber]

	if loc[0] > 0 && loc[1] < len(line) {
		rangeA, rangeB = loc[0]-1, loc[1]+1
	} else if loc[0] == 0 {
		rangeA, rangeB = loc[0], loc[1]+1
	} else {
		rangeA, rangeB = loc[0]-1, loc[1]
	}

	if lineNumber > 0 {
		prevLine := input[lineNumber-1]
		if symbolRegex.MatchString(prevLine[rangeA:rangeB]) {
			return true
		}
	}
	if lineNumber < len(input)-1 {
		nextLine := input[lineNumber+1]
		if symbolRegex.MatchString(nextLine[rangeA:rangeB]) {
			return true
		}
	}
	if symbolRegex.MatchString(line[rangeA:rangeB]) {
		return true
	}
	return false
}
