package main

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strconv"
	"unicode"
)

func puzzle1() (int, error) {
	file, err := os.Open("inputs/puzzle1.txt")
	if err != nil {
		return 0, errors.New("could not read file input")
	}

	defer file.Close()
	reader := bufio.NewReader(file)

	calValues := []int{}

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

		digits := []string{}

		// Iterate each byte of the line
		for _, b := range line {

			if unicode.IsDigit(rune(b)) {
				digits = append(digits, string(b))
			}
		}

		if len(digits) > 0 {
			first := digits[0]
			last := digits[len(digits) - 1]

			number := first + last
			calValue, err := strconv.Atoi(number)

			if err == nil {
				calValues = append(calValues, calValue)
			}
		}
	}

	calibration := 0
	for _, calValue := range calValues {
		calibration += calValue
	}

	return calibration, nil
}