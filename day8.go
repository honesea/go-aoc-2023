package main

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strings"
)

func day8() (int, int, error) {
	answerP1, err := d8p1()
	if err != nil {
		return 0, 0, err
	}

	answerP2, err := d8p2()
	if err != nil {
		return 0, 0, err
	}

	return answerP1, answerP2, err
}

func d8p1() (int, error) {
	file, err := os.Open("inputs/d8p1.txt")
	if err != nil {
		return 0, errors.New("could not read file input")
	}

	defer file.Close()
	reader := bufio.NewReader(file)

	instructions := []string{}
	nodes := map[string]struct {
		left  string
		right string
	}{}

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

		if len(instructions) == 0 {
			instructions = strings.Split(string(line), "")
		}

		if len(line) == 0 {
			continue
		}

		nodeId, left, right, err := parseDesertMapNode(line)
		if err != nil {
			continue
		}

		nodes[nodeId] = struct {
			left  string
			right string
		}{
			left:  left,
			right: right,
		}
	}

	step := 0
	currentNode := "AAA"
	for {
		instruction := instructions[step%len(instructions)]

		if instruction == "L" {
			currentNode = nodes[currentNode].left
		} else {
			currentNode = nodes[currentNode].right
		}

		step++
		if currentNode == "ZZZ" {
			break
		}
	}

	return step, nil
}

func d8p2() (int, error) {
	file, err := os.Open("inputs/d8p2.txt")
	if err != nil {
		return 0, errors.New("could not read file input")
	}

	defer file.Close()
	reader := bufio.NewReader(file)

	instructions := []string{}
	nodes := map[string]struct {
		left  string
		right string
	}{}

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

		if len(instructions) == 0 {
			instructions = strings.Split(string(line), "")
		}

		if len(line) == 0 {
			continue
		}

		nodeId, left, right, err := parseDesertMapNode(line)
		if err != nil {
			continue
		}

		nodes[nodeId] = struct {
			left  string
			right string
		}{
			left:  left,
			right: right,
		}
	}

	startingNodes := []string{}
	for id := range nodes {
		if id[2:] == "A" {
			startingNodes = append(startingNodes, id)
		}
	}

	cycleLengths := []int{}
	for _, node := range startingNodes {
		step := 0

		currentNode := node
		for currentNode[2:] != "Z" {
			instruction := instructions[step%len(instructions)]

			if instruction == "L" {
				currentNode = nodes[currentNode].left
			} else {
				currentNode = nodes[currentNode].right
			}

			step++
		}

		cycleLengths = append(cycleLengths, step)
	}

	totalCycleLength := lcm(cycleLengths...)

	return totalCycleLength, nil
}

func parseDesertMapNode(line []byte) (nodeId, left, right string, err error) {
	nodeStr := strings.Split(string(line), " = ")
	if len(nodeStr) != 2 {
		return "", "", "", errors.New("incorrect line format")
	}

	nodePathsStr := strings.Split(nodeStr[1], ", ")
	if len(nodeStr) != 2 {
		return "", "", "", errors.New("incorrect line format")
	}

	nodeId = nodeStr[0]
	left = nodePathsStr[0][1:]
	right = nodePathsStr[1][:len(nodePathsStr[1])-1]

	return nodeId, left, right, nil
}

// lcm calculates the Least Common Multiple of integers
func lcm(nums ...int) int {
	result := nums[0]
	for i := 1; i < len(nums); i++ {
		result = lcmOfTwo(result, nums[i])
	}
	return result
}

// lcmOfTwo calculates the Least Common Multiple of two integers
func lcmOfTwo(a, b int) int {
	return a * b / gcd(a, b)
}

// gcd calculates the Greatest Common Divisor of two integers
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
