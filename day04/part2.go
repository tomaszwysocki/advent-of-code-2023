package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

var numberRegex = regexp.MustCompile(`\d+`)

func main() {
	copies := make(map[int]int)
	var sum, lineNumber, numberOfLines int

	f1, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f1.Close()

	f2, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f2.Close()

	scanner1 := bufio.NewScanner(f1)
	scanner2 := bufio.NewScanner(f2)

	for scanner1.Scan() {
		numberOfLines++
	}

	for i := 0; i < numberOfLines; i++ {
		copies[i] = 1
	}

	for scanner2.Scan() {
		line := scanner2.Text()
		winningNumbers, myNumbers := processLine(line)
		matches := checkMatches(winningNumbers, myNumbers)
		for i := 1; i <= matches; i++ {
			copies[lineNumber+i] += copies[lineNumber]
		}
		sum += copies[lineNumber]
		lineNumber++
	}
	fmt.Println(sum)
}

// checkIfWinning takes as arguments a slice of winning numbers and a slice
// of "my numbers" and returns how many points each card is worth
func checkMatches(winning, my []string) int {
	var matches int
	for _, w := range winning {
		for _, m := range my {
			if w == m {
				matches++
			}
		}
	}
	return matches
}

// processLine takes as an argument a line of input and returns a slice
// of winningNumbers and a slice of "your numbers"
func processLine(line string) ([]string, []string) {
	line = strings.Split(line, ":")[1]
	winningNumbersStr := strings.Split(line, "|")[0]
	myNumbersStr := strings.Split(line, "|")[1]
	winningNumbers := numberRegex.FindAllString(winningNumbersStr, -1)
	myNumbers := numberRegex.FindAllString(myNumbersStr, -1)
	return winningNumbers, myNumbers
}
