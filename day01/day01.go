package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var stringsToNumberValues = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func parseNumbersFromLine(line string) int64 {
	foundNumbers := []string{}

	var wordSoFar = ""
	for c := 0; c < len(line); c++ {
		wordSoFar = wordSoFar + string(line[c])

		// Check if current is number
		parsedNumber, err := strconv.ParseInt(string(line[c]), 10, 64)
		if err == nil {
			fmt.Println("This is a number", parsedNumber)

			foundNumbers = append(foundNumbers, string(line[c]))

			// If we run across a number we know there is not a
			// word number so just reset
			wordSoFar = ""
			continue
		}

		for key, value := range stringsToNumberValues {
			matched, _ := regexp.MatchString(key, wordSoFar)
			if matched {
				fmt.Println("CONTAINS STR ", rune(value))
				foundNumbers = append(foundNumbers, strconv.FormatInt(int64(value), 10))
				wordSoFar = string(wordSoFar[len(wordSoFar)-1]) + ""
			}
		}
	}

	if len(foundNumbers) <= 0 {
		return 0
	}

	combinedNumber := foundNumbers[0] + foundNumbers[len(foundNumbers)-1]
	fmt.Println("Combined number", combinedNumber)
	finalNumber, err := strconv.ParseInt(combinedNumber, 10, 64)

	fmt.Println("Final number", finalNumber)

	if err != nil {
		return 0
	}

	return finalNumber
}

func main() {
	file, err := os.Open("input.txt")
	//file, err := os.Open("testInput.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)
	var totalAmount int64 = 0

	for scanner.Scan() {
		lineStr := scanner.Text()
		fmt.Println(lineStr)
		lineNumber := parseNumbersFromLine(scanner.Text())
		fmt.Println("Line number result is: ", lineNumber)
		if lineNumber > 0 {
			totalAmount = totalAmount + lineNumber
		}

		fmt.Println("\n\n")
	}

	fmt.Println("Total value is: ", totalAmount)
}
