package main

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strings"
)

type MirrorPattern struct {
	pattern [][]bool
}

func day13() (int, int, error) {
	answerP1, err := d13p1()
	if err != nil {
		return 0, 0, err
	}

	answerP2, err := d13p2()
	if err != nil {
		return 0, 0, err
	}

	return answerP1, answerP2, err
}

func d13p1() (int, error) {
	file, err := os.Open("inputs/d13p1.txt")
	if err != nil {
		return 0, errors.New("could not read file input")
	}

	defer file.Close()
	reader := bufio.NewReader(file)

	mirrorPatterns := []MirrorPattern{}
	mirrorPattern := MirrorPattern{}

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

		patternLineStr := strings.Split(string(line), "")
		if len(patternLineStr) == 0 {
			mirrorPatterns = append(mirrorPatterns, mirrorPattern)
			mirrorPattern = MirrorPattern{}
			continue
		}

		patternLine := []bool{}
		for _, character := range patternLineStr {
			isRock := character == "#"
			patternLine = append(patternLine, isRock)
		}

		mirrorPattern.pattern = append(mirrorPattern.pattern, patternLine)
	}

	// Add last mirror pattern
	mirrorPatterns = append(mirrorPatterns, mirrorPattern)

	total := 0
	for _, mirrorPattern = range mirrorPatterns {

		for i := 1; i < len(mirrorPattern.pattern); i++ {
			if checkHorizontalReflection(mirrorPattern.pattern, i-1, i) {
				total += i * 100
				break
			}
		}

		for i := 1; i < len(mirrorPattern.pattern[0]); i++ {
			if checkVerticalReflection(mirrorPattern.pattern, i-1, i) {
				total += i
				break
			}
		}
	}

	return total, nil
}

func d13p2() (int, error) {
	file, err := os.Open("inputs/d13p2.txt")
	if err != nil {
		return 0, errors.New("could not read file input")
	}

	defer file.Close()
	reader := bufio.NewReader(file)

	mirrorPatterns := []MirrorPattern{}
	mirrorPattern := MirrorPattern{}

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

		patternLineStr := strings.Split(string(line), "")
		if len(patternLineStr) == 0 {
			mirrorPatterns = append(mirrorPatterns, mirrorPattern)
			mirrorPattern = MirrorPattern{}
			continue
		}

		patternLine := []bool{}
		for _, character := range patternLineStr {
			isRock := character == "#"
			patternLine = append(patternLine, isRock)
		}

		mirrorPattern.pattern = append(mirrorPattern.pattern, patternLine)
	}

	// Add last mirror pattern
	mirrorPatterns = append(mirrorPatterns, mirrorPattern)

	total := 0
	for _, mirrorPattern = range mirrorPatterns {
		hasHorizontalReflection := false

		for i := 1; i < len(mirrorPattern.pattern); i++ {
			if checkSmudgedHorizontalReflection(mirrorPattern.pattern, i-1, i) {
				total += i * 100
				hasHorizontalReflection = true
				break
			}
		}

		if hasHorizontalReflection {
			continue
		}

		for i := 1; i < len(mirrorPattern.pattern[0]); i++ {
			if checkSmudgedVerticalReflection(mirrorPattern.pattern, i-1, i) {
				total += i
				break
			}
		}
	}

	return total, nil
}

func checkHorizontalReflection(pattern [][]bool, upperRow, lowerRow int) bool {
	if rowMismatchingChars(pattern, upperRow, lowerRow) > 0 {
		return false
	}

	for {
		upperRow = upperRow - 1
		lowerRow = lowerRow + 1

		// if index doesn't exist break and valid
		if upperRow < 0 || lowerRow == len(pattern) {
			break
		}

		// if not match not valid
		if rowMismatchingChars(pattern, upperRow, lowerRow) > 0 {
			return false
		}
	}

	return true
}

func checkVerticalReflection(pattern [][]bool, leftCol, rightCol int) bool {
	if colMismatchedChars(pattern, leftCol, rightCol) > 0 {
		return false
	}

	for {
		leftCol = leftCol - 1
		rightCol = rightCol + 1

		// if index doesn't exist break and valid
		if leftCol < 0 || rightCol == len(pattern[0]) {
			break
		}

		// if not match not valid
		if colMismatchedChars(pattern, leftCol, rightCol) > 0 {
			return false
		}
	}

	return true
}

func checkSmudgedHorizontalReflection(pattern [][]bool, upperRow, lowerRow int) bool {
	initMismatchingChars := rowMismatchingChars(pattern, upperRow, lowerRow)
	if initMismatchingChars > 1 {
		return false
	}

	smudges := initMismatchingChars
	for {
		upperRow = upperRow - 1
		lowerRow = lowerRow + 1

		// if index doesn't exist break and valid
		if upperRow < 0 || lowerRow == len(pattern) {
			break
		}

		mismatchingChars := rowMismatchingChars(pattern, upperRow, lowerRow)
		if mismatchingChars == 1 && smudges == 1 {
			return false
		} else if mismatchingChars == 1 {
			smudges = 1
			continue
		} else if mismatchingChars > 0 {
			return false
		}
	}

	return smudges == 1
}

func checkSmudgedVerticalReflection(pattern [][]bool, leftCol, rightCol int) bool {
	initMismatchingChars := colMismatchedChars(pattern, leftCol, rightCol)
	if initMismatchingChars > 1 {
		return false
	}

	smudges := initMismatchingChars
	for {
		leftCol = leftCol - 1
		rightCol = rightCol + 1

		// if index doesn't exist break and valid
		if leftCol < 0 || rightCol == len(pattern[0]) {
			break
		}

		mismatchingChars := colMismatchedChars(pattern, leftCol, rightCol)
		if mismatchingChars == 1 && smudges == 1 {
			return false
		} else if mismatchingChars == 1 {
			smudges = 1
			continue
		} else if mismatchingChars > 0 {
			return false
		}
	}

	return smudges == 1
}

func rowMismatchingChars(pattern [][]bool, row1, row2 int) int {
	mismatchedChars := 0

	for i := 0; i < len(pattern[0]); i++ {
		row1Char := pattern[row1][i]
		row2Char := pattern[row2][i]

		if row1Char != row2Char {
			mismatchedChars++
		}
	}

	return mismatchedChars
}

func colMismatchedChars(pattern [][]bool, col1, col2 int) int {
	mismatchedChars := 0

	for i := 0; i < len(pattern); i++ {
		col1Char := pattern[i][col1]
		col2Char := pattern[i][col2]

		if col1Char != col2Char {
			mismatchedChars++
		}
	}

	return mismatchedChars
}
