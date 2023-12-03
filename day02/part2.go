package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
)

const (
	redCubes   = 12
	greenCubes = 13
	blueCubes  = 14
)

func main() {
	greenRegex := regexp.MustCompile(`(\d+) green`)
	redRegex := regexp.MustCompile(`(\d+) red`)
	blueRegex := regexp.MustCompile(`(\d+) blue`)

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var sum int
	for gameID := 1; scanner.Scan(); gameID++ {
		var power int
		line := scanner.Text()

		greenNumbers := findNumbersRegex(greenRegex, line)
		redNumbers := findNumbersRegex(redRegex, line)
		blueNumbers := findNumbersRegex(blueRegex, line)

		maxGreen := slices.Max(greenNumbers)
		maxRed := slices.Max(redNumbers)
		maxBlue := slices.Max(blueNumbers)

		power = maxGreen * maxRed * maxBlue
		sum += power
	}
	fmt.Println(sum)
}

// findNumbersRegex takes a regex pattern and a line of text
// (e.g., "Game1: 3 blue...") and returns a slice of all numbers
// associated with the color that regex matches
func findNumbersRegex(regex *regexp.Regexp, line string) []int {
	var numbers []int
	matches := regex.FindAllStringSubmatch(line, -1)
	for _, match := range matches {
		matchInt, _ := strconv.Atoi(match[1])
		numbers = append(numbers, matchInt)
	}
	return numbers
}
