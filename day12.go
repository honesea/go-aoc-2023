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

type CachedSpringLine struct {
	groupings      []int
	dmgSpringCount int
	currentIndex   int
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
		cache := make(map[string]CachedSpringLine)
		permutations = append(permutations, countPermutations(cache, springLine, ""))
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
	for i, springLine := range springLines {
		fmt.Printf("%v / %v\n", i, len(springLines))
		cache := make(map[string]CachedSpringLine)
		permutations = append(permutations, countPermutations(cache, springLine, ""))
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

func countPermutations(cache map[string]CachedSpringLine, springLine SpringLine, perm string) int {
	cachedSpringLine := getSpringLineFromCache(cache, perm)

	fmt.Println(perm)

	// Handle most recent addition to permutation
	if perm != "" {
		if perm[len(perm)-1] == '#' {
			cachedSpringLine.dmgSpringCount++
		} else if perm[len(perm)-1] == '.' && cachedSpringLine.dmgSpringCount > 0 {
			cachedSpringLine.groupings = append(cachedSpringLine.groupings, cachedSpringLine.dmgSpringCount)
			cachedSpringLine.dmgSpringCount = 0
		}

		cachedSpringLine.currentIndex++
	}

	for {

		// End of string so calculate if correct and return count
		if cachedSpringLine.currentIndex == len(springLine.springs) {
			break
		}

		character := springLine.springs[cachedSpringLine.currentIndex]

		if character == "?" {
			cache[perm] = cachedSpringLine
			return countPermutations(cache, springLine, perm+"#") + countPermutations(cache, springLine, perm+".")
		} else if character == "#" {
			cachedSpringLine.dmgSpringCount++
		} else if character == "." && cachedSpringLine.dmgSpringCount > 0 {
			cachedSpringLine.groupings = append(cachedSpringLine.groupings, cachedSpringLine.dmgSpringCount)
			cachedSpringLine.dmgSpringCount = 0
		}

		cachedSpringLine.currentIndex++
	}

	if cachedSpringLine.dmgSpringCount > 0 {
		cachedSpringLine.groupings = append(cachedSpringLine.groupings, cachedSpringLine.dmgSpringCount)
		cachedSpringLine.dmgSpringCount = 0
	}

	if len(springLine.groupings) != len(cachedSpringLine.groupings) {
		return 0
	}

	for i := 0; i < len(springLine.groupings); i++ {
		if springLine.groupings[i] != cachedSpringLine.groupings[i] {
			return 0
		}
	}

	return 1
}

func getSpringLineFromCache(cache map[string]CachedSpringLine, perm string) CachedSpringLine {
	if len(perm) == 0 {
		return CachedSpringLine{}
	}

	cachedPerm := perm[:len(perm)-1]

	groups, ok := cache[cachedPerm]
	if ok {
		return groups
	}

	return CachedSpringLine{}
}
