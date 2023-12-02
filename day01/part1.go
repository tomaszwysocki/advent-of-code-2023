package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {
	sum := 0

	inputFile, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		var digits []int
		for _, c := range scanner.Text() {
			if unicode.IsDigit(c) {
				digit, err := strconv.Atoi(string(c))
				if err != nil {
					log.Fatal(err)
				}
				digits = append(digits, digit)
			}
		}
		firstDigit, lastDigit := digits[0], digits[len(digits)-1]
		calibrationValueStr := fmt.Sprintf("%v%v", firstDigit, lastDigit)
		calibrationValue, err := strconv.Atoi(calibrationValueStr)
		if err != nil {
			log.Fatal(err)
		}
		sum += calibrationValue
	}
	fmt.Println(sum)
}
