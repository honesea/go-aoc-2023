package main

import (
	"bufio"
	"errors"
	"io"
	"os"
)

func day14() (int, int, error) {
	answerP1, err := d14p1()
	if err != nil {
		return 0, 0, err
	}

	answerP2, err := d14p2()
	if err != nil {
		return 0, 0, err
	}

	return answerP1, answerP2, err
}

func d14p1() (int, error) {
	file, err := os.Open("inputs/d14p1.txt")
	if err != nil {
		return 0, errors.New("could not read file input")
	}

	defer file.Close()
	reader := bufio.NewReader(file)

	platformStrings := []string{}

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

		platformStrings = append(platformStrings, string(line))
	}

	platform := make([][]byte, len(platformStrings))
	for row := range platformStrings {
		platformRow := make([]byte, len(platformStrings[row]))
		for col := range platformStrings[row] {
			platformRow[col] = platformStrings[row][col]
		}
		platform[row] = platformRow
	}

	tiltedPlatformLines := tiltPlatform(platform)

	total := 0
	for row := range tiltedPlatformLines {
		for col := range tiltedPlatformLines[row] {
			if tiltedPlatformLines[row][col] == 'O' {
				total += len(tiltedPlatformLines) - row
			}
		}
	}

	return total, nil
}

func d14p2() (int, error) {
	file, err := os.Open("inputs/d14p2.txt")
	if err != nil {
		return 0, errors.New("could not read file input")
	}

	defer file.Close()
	reader := bufio.NewReader(file)

	platformStrings := []string{}

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

		platformStrings = append(platformStrings, string(line))
	}

	platform := make([][]byte, len(platformStrings))
	for row := range platformStrings {
		platformRow := make([]byte, len(platformStrings[row]))
		for col := range platformStrings[row] {
			platformRow[col] = platformStrings[row][col]
		}
		platform[row] = platformRow
	}

	platform = spinPlatform(platform, 1000000000)

	total := 0
	for row := range platform {
		for col := range platform[row] {
			if platform[row][col] == 'O' {
				total += len(platform) - row
			}
		}
	}

	return total, nil
}

func tiltPlatform(platform [][]byte) [][]byte {
	tiltedPlatform := make([][]byte, len(platform))
	for row := range tiltedPlatform {
		tiltedPlatform[row] = make([]byte, len(platform[0]))

		for col := range tiltedPlatform[row] {
			tiltedPlatform[row][col] = platform[row][col]
		}
	}

	for col := range tiltedPlatform[0] {
		for row := range tiltedPlatform {
			if tiltedPlatform[row][col] != 'O' {
				continue
			}

			// Backtrack to find available space
			for i := row - 1; i >= 0; i-- {
				if tiltedPlatform[i][col] != '.' {
					break
				}

				tiltedPlatform[i][col] = 'O'
				tiltedPlatform[i+1][col] = '.'
			}
		}
	}

	return tiltedPlatform
}

func spinPlatform(platform [][]byte, spins int) [][]byte {
	platformStates := [][][]byte{}
	cycleFound := false

	for i := 0; i < spins; i++ {
		platformStates = append(platformStates, platform)

		// NORTH
		platform = tiltPlatform(platform)
		// WEST
		platform = rotate90(platform)
		platform = tiltPlatform(platform)
		// SOUTH
		platform = rotate90(platform)
		platform = tiltPlatform(platform)
		// EAST
		platform = rotate90(platform)
		platform = tiltPlatform(platform)
		// NORTH
		platform = rotate90(platform)

		if cycleFound {
			continue
		}

		for j := 0; j < len(platformStates); j++ {
			if platformsMatch(platform, platformStates[j]) {
				spinsPerCycle := i - j + 1
				spinsLeft := spins - i - 1
				spinsLeftAfterCycles := spinsLeft % spinsPerCycle
				i = spins - spinsLeftAfterCycles - 1
				cycleFound = true
				break
			}
		}
	}

	return platform
}

func rotate90(matrix [][]byte) [][]byte {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return matrix
	}

	rowCount := len(matrix)
	colCount := len(matrix[0])
	rotated := make([][]byte, colCount)

	for i := range rotated {
		rotated[i] = make([]byte, rowCount)
		for j := range rotated[i] {
			rotated[i][j] = matrix[rowCount-j-1][i]
		}
	}

	return rotated
}

func platformsMatch(platform1, platform2 [][]byte) bool {
	for row := range platform1 {
		for col := range platform1[row] {
			if platform1[row][col] != platform2[row][col] {
				return false
			}
		}
	}

	return true
}
