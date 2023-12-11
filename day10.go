package main

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strings"
)

type Pipe struct {
	character string
	north     bool
	east      bool
	south     bool
	west      bool
}

var pipeMap map[string]Pipe = map[string]Pipe{
	"|": {
		character: "|",
		north:     true,
		east:      false,
		south:     true,
		west:      false,
	},
	"-": {
		character: "-",
		north:     false,
		east:      true,
		south:     false,
		west:      true,
	},
	"L": {
		character: "L",
		north:     true,
		east:      true,
		south:     false,
		west:      false,
	},
	"J": {
		character: "J",
		north:     true,
		east:      false,
		south:     false,
		west:      true,
	},
	"7": {
		character: "7",
		north:     false,
		east:      false,
		south:     true,
		west:      true,
	},
	"F": {
		character: "F",
		north:     false,
		east:      true,
		south:     true,
		west:      false,
	},
	".": {
		character: ".",
		north:     false,
		east:      false,
		south:     false,
		west:      false,
	},
	"S": {
		character: "S",
		north:     true,
		east:      true,
		south:     true,
		west:      true,
	},
}

func day10() (int, int, error) {
	answerP1, err := d10p1()
	if err != nil {
		return 0, 0, err
	}

	answerP2, err := d10p2()
	if err != nil {
		return 0, 0, err
	}

	return answerP1, answerP2, err
}

func d10p1() (int, error) {
	file, err := os.Open("inputs/d10p1.txt")
	if err != nil {
		return 0, errors.New("could not read file input")
	}

	defer file.Close()
	reader := bufio.NewReader(file)

	pipeMatrix := [][]Pipe{}
	start := Coord{}

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

		pipesStr := strings.Split(string(line), "")

		pipeLine := make([]Pipe, len(pipesStr))
		for i, pipeStr := range pipesStr {
			pipe, ok := pipeMap[pipeStr]
			if ok {
				pipeLine[i] = pipe
			}
		}

		for x, pipe := range pipeLine {
			if pipe.character == "S" {
				start = Coord{
					x: x,
					y: len(pipeMatrix),
				}
			}
		}

		pipeMatrix = append(pipeMatrix, pipeLine)
	}

	loop := findLoop(pipeMatrix, start)
	return len(loop) / 2, nil
}

func d10p2() (int, error) {
	file, err := os.Open("inputs/d10p2.txt")
	if err != nil {
		return 0, errors.New("could not read file input")
	}

	defer file.Close()
	reader := bufio.NewReader(file)

	pipeMatrix := [][]Pipe{}
	start := Coord{}

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

		pipesStr := strings.Split(string(line), "")

		pipeLine := make([]Pipe, len(pipesStr))
		for i, pipeStr := range pipesStr {
			pipe, ok := pipeMap[pipeStr]
			if ok {
				pipeLine[i] = pipe
			}
		}

		for x, pipe := range pipeLine {
			if pipe.character == "S" {
				start = Coord{
					x: x,
					y: len(pipeMatrix),
				}
			}
		}

		pipeMatrix = append(pipeMatrix, pipeLine)
	}

	loop := findLoop(pipeMatrix, start)

	for y := range pipeMatrix {
		for x := range pipeMatrix[y] {
			if !loopContainsCoord(loop, x, y) {
				pipeMatrix[y][x] = Pipe{
					character: ".",
					north:     false,
					east:      false,
					south:     false,
					west:      false,
				}
			}

			if pipeMatrix[y][x].character == "S" {
				pipeMatrix[y][x] = replaceStartingPipe(pipeMatrix, x, y)
			}
		}
	}

	tilesInside := 0
	for y, pipeLine := range pipeMatrix {
		intersections := 0
		lastChar := ""

		for x := range pipeLine {
			// horizontal pipes are not edges or empty sections
			if pipeMatrix[y][x].character == "-" {
				continue
			}

			// Count empty section if intersections are odd meaning tile is inside
			if intersections%2 == 1 && pipeMatrix[y][x].character == "." {
				lastChar = pipeMatrix[y][x].character
				tilesInside++
				continue
			}

			// If pipe is part of a diagonal then don't conunt as an intersection
			if (pipeMatrix[y][x].character == "7" && lastChar == "L") || (pipeMatrix[y][x].character == "J" && lastChar == "F") {
				lastChar = pipeMatrix[y][x].character
				continue
			}

			// Pipe piece is an edge
			if loopContainsCoord(loop, x, y) {
				lastChar = pipeMatrix[y][x].character
				intersections++
				continue
			}

			lastChar = pipeMatrix[y][x].character
		}
	}

	return tilesInside, nil
}

func findLoop(pipeMatrix [][]Pipe, start Coord) []Coord {
	possibleStartingCoords := getSurroundingCoords(start)

	pipesConnectedToStart := []Coord{}
	for _, possibleStartingCoord := range possibleStartingCoords {
		if canTraverse(start, start, possibleStartingCoord, pipeMatrix) {
			pipesConnectedToStart = append(pipesConnectedToStart, possibleStartingCoord)
		}
	}

	prev := start
	current := pipesConnectedToStart[0]
	last := pipesConnectedToStart[1]
	loop := []Coord{prev, current}

	for {
		for _, next := range getSurroundingCoords(current) {
			if canTraverse(prev, current, next, pipeMatrix) {
				prev = current
				current = next
				loop = append(loop, current)
				break
			}
		}

		if current == last {
			break
		}
	}

	return loop
}

func getSurroundingCoords(start Coord) []Coord {
	return []Coord{
		{x: start.x, y: start.y - 1},
		{x: start.x + 1, y: start.y},
		{x: start.x, y: start.y + 1},
		{x: start.x - 1, y: start.y},
	}
}

func canTraverse(prev, current, next Coord, pipeMatrix [][]Pipe) bool {
	if prev.x == next.x && prev.y == next.y {
		return false
	}

	// Traversing north
	if current.y > next.y && next.y >= 0 {
		return pipeMatrix[current.y][current.x].north && pipeMatrix[next.y][next.x].south
	}

	// Traversing east
	if current.x < next.x && next.x < len(pipeMatrix[0]) {
		return pipeMatrix[current.y][current.x].east && pipeMatrix[next.y][next.x].west
	}

	// Traversing south
	if current.y < next.y && next.y < len(pipeMatrix) {
		return pipeMatrix[current.y][current.x].south && pipeMatrix[next.y][next.x].north
	}

	// Traversing west
	if current.x > next.x && next.x >= 0 {
		return pipeMatrix[current.y][current.x].west && pipeMatrix[next.y][next.x].east
	}

	return false
}

func replaceStartingPipe(pipeMatrix [][]Pipe, x, y int) Pipe {
	if y-1 >= 0 && x+1 < len(pipeMatrix[0]) && pipeMatrix[y-1][x].south && pipeMatrix[y][x+1].west {
		return pipeMap["L"]
	}

	if y+1 < len(pipeMatrix) && x+1 < len(pipeMatrix[0]) && pipeMatrix[y+1][x].north && pipeMatrix[y][x+1].west {
		return pipeMap["F"]
	}

	if y+1 < len(pipeMatrix) && x-1 >= 0 && pipeMatrix[y+1][x].north && pipeMatrix[y][x-1].east {
		return pipeMap["7"]
	}

	if y-1 >= 0 && x-1 >= 0 && pipeMatrix[y-1][x].south && pipeMatrix[y][x-1].east {
		return pipeMap["J"]
	}

	return Pipe{}
}

func loopContainsCoord(loop []Coord, x, y int) bool {
	for _, coord := range loop {
		if coord.x == x && coord.y == y {
			return true
		}
	}

	return false
}
