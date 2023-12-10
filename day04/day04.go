package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseRow(row string) (map[int]bool, []int) {
	gameNumbers := strings.Split(strings.Trim(row, " "), ":")[1]
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

	return winningMap, cardNumbersArr
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
	for scanner.Scan() {
		entireLine := scanner.Text()
		winMap, cardNumbers := parseRow(entireLine)
		//fmt.Println("Entire line", entireLine)

		points := 0
		for i := 0; i < len(cardNumbers); i++ {
			num := cardNumbers[i]
			if winMap[num] {
				if points == 0 {
					points = 1
				} else {
					points = points * 2
				}

				fmt.Println("This is a win number!", num, " cur p: ", points)
			}
		}

		totalNum += points
	}

	fmt.Println("Total num: ", totalNum)
}
