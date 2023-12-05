package main

import (
	"bufio"
	"errors"
	"io"
	"math"
	"os"
	"strconv"
	"unicode"
)

type PartNumber struct {
	value  int
	coords []struct {
		x int
		y int
	}
}

func day3() (int, int, error) {
	answerP1, err := d3p1()
	if err != nil {
		return 0, 0, err
	}

	answerP2, err := d3p2()
	if err != nil {
		return 0, 0, err
	}

	return answerP1, answerP2, err
}

func d3p1() (int, error) {
	file, err := os.Open("inputs/d3p1.txt")
	if err != nil {
		return 0, errors.New("could not read file input")
	}

	defer file.Close()
	reader := bufio.NewReader(file)

	partMatrix := [][]byte{}

	for {
		line, isPrefix, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}

			return 0, errors.New("there was an issue reading the file")
		}

		if isPrefix {
			return 0, errors.New("line too long to parse")
		}

		lineCopy := make([]byte, len(line))
		copy(lineCopy, line)
		partMatrix = append(partMatrix, lineCopy)
	}

	partNumbers := []int{}
	for y, row := range partMatrix {
		currentNumber := ""
		isPartNumber := false

		for x, character := range row {
			if unicode.IsDigit(rune(character)) {
				currentNumber = currentNumber + string(character)
				if !isPartNumber {
					isPartNumber = checkSpecialCharacter(partMatrix, x, y)
				}
			} else {
				partNumber, err := strconv.Atoi(currentNumber)
				if err == nil && isPartNumber {
					partNumbers = append(partNumbers, partNumber)
				}

				currentNumber = ""
				isPartNumber = false
			}
		}

		if currentNumber != "" && isPartNumber {
			partNumber, err := strconv.Atoi(currentNumber)
			if err == nil {
				partNumbers = append(partNumbers, partNumber)
			}
		}
	}

	return sum(partNumbers), nil
}

func d3p2() (int, error) {
	file, err := os.Open("inputs/d3p2.txt")
	if err != nil {
		return 0, errors.New("could not read file input")
	}

	defer file.Close()
	reader := bufio.NewReader(file)

	partMatrix := [][]byte{}

	for {
		line, isPrefix, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}

			return 0, errors.New("there was an issue reading the file")
		}

		if isPrefix {
			return 0, errors.New("line too long to parse")
		}

		lineCopy := make([]byte, len(line))
		copy(lineCopy, line)
		partMatrix = append(partMatrix, lineCopy)
	}

	partNumbers := []PartNumber{}
	for y, row := range partMatrix {
		currentNumber := ""
		coords := []struct {
			x int
			y int
		}{}

		for x, character := range row {
			if unicode.IsDigit(rune(character)) {
				currentNumber = currentNumber + string(character)
				coords = append(coords, struct {
					x int
					y int
				}{x: x, y: y})
			} else {
				partNumberValue, err := strconv.Atoi(currentNumber)
				if err == nil {
					partNumber := PartNumber{
						value:  partNumberValue,
						coords: coords,
					}
					partNumbers = append(partNumbers, partNumber)
				}

				currentNumber = ""
				coords = []struct {
					x int
					y int
				}{}
			}
		}

		if currentNumber != "" {
			partNumberValue, err := strconv.Atoi(currentNumber)
			if err == nil {
				partNumber := PartNumber{
					value:  partNumberValue,
					coords: coords,
				}
				partNumbers = append(partNumbers, partNumber)
			}
		}
	}

	gearRatios := []int{}
	for y, row := range partMatrix {
		for x, character := range row {
			if string(character) == "*" {
				gearRatios = append(gearRatios, calculateGearRatio(partNumbers, x, y))
			}
		}
	}

	return sum(gearRatios), nil
}

func checkSpecialCharacter(partMatrix [][]byte, x int, y int) bool {
	for row := y - 1; row < y+2; row++ {
		if row < 0 || row >= len(partMatrix) {
			continue
		}

		for col := x - 1; col < x+2; col++ {
			if col < 0 || col >= len(partMatrix[row]) {
				continue
			}

			character := partMatrix[row][col]
			if !unicode.IsDigit(rune(character)) && string(character) != "." {
				return true
			}
		}
	}

	return false
}

func calculateGearRatio(partNumbers []PartNumber, x int, y int) int {
	partNumbersInRange := []int{}
	for _, partNumber := range partNumbers {
		for _, coord := range partNumber.coords {
			xDiff := float64(x - coord.x)
			yDiff := float64(y - coord.y)

			if math.Abs(xDiff) <= 1 && math.Abs(yDiff) <= 1 {
				partNumbersInRange = append(partNumbersInRange, partNumber.value)
				break
			}
		}
	}

	if len(partNumbersInRange) == 2 {
		return partNumbersInRange[0] * partNumbersInRange[1]
	}

	return 0
}
