package main

import (
	"fmt"
	"os"
	"strconv"
)

var puzzles = map[int]func() (int, int, error){
	1:  day1,
	2:  day2,
	3:  day3,
	4:  day4,
	5:  day5,
	6:  day6,
	7:  day7,
	8:  day8,
	9:  day9,
	10: day10,
	11: day11,
	12: day12,
	13: day13,
	14: day14,
}

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Error: please enter a valid puzzle number as an arg")
		return
	}

	puzzleArg, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Error: please enter a valid puzzle number as an arg")
		return
	}

	day, ok := puzzles[puzzleArg]
	if !ok {
		fmt.Println("Error: please enter a valid puzzle number as an arg")
		return
	}

	answerP1, answerP2, err := day()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Answer p1: %v\nAnswer p2: %v\n", answerP1, answerP2)
	}
}
