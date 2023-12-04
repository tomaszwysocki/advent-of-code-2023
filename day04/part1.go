package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strings"
)

var numberRegex = regexp.MustCompile(`\d+`)

func main() {
	var sum, lineNumber int

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		winningNumbers, myNumbers := processLine(line)
		sum += checkIfWinning(winningNumbers, myNumbers)
		lineNumber++
	}
	fmt.Println(sum)
}

// checkIfWinning takes as arguments a slice of winning numbers and a slice
// of "my numbers" and returns how many points each card is worth
func checkIfWinning(winning, my []string) int {
	var match, points int
	for _, w := range winning {
		for _, m := range my {
			if w == m {
				match++
			}
		}
	}
	if match > 0 {
		points = int(math.Pow(2, float64(match-1)))
	}
	return points
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
