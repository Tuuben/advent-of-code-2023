package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var DELETED = -3
var EMPTY = -2
var SYMBOL = -1

type cursorPos struct {
	x int
	y int
}

func isSymbol(char string) bool {
	symbolPattern := regexp.MustCompile("[[:punct:]]")
	return symbolPattern.MatchString(char)
}

func isNumber(char string) bool {
	numberPattern := regexp.MustCompile(`^[0-9]+$`)
	return numberPattern.MatchString(char)
}

func getValueRow(line string) []int {
	var row []int
	for i := 0; i < len(line); i++ {
		char := string(line[i])

		if char == "." {
			row = append(row, -2)
			continue
		}

		if isSymbol(char) {
			row = append(row, -1)
			continue
		}

		if isNumber(char) {
			num, err := strconv.Atoi(char)

			if err != nil {
				continue
			}

			row = append(row, num)
			continue
		}

		row = append(row, 0)
	}

	return row
}

func getCursorPositions(x int, y int, rangeX [2]int, rangeY [2]int) []cursorPos {
	positions := []cursorPos{}
	for cursorY := y - 1; cursorY <= y+1; cursorY++ {
		for cursorX := x - 1; cursorX <= x+1; cursorX++ {
			if cursorY == y && cursorX == x {
				continue
			}

			// Outside range, dont look
			if cursorX < rangeX[0] || cursorX >= rangeX[1] ||
				cursorY < rangeY[0] || cursorY >= rangeY[1] {
				continue
			}

			positions = append(positions, cursorPos{
				x: cursorX,
				y: cursorY,
			})
		}
	}

	return positions
}

func getNumberValueFromRow(valueRow []int, atIndex int) (int, []int) {
	startIndex := atIndex
	numberStr := ""
	collected := []int{}

	for backI := atIndex; backI >= 0; backI-- {
		if valueRow[backI] < 0 {
			break
		}

		startIndex = backI
	}

	for forwardI := startIndex; forwardI < len(valueRow); forwardI++ {
		if valueRow[forwardI] < 0 {
			break
		}

		numberStr += strconv.Itoa(valueRow[forwardI])
		collected = append(collected, forwardI)
	}

	val, err := strconv.Atoi(numberStr)

	if err != nil {
		return 0, []int{}
	}

	return val, collected
}

func parseCollection(valueCol [][]int) int {
	totalNum := 0
	for y := 0; y < len(valueCol); y++ {
		for x := 0; x < len(valueCol[y]); x++ {
			if valueCol[y][x] == SYMBOL {
				cursorPositions := getCursorPositions(x, y, [2]int{0, len(valueCol[y])}, [2]int{0, len(valueCol)})

				for cursorI := 0; cursorI < len(cursorPositions); cursorI++ {
					pos := cursorPositions[cursorI]

					// TODO: This need to somehow get if this position is already checked
					if valueCol[pos.y][pos.x] >= 0 {
						num, collected := getNumberValueFromRow(valueCol[pos.y], pos.x)
						totalNum += num

						for collI := 0; collI < len(collected); collI++ {
							removeI := collected[collI]
							valueCol[pos.y][removeI] = DELETED
						}
					}
				}
			}
		}
	}

	return totalNum
}

func main() {
	valueCollection := [][]int{}
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

	for scanner.Scan() {
		entireLine := scanner.Text()
		rowValues := getValueRow(entireLine)
		valueCollection = append(valueCollection, rowValues)
	}

	fmt.Println("Value col", valueCollection)
	totalSum := parseCollection(valueCollection)
	fmt.Println("\n\nTotal sum: ", totalSum)
}
