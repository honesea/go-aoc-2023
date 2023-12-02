package main

import (
	"fmt"
	"os"
	"strconv"
)

var puzzles = map[int]func()(int, error){
	1: puzzle1,
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

	puzzle, ok := puzzles[puzzleArg]
	if !ok {
		fmt.Println("Error: please enter a valid puzzle number as an arg")
		return
	}

	answer, err := puzzle()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Answer: %v\n", answer)
	}
}