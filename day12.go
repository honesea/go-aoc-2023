package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type SpringLine struct {
	springs   []string
	groupings []int
}

func day12() (int, int, error) {
	answerP1, err := d12p1()
	if err != nil {
		return 0, 0, err
	}

	answerP2, err := d12p2()
	if err != nil {
		return 0, 0, err
	}

	return answerP1, answerP2, err
}

func d12p1() (int, error) {
	file, err := os.Open("inputs/d12p1.txt")
	if err != nil {
		return 0, errors.New("could not read file input")
	}

	defer file.Close()
	reader := bufio.NewReader(file)

	springLines := []SpringLine{}

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

		springLines = append(springLines, parseSpringLine(line))
	}

	permutations := []int{}
	// for _, springLine := range springLines {
	// 	permutations = append(permutations, calculateSpringPermutations(springLine))
	// }

	return sum(permutations), nil
}

func d12p2() (int, error) {
	file, err := os.Open("inputs/d12p2.txt")
	if err != nil {
		return 0, errors.New("could not read file input")
	}

	defer file.Close()
	reader := bufio.NewReader(file)

	springLines := []SpringLine{}

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

		springLines = append(springLines, parseSpringLineLong(line))
	}

	permutations := []int{}
	for _, springLine := range springLines {
		// fmt.Println(springLine)
		permutations = append(permutations, calculateSpringPermutationsFast(springLine))
	}

	return sum(permutations), nil
}

func parseSpringLine(line []byte) SpringLine {
	springStr := strings.Split(string(line), " ")
	if len(springStr) != 2 {
		return SpringLine{}
	}

	springs := strings.Split(springStr[0], "")
	groupingsStr := strings.Split(springStr[1], ",")
	groupings := make([]int, len(groupingsStr))
	for i := 0; i < len(groupings); i++ {
		grouping, err := strconv.Atoi(groupingsStr[i])
		if err == nil {
			groupings[i] = grouping
		}
	}

	return SpringLine{
		springs:   springs,
		groupings: groupings,
	}
}

func parseSpringLineLong(line []byte) SpringLine {
	springStr := strings.Split(string(line), " ")
	if len(springStr) != 2 {
		return SpringLine{}
	}

	springs := strings.Split(springStr[0], "")
	groupingsStr := strings.Split(springStr[1], ",")
	groupings := make([]int, len(groupingsStr))
	for i := 0; i < len(groupings); i++ {
		grouping, err := strconv.Atoi(groupingsStr[i])
		if err == nil {
			groupings[i] = grouping
		}
	}

	expandedGroupingsList := make([]int, len(groupings)*5)
	for i := range expandedGroupingsList {
		expandedGroupingsList[i] = groupings[i%len(groupings)]
	}

	expandedSpringList := make([]string, ((len(springs)+1)*5)-1)
	springIdx := 0
	for i := range expandedSpringList {
		if springIdx == len(springs) {
			expandedSpringList[i] = "?"
			springIdx = 0
			continue
		}

		expandedSpringList[i] = springs[springIdx]
		springIdx++
	}

	return SpringLine{
		springs:   expandedSpringList,
		groupings: expandedGroupingsList,
	}
}

func calculateSpringPermutations(springLine SpringLine) int {
	unknownSprings := []int{}
	for i, spring := range springLine.springs {
		if spring == "?" {
			unknownSprings = append(unknownSprings, i)
		}
	}

	// Creating starting permutation
	perm := make([]string, len(unknownSprings))
	for i := 0; i < len(unknownSprings); i++ {
		perm[i] = "?"
	}

	return countPermutations(perm, unknownSprings, springLine)
}

func calculateSpringPermutationsFast(springLine SpringLine) int {
	unknownSprings := []int{}
	for i, spring := range springLine.springs {
		if spring == "?" {
			unknownSprings = append(unknownSprings, i)
		}
	}

	// Creating starting permutation
	perm := make([]string, len(unknownSprings))
	for i := 0; i < len(unknownSprings); i++ {
		perm[i] = "?"
	}

	return countPermutationsFast(perm, unknownSprings, springLine)
}

func countPermutations(perm []string, unknownSprings []int, springLine SpringLine) int {
	for i := 0; i < len(perm); i++ {
		if perm[i] == "?" {
			opPerm := make([]string, len(perm))
			dmgPerm := make([]string, len(perm))

			copy(opPerm, perm)
			copy(dmgPerm, perm)

			opPerm[i] = "."
			dmgPerm[i] = "#"

			opPermCount := countPermutations(opPerm, unknownSprings, springLine)
			dmgPermCount := countPermutations(dmgPerm, unknownSprings, springLine)

			return opPermCount + dmgPermCount
		}
	}

	// validate permutation matches grouping
	filledSprings := make([]string, len(springLine.springs))
	copy(filledSprings, springLine.springs)
	springsfilled := 0
	for i := range filledSprings {
		if filledSprings[i] == "?" {
			filledSprings[i] = perm[springsfilled]
			springsfilled++
		}
	}

	springGroups := []int{}
	groupedSpringCount := 0
	for i := range filledSprings {
		if filledSprings[i] == "#" {
			groupedSpringCount++
		} else if groupedSpringCount > 0 {
			springGroups = append(springGroups, groupedSpringCount)
			groupedSpringCount = 0
		}
	}

	if groupedSpringCount > 0 {
		springGroups = append(springGroups, groupedSpringCount)
	}

	for i := range springLine.groupings {
		if len(springLine.groupings) != len(springGroups) || springLine.groupings[i] != springGroups[i] {
			return 0
		}
	}

	return 1
}

func countPermutationsFast(perm []string, unknownSprings []int, springLine SpringLine) int {
	// fmt.Println(perm)
	for i := 0; i < len(perm); i++ {
		if perm[i] == "?" {
			opPerm := make([]string, len(perm))
			dmgPerm := make([]string, len(perm))

			copy(opPerm, perm)
			copy(dmgPerm, perm)

			opPerm[i] = "."
			dmgPerm[i] = "#"

			// fmt.Println(opPerm)
			opValid := checkPermValid(opPerm[:i+1], springLine)
			dmgValid := checkPermValid(dmgPerm[:i+1], springLine)

			// if i > 5 {
			// 	return 0
			// }

			totalPermCount := 0
			if opValid {
				totalPermCount += countPermutationsFast(opPerm, unknownSprings, springLine)
			}
			if dmgValid {
				totalPermCount += countPermutationsFast(dmgPerm, unknownSprings, springLine)
			}

			return totalPermCount

			// Check if opPerm and dmgPerm are valid

			// opPermCount := countPermutations(opPerm, unknownSprings, springLine)
			// dmgPermCount := countPermutations(dmgPerm, unknownSprings, springLine)

			// return opPermCount + dmgPermCount
		}
	}

	// fmt.Println(perm)

	fmt.Println(perm)

	if checkPermValid(perm, springLine) {
		return 1
	}

	return 0

	// validate permutation matches grouping
	// filledSprings := make([]string, len(springLine.springs))
	// copy(filledSprings, springLine.springs)
	// springsfilled := 0
	// for i := range filledSprings {
	// 	if filledSprings[i] == "?" {
	// 		filledSprings[i] = perm[springsfilled]
	// 		springsfilled++
	// 	}
	// }

	// springGroups := []int{}
	// groupedSpringCount := 0
	// for i := range filledSprings {
	// 	if filledSprings[i] == "#" {
	// 		groupedSpringCount++
	// 	} else if groupedSpringCount > 0 {
	// 		springGroups = append(springGroups, groupedSpringCount)
	// 		groupedSpringCount = 0
	// 	}
	// }

	// if groupedSpringCount > 0 {
	// 	springGroups = append(springGroups, groupedSpringCount)
	// }

	// for i := range springLine.groupings {
	// 	if len(springLine.groupings) != len(springGroups) || springLine.groupings[i] != springGroups[i] {
	// 		return 0
	// 	}
	// }

	// return 1
}

// permSubArray is the array of all springs that have been set in the permutation
func checkPermValid(permSubArray []string, springLine SpringLine) bool {
	// fmt.Println(permSubArray)

	filledSprings := make([]string, len(springLine.springs))
	copy(filledSprings, springLine.springs)

	currentCharIdx := 0
	for i := range filledSprings {
		if currentCharIdx == len(permSubArray) {
			break
		}

		if filledSprings[i] == "?" {
			filledSprings[i] = permSubArray[currentCharIdx]
			currentCharIdx++
		}
	}

	// fmt.Println(springLine.springs)
	// fmt.Println(filledSprings)

	currentSpringGroup := 0
	groupedSprings := 0
	for i := range filledSprings {
		if filledSprings[i] == "?" {
			break
		}

		if filledSprings[i] == "#" {
			groupedSprings++
		}

		if filledSprings[i] == "." && groupedSprings > 0 {
			currentSpringGroup++
			groupedSprings = 0
		}

		if currentSpringGroup >= len(springLine.groupings) || groupedSprings > springLine.groupings[currentSpringGroup] {
			// fmt.Println("invalid:", filledSprings)
			return false
		}
	}
	// fmt.Println("ok:", filledSprings)
	return true
}
