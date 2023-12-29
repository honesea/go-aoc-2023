package main

import (
	"bufio"
	"errors"
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
	for _, springLine := range springLines {
		cache := make(map[[3]int]int)
		permutations = append(permutations, countPermutationsDP(springLine.springs, springLine.groupings, 0, 0, 0, cache))
	}

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
		cache := make(map[[3]int]int)
		permutations = append(permutations, countPermutationsDP(springLine.springs, springLine.groupings, 0, 0, 0, cache))

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

func countPermutationsDP(springs []string, groupings []int, pos, consecutiveHash, groupIdx int, cache map[[3]int]int) int {
	if pos == len(springs) {
		if consecutiveHash > 0 && groupings[groupIdx] != consecutiveHash {
			return 0
		}

		// Handle final group index increment
		if consecutiveHash > 0 {
			groupIdx++
		}

		if groupIdx != len(groupings) {
			return 0
		}

		return 1
	}

	// Check if this state has already been computed
	state := [3]int{pos, consecutiveHash, groupIdx}
	if val, exists := cache[state]; exists {
		return val
	}

	totalPermutations := 0
	tryBothPaths := false

	if springs[pos] == "?" {
		tryBothPaths = true
	}

	if tryBothPaths || springs[pos] == "." {
		if consecutiveHash > 0 && groupings[groupIdx] == consecutiveHash {
			totalPermutations += countPermutationsDP(springs, groupings, pos+1, 0, groupIdx+1, cache)
		} else if consecutiveHash == 0 {
			totalPermutations += countPermutationsDP(springs, groupings, pos+1, consecutiveHash, groupIdx, cache)
		}
	}

	if tryBothPaths || springs[pos] == "#" {
		if groupIdx < len(groupings) && consecutiveHash+1 <= groupings[groupIdx] {
			totalPermutations += countPermutationsDP(springs, groupings, pos+1, consecutiveHash+1, groupIdx, cache)
		}
	}

	cache[state] = totalPermutations
	return totalPermutations
}
