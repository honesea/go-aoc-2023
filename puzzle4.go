package main

import (
	"bufio"
	"errors"
	"io"
	"os"
)

// Using types and parser from puzzle 3

func puzzle4() (int, error) {
	file, err := os.Open("inputs/puzzle4.txt")
	if err != nil {
		return 0, errors.New("could not read file input")
	}

	defer file.Close()
	reader := bufio.NewReader(file)

	games := []CubeGame{}

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

		cubeGame, err := parseCubeGame(line)
		if err != nil {
			return 0, err
		}
		
		games = append(games, cubeGame)
	}

	gamePowersTotal := 0
	for _, game := range games {
		minRedCubes := 0
		minGreenCubes := 0
		minBlueCubes := 0

		for _, set := range game.cubeSets {
			minRedCubes = max(minRedCubes, set.red)
			minGreenCubes = max(minGreenCubes, set.green)
			minBlueCubes = max(minBlueCubes, set.blue)
		}

		gamePower := minRedCubes * minGreenCubes * minBlueCubes
		gamePowersTotal += gamePower
	}

	return gamePowersTotal, nil
}