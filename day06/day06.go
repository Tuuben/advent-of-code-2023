package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parse(line string) []int {
	numberStr := strings.Trim(strings.Split(line, ":")[1], " ")
	numbersStrs := strings.Split(numberStr, " ")
	fmt.Println("NUMBER STR:", numbersStrs)

	numbers := []int{}
	for i := 0; i < len(numbersStrs); i++ {
		num, _ := strconv.Atoi(strings.Trim(numbersStrs[i], " "))

		if num == 0 {
			continue
		}

		fmt.Println("NUMBER:", num)
		numbers = append(numbers, num)
	}

	return numbers
}

func main() {
	file, err := os.Open("test-input.txt")
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

	times := []int{}
	distances := []int{}

	for scanner.Scan() {
		lineStr := scanner.Text()
		fmt.Println(lineStr)

		if len(times) == 0 {
			times = parse(lineStr)
		} else {
			distances = parse(lineStr)
		}
	}

	counts := []int{}
	fmt.Println("times", times)
	fmt.Println("distances", distances)

	/*
		Time:      7  15   30
		Distance:  9  40  200

		hold for 2 = 2 miliseconds per 1 second
		5 * 2 = 10 distance

		hold for 3 = 3 miliseconds per 1 second
		4 * 3 = 12 distance
	*/
	for y := 0; y < len(times); y++ {
		time := times[y]
		distance := distances[y]
		count := 0
		for i := 0; i < time; i++ {
			totalDistance := (distance - i) * i
			fmt.Println("total distance: ", totalDistance)

			if totalDistance > distance {
				count++
			}
		}

		counts = append(counts, count)
	}

	fmt.Println("counts: ", counts)

}
