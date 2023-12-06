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

func day5() (int, int, error) {
	answerP1, err := d5p1()
	if err != nil {
		return 0, 0, err
	}

	answerP2, err := d5p2()
	if err != nil {
		return 0, 0, err
	}

	return answerP1, answerP2, err
}

type AlmanacMap struct {
	source string
	dest   string
	ranges []AlmanacMapRange
}

type AlmanacMapRange struct {
	destStart   int
	sourceStart int
	mapRange    int
}

func d5p1() (int, error) {
	file, err := os.Open("inputs/d5p1.txt")
	if err != nil {
		return 0, errors.New("could not read file input")
	}

	defer file.Close()
	reader := bufio.NewReader(file)

	seeds := []int{}
	maps := []AlmanacMap{}
	currentMap := AlmanacMap{}

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

		// Ignore empty line
		if len(line) == 0 {
			continue
		}

		// Populate seed list
		if len(seeds) == 0 {
			seeds = parseSeeds(line)
		}

		// Start of a new map range
		mapNameStr := strings.Split(string(line), " ")
		if len(mapNameStr) == 2 && mapNameStr[1] == "map:" {
			// Save current map
			if currentMap.source != "" && currentMap.dest != "" {
				maps = append(maps, currentMap)
			}

			objectStr := strings.Split(mapNameStr[0], "-")
			if len(objectStr) == 3 {
				source := objectStr[0]
				dest := objectStr[2]
				currentMap = AlmanacMap{
					source: source,
					dest:   dest,
					ranges: []AlmanacMapRange{},
				}
			}
		}

		mapRange := parseAlmanacMapRange(line)
		if mapRange != (AlmanacMapRange{}) {
			currentMap.ranges = append(currentMap.ranges, parseAlmanacMapRange(line))
		}
	}

	// Save final map
	if currentMap.source != "" && currentMap.dest != "" {
		maps = append(maps, currentMap)
	}

	location := math.MaxInt
	for _, seed := range seeds {
		newLocation := findSeedLocation(maps, seed)
		if newLocation < location {
			location = newLocation
		}
	}

	return location, nil
}

func d5p2() (int, error) {
	file, err := os.Open("inputs/d5p2.txt")
	if err != nil {
		return 0, errors.New("could not read file input")
	}

	defer file.Close()
	reader := bufio.NewReader(file)

	seedRanges := []struct {
		start int
		end   int
	}{}
	maps := []AlmanacMap{}
	currentMap := AlmanacMap{}

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

		// Ignore empty line
		if len(line) == 0 {
			continue
		}

		// Populate seed list
		if len(seedRanges) == 0 {
			seedRanges = parseSeedRanges(line)
		}

		// Start of a new map range
		mapNameStr := strings.Split(string(line), " ")
		if len(mapNameStr) == 2 && mapNameStr[1] == "map:" {
			// Save current map
			if currentMap.source != "" && currentMap.dest != "" {
				maps = append(maps, currentMap)
			}

			objectStr := strings.Split(mapNameStr[0], "-")
			if len(objectStr) == 3 {
				source := objectStr[0]
				dest := objectStr[2]
				currentMap = AlmanacMap{
					source: source,
					dest:   dest,
					ranges: []AlmanacMapRange{},
				}
			}
		}

		mapRange := parseAlmanacMapRange(line)
		if mapRange != (AlmanacMapRange{}) {
			currentMap.ranges = append(currentMap.ranges, parseAlmanacMapRange(line))
		}
	}

	// Save final map
	if currentMap.source != "" && currentMap.dest != "" {
		maps = append(maps, currentMap)
	}

	// Should be optimised by doing range check rather than iterating through seed id
	location := math.MaxInt
	for _, seedRange := range seedRanges {
		for i := seedRange.start; i < seedRange.end; i++ {
			newLocation := findSeedLocation(maps, i)
			if newLocation < location {
				location = newLocation
			}
		}
	}

	return location, nil
}

func parseSeeds(line []byte) []int {
	seedStr := strings.Split(string(line), ":")
	if len(seedStr) != 2 {
		return []int{}
	}

	seedNumbersStr := strings.Split(seedStr[1], " ")

	seeds := []int{}
	for _, seedNumberStr := range seedNumbersStr {
		seedNumber, err := strconv.Atoi(seedNumberStr)
		if err == nil {
			seeds = append(seeds, seedNumber)
		}
	}

	return seeds
}

func parseSeedRanges(line []byte) []struct {
	start int
	end   int
} {
	seedStr := strings.Split(string(line), ":")
	if len(seedStr) != 2 {
		return []struct {
			start int
			end   int
		}{}
	}

	seedNumbersStr := strings.Split(seedStr[1], " ")

	seeds := []int{}
	for _, seedNumberStr := range seedNumbersStr {
		seedNumber, err := strconv.Atoi(seedNumberStr)
		if err == nil {
			seeds = append(seeds, seedNumber)
		}
	}

	seedRanges := []struct {
		start int
		end   int
	}{}
	for i := 0; i < len(seeds); i += 2 {
		rangeStart := seeds[i]
		rangeLength := seeds[i+1]
		seedRange := struct {
			start int
			end   int
		}{
			start: rangeStart,
			end:   rangeStart + rangeLength,
		}
		seedRanges = append(seedRanges, seedRange)
	}

	return seedRanges
}

func parseAlmanacMapRange(line []byte) AlmanacMapRange {
	rangeStr := strings.Split(string(line), " ")
	if len(rangeStr) != 3 {
		return AlmanacMapRange{}
	}
	destStart, destErr := strconv.Atoi(rangeStr[0])
	sourceStart, sourceErr := strconv.Atoi(rangeStr[1])
	mapRange, rangeErr := strconv.Atoi(rangeStr[2])
	if destErr != nil || sourceErr != nil || rangeErr != nil {
		return AlmanacMapRange{}
	}

	return AlmanacMapRange{
		destStart:   destStart,
		sourceStart: sourceStart,
		mapRange:    mapRange,
	}
}

func findSeedLocation(maps []AlmanacMap, seed int) int {
	currentObj := "seed"
	currentId := seed

	for {
		if currentObj == "location" {
			break
		}

		currentMap := AlmanacMap{}
		for _, almanacMap := range maps {
			if almanacMap.source == currentObj {
				currentMap = almanacMap
				break
			}
		}

		// Couldn't find object map
		if currentMap.source == "" && currentMap.dest == "" {
			return 0
		}

		// Find source object range
		currentRange := AlmanacMapRange{}
		for _, amr := range currentMap.ranges {
			if currentId >= amr.sourceStart && currentId < amr.sourceStart+amr.mapRange {
				currentRange = amr
				break
			}
		}

		// Update object id if range found
		if currentRange != (AlmanacMapRange{}) {
			currentId = currentRange.destStart + (currentId - currentRange.sourceStart)
		}

		currentObj = currentMap.dest
	}

	return currentId
}
