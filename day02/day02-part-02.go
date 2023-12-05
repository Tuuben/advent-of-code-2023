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

func parseOutSets(str string) map[string]int {
	trimmed := strings.Trim(str, " ")
	allSets := strings.Split(trimmed, ":")[1]
	sets := strings.Split(allSets, ";")

	var mappedSet = map[string]int{}

	for setI := 0; setI < len(sets); setI++ {
		set := sets[setI]
		fmt.Println("Set :", set)
		cubes := strings.Split(set, ",")

		for cubeI := 0; cubeI < len(cubes); cubeI++ {
			cube := strings.Trim(cubes[cubeI], " ")
			keyValues := strings.Split(cube, " ")
			key := keyValues[1]
			value, err := strconv.Atoi(keyValues[0])

			if err != nil {
				continue
			}

			// if current is larger
			if mappedSet[key] > value {
				continue
			}

			mappedSet[key] = value
		}
	}

	return mappedSet
}

func setToPower(set map[string]int) int {
	var totalPower int = 1
	for _, v := range set {
		totalPower = totalPower * v
	}

	return totalPower
}

// part 2 because p1 code is complete shit
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

	totalSum := 0

	for scanner.Scan() {
		entireLine := scanner.Text()
		fmt.Println("Entire line: ", entireLine)

		parsedSet := parseOutSets(entireLine)
		powerOfSet := setToPower(parsedSet)

		fmt.Println("Set: ", parsedSet)
		fmt.Println("Total power of set: ", powerOfSet)

		totalSum += powerOfSet
	}

	fmt.Println("\n\nTotal sum: ", totalSum)
}
