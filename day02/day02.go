package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var gameTestData = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

var gameMinDices = map[string]int{}

func parseGameNumber(str string) int64 {
	split := strings.Split(str, ":")[0]

	fmt.Println("Split game number str", split)
	gameNr := strings.Split(strings.Trim(split, " "), " ")

	fmt.Println("Split is", gameNr[1])

	gameNrParsed, err := strconv.ParseInt(gameNr[1], 10, 64)

	if err == nil {
		return gameNrParsed
	}

	return 0
}

func checkIsSetValid(str string) bool {
	dices := strings.SplitAfter(str, ",")

	for i := 0; i < len(dices); i++ {
		diceStr := strings.TrimRight(strings.TrimSpace(dices[i]), ",|;")
		split := strings.SplitAfter(diceStr, " ")

		var value = string(strings.Trim(split[0], " "))
		var key = string(split[1])

		parsedVal, _ := strconv.Atoi(value)

		if gameTestData[key] < int(parsedVal) {
			fmt.Println("This is not a valid value", parsedVal)
			return false
		}
	}

	return true
}

func parseSets(str string) bool {
	allSets := strings.SplitAfter(str, ":")[1]
	sets := strings.Split(allSets, ";")

	for i := 0; i < len(sets); i++ {
		isValid := checkIsSetValid(sets[i])

		if isValid == false {
			return false
		}
	}

	return true
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

	gameTotal := 0

	for scanner.Scan() {
		lineStr := scanner.Text()
		fmt.Println(lineStr)

		gameNr := parseGameNumber(lineStr)
		fmt.Println("Game nr", gameNr)
		isValid := parseSets(lineStr)

		if isValid {
			gameTotal = gameTotal + int(gameNr)
		}

		fmt.Println("\n")
	}

	fmt.Println("Game total: ", gameTotal)
}
