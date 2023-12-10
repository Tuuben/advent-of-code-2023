package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func extractNumberFromString(input string) (int, error) {
	// Define a regular expression pattern to match numbers
	re := regexp.MustCompile(`(\d+)`)

	// Find the first match in the input string
	matches := re.FindStringSubmatch(input)
	if len(matches) < 2 {
		return 0, fmt.Errorf("no number found in the string")
	}

	// Extract the matched number
	numberStr := matches[1]

	// Convert the string representation of the number to an actual number
	number, err := strconv.Atoi(numberStr)
	if err != nil {
		return 0, err
	}

	return number, nil
}

func parseRow(row string) (map[int]bool, []int, int) {
	entireRow := strings.Split(strings.Trim(row, " "), ":")
	gameNumbers := entireRow[1]
	//gameNumberStr := strings.Split(entireRow[0], " ")[1]
	gameNum, _ := extractNumberFromString(entireRow[0])

	fmt.Println("Game numbers", gameNumbers)
	numbers := strings.Split(strings.Trim(gameNumbers, " "), "|")

	winningNumbers := strings.Trim(numbers[0], " ")
	cardNumbers := strings.Trim(numbers[1], " ")

	cardNumbersSplit := strings.Split(cardNumbers, " ")
	cardNumbersArr := []int{}
	for cardNumI := 0; cardNumI < len(cardNumbersSplit); cardNumI++ {
		num, err := strconv.Atoi(string(cardNumbersSplit[cardNumI]))
		if err != nil {
			continue
		}

		cardNumbersArr = append(cardNumbersArr, num)
	}

	winningNumbersSplit := strings.Split(winningNumbers, " ")
	winningMap := map[int]bool{}
	for winI := 0; winI < len(winningNumbersSplit); winI++ {
		num, err := strconv.Atoi(winningNumbersSplit[winI])
		if err != nil {
			continue
		}
		winningMap[num] = true
	}

	fmt.Println("Winning numbers", winningMap)
	fmt.Println("Card numbers", cardNumbersArr)

	return winningMap, cardNumbersArr, gameNum
}

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)

	totalNum := 0
	aggregatedCards := map[int]int{}
	for scanner.Scan() {
		entireLine := scanner.Text()
		winMap, cardNumbers, cardNumber := parseRow(entireLine)

		cardI := cardNumber
		// we add the og card
		aggregatedCards[cardI] = aggregatedCards[cardI] + 1
		repeatCount := aggregatedCards[cardI]

		for a := 0; a < repeatCount; a++ {
			points := 0
			for i := 0; i < len(cardNumbers); i++ {
				num := cardNumbers[i]
				if winMap[num] {
					points += 1
					nextI := cardI + points
					aggregatedCards[nextI] = aggregatedCards[nextI] + 1
				}
			}
		}
	}

	fmt.Println("Aggregated cards", aggregatedCards)
	for _, v := range aggregatedCards {
		totalNum = totalNum + v
	}
	fmt.Println("Total num: ", totalNum)
}
