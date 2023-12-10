package main

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strconv"
	"strings"
)

func day9() (int, int, error) {
	answerP1, err := d9p1()
	if err != nil {
		return 0, 0, err
	}

	answerP2, err := d9p2()
	if err != nil {
		return 0, 0, err
	}

	return answerP1, answerP2, err
}

func d9p1() (int, error) {
	file, err := os.Open("inputs/d9p1.txt")
	if err != nil {
		return 0, errors.New("could not read file input")
	}

	defer file.Close()
	reader := bufio.NewReader(file)

	histories := [][]int{}

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

		readingsStr := strings.Split(string(line), " ")
		history := make([]int, len(readingsStr))

		for i := 0; i < len(history); i++ {
			reading, err := strconv.Atoi(readingsStr[i])
			if err == nil {
				history[i] = reading
			}
		}

		histories = append(histories, history)
	}

	sensorPredictions := make([]int, len(histories))
	for i := 0; i < len(sensorPredictions); i++ {
		sensorPredictions[i] = histories[i][len(histories[i])-1] + getNextLayerDiff(histories[i])
	}

	return sum(sensorPredictions), nil
}

func d9p2() (int, error) {
	file, err := os.Open("inputs/d9p2.txt")
	if err != nil {
		return 0, errors.New("could not read file input")
	}

	defer file.Close()
	reader := bufio.NewReader(file)

	histories := [][]int{}

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

		readingsStr := strings.Split(string(line), " ")
		history := make([]int, len(readingsStr))

		for i := 0; i < len(history); i++ {
			reading, err := strconv.Atoi(readingsStr[i])
			if err == nil {
				history[i] = reading
			}
		}

		histories = append(histories, history)
	}

	sensorPredictions := make([]int, len(histories))
	for i := 0; i < len(sensorPredictions); i++ {
		sensorPredictions[i] = histories[i][0] - getPreviousLayerDiff(histories[i])
	}

	return sum(sensorPredictions), nil
}

func getNextLayerDiff(sequence []int) int {
	if len(sequence) <= 1 {
		return 0
	}

	diffSequence := make([]int, len(sequence)-1)
	for i := 1; i < len(sequence); i++ {
		diffSequence[i-1] = sequence[i] - sequence[i-1]
	}

	for _, diff := range diffSequence {
		if diff != 0 {
			nextDiff := getNextLayerDiff(diffSequence)
			return diffSequence[len(diffSequence)-1] + nextDiff
		}
	}

	return 0
}

func getPreviousLayerDiff(sequence []int) int {
	if len(sequence) <= 1 {
		return 0
	}

	diffSequence := make([]int, len(sequence)-1)
	for i := 1; i < len(sequence); i++ {
		diffSequence[i-1] = sequence[i] - sequence[i-1]
	}

	for _, diff := range diffSequence {
		if diff != 0 {
			nextDiff := getPreviousLayerDiff(diffSequence)
			return diffSequence[0] - nextDiff
		}
	}

	return 0
}
