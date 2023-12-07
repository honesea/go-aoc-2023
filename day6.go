package main

import (
	"bufio"
	"errors"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func day6() (int, int, error) {
	answerP1, err := d6p1()
	if err != nil {
		return 0, 0, err
	}

	answerP2, err := d6p2()
	if err != nil {
		return 0, 0, err
	}

	return answerP1, answerP2, err
}

func d6p1() (int, error) {
	file, err := os.Open("inputs/d6p1.txt")
	if err != nil {
		return 0, errors.New("could not read file input")
	}

	defer file.Close()
	reader := bufio.NewReader(file)

	raceTimes := []int{}
	raceDistances := []int{}

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

		if len(raceTimes) == 0 {
			raceTimes = parseRaceLine(line)
		} else {
			raceDistances = parseRaceLine(line)
		}
	}

	totalWins := 1 // Result needs to be multiplied so start at 1
	for i := 0; i < len(raceTimes); i++ {
		start, end := calcWinningTimes(raceTimes[i], raceDistances[i])
		totalWins *= end - start - 1 // Minus one to account for exclusion of first number
	}

	return totalWins, nil
}

func d6p2() (int, error) {
	file, err := os.Open("inputs/d6p2.txt")
	if err != nil {
		return 0, errors.New("could not read file input")
	}

	defer file.Close()
	reader := bufio.NewReader(file)

	raceTime := 0
	raceDistance := 0

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

		if raceTime == 0 {
			raceTime = parseBigRaceLine(line)
		} else {
			raceDistance = parseBigRaceLine(line)
		}
	}

	start, end := calcWinningTimes(raceTime, raceDistance)
	totalWins := end - start - 1 // Minus one to account for exclusion of first number

	return totalWins, nil
}

func parseRaceLine(line []byte) []int {
	raceStr := strings.Split(string(line), ":")
	if len(raceStr) != 2 {
		return []int{}
	}

	values := []int{}
	for _, valStr := range strings.Split(raceStr[1], " ") {
		val, err := strconv.Atoi(valStr)
		if err == nil {
			values = append(values, val)
		}
	}

	return values
}

func parseBigRaceLine(line []byte) int {
	raceStr := strings.Split(string(line), ":")
	if len(raceStr) != 2 {
		return 0
	}

	valuesStr := []string{}
	for _, valStr := range strings.Split(raceStr[1], " ") {
		_, err := strconv.Atoi(valStr)
		if err == nil {
			valuesStr = append(valuesStr, valStr)
		}
	}

	bigValueStr := strings.Join(valuesStr, "")
	val, err := strconv.Atoi(bigValueStr)
	if err != nil {
		return 0
	}

	return val
}

func calcWinningTimes(time int, distance int) (int, int) {
	discriminant := math.Pow(float64(time), 2) - float64(4*distance)
	if discriminant <= 0 {
		return 0, 0
	}

	sqrtDesc := math.Sqrt(discriminant)
	val1 := int(math.Floor((float64(-time) + sqrtDesc) / -2))
	val2 := int(math.Ceil((float64(-time) - sqrtDesc) / -2))

	return val1, val2
}
