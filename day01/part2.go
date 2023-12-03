package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var digitsAsWords = []string{
	"one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
}

func main() {
	sum := 0
	inputFile, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		firstDigit, err := findDigit(line)
		if err != nil {
			log.Fatal(err)
		}
		lastDigit, err := findDigitRight(line)
		if err != nil {
			log.Fatal(err)
		}
		calibrationValueStr := fmt.Sprintf("%v%v", firstDigit, lastDigit)
		calibrationValue, err := strconv.Atoi(calibrationValueStr)
		if err != nil {
			log.Fatal(err)
		}
		sum += calibrationValue
	}
	fmt.Println(sum)
}

func reverseString(input string) string {
	inputRunes := []rune(input)
	length := len(inputRunes)
	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		inputRunes[i], inputRunes[j] = inputRunes[j], inputRunes[i]
	}
	reversedString := string(inputRunes)
	return reversedString
}

// findDigit returns a first digit that it finds or an
// error if it doesn't find a digit
func findDigit(line string) (digit int, err error) {
	end := 1
	for _, c := range line {
		if unicode.IsDigit(c) {
			digit, err := strconv.Atoi(string(c))
			if err != nil {
				log.Fatal(err)
			}
			return digit, nil
		} else {
			for i, d := range digitsAsWords {
				if strings.Contains(line[:end], d) {
					return i + 1, nil
				}
			}
		}
		end++
	}
	return 0, fmt.Errorf("findDigit %q: no digit found", line)
}

// findDigitRight returns a first digit it finds starting
// from the right side or an error if it doesn't find a digit
func findDigitRight(line string) (digit int, err error) {
	start := len(line) - 2
	lineReversed := reverseString(line)
	for idx, _ := range line {
		c := rune(lineReversed[idx])
		if unicode.IsDigit(c) {
			digit, err := strconv.Atoi(string(c))
			if err != nil {
				log.Fatal(err)
			}
			return digit, nil
		} else {
			for i, d := range digitsAsWords {
				if strings.Contains(line[start:], d) {
					return i + 1, nil
				}
			}
		}
		start--
	}
	return 0, fmt.Errorf("findDigitRight %q: no digit found", line)
}
