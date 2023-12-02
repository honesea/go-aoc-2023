package main

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var stringToNumber = map[string]string{
	"one": "1",
	"two": "2",
	"three": "3",
	"four": "4",
	"five": "5",
	"six": "6",
	"seven": "7",
	"eight": "8",
	"nine": "9",
}

func puzzle2() (int, error) {
	file, err := os.Open("inputs/puzzle2.txt")
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
		availableCharacters := ""

		// Iterate each byte of the line
		for _, b := range line {
			if unicode.IsDigit(rune(b)) {
				digits = append(digits, string(b))
				availableCharacters = ""
			}

			availableCharacters += string(b)
			for number, val := range stringToNumber {
				if strings.Contains(availableCharacters, number) {
					digits = append(digits, val)
					availableCharacters = string(number[len(number)-1]) // handle case where end of one number can be start of another
					break
				}
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