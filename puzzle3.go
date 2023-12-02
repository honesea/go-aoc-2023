package main

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strconv"
	"strings"
)

type CubeSet struct {
	red int
	green int
	blue int
}

type CubeGame struct {
	id int
	cubeSets []CubeSet
}

const maxRedCubes = 12
const maxGreenCubes = 13
const maxBlueCubes = 14

func puzzle3() (int, error) {
	file, err := os.Open("inputs/puzzle3.txt")
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

	gameIdTotal := 0
	OUTER:
	for _, game := range games {
		for _, set := range game.cubeSets {
			if set.red > maxRedCubes || set.green > maxGreenCubes || set.blue > maxBlueCubes {
				continue OUTER
			}
		}

		gameIdTotal += game.id
	}

	return gameIdTotal, nil
}

func parseCubeGame(line []byte) (CubeGame, error) {
	gameStr := strings.Split(string(line), ":")
	if len(gameStr) != 2 {
		return CubeGame{}, errors.New("game line formatted incorrectly")
	}

	gameIdStr := strings.Split(gameStr[0], " ")
	if len(gameIdStr) != 2 {
		return CubeGame{}, errors.New("game line formatted incorrectly")
	}

	gameId, err := strconv.Atoi(gameIdStr[1])
	if err != nil {
		return CubeGame{}, errors.New("game line formatted incorrectly")
	}

	cubeGame := CubeGame{
		id: gameId,
		cubeSets: []CubeSet{},
	}
	
	for _, cubeSetStr := range strings.Split(gameStr[1], ";") {
		cubeSet := CubeSet{}

		for _, cubeCountStr := range strings.Split(cubeSetStr, ",") {
			cubeCountValStr := strings.Split(cubeCountStr, " ")

			// Ignore empty or badly formatted tokens
			if len(cubeCountValStr) != 3 { // white space + count + color
				continue
			}

			val, err := strconv.Atoi(cubeCountValStr[1])
			if err != nil {
				continue
			}

			switch color := cubeCountValStr[2]; color {
			case "red":
				cubeSet.red = val
			case "green":
				cubeSet.green = val
			case "blue":
				cubeSet.blue = val
			}
		}

		cubeGame.cubeSets = append(cubeGame.cubeSets, cubeSet)
	}

	return cubeGame, nil
}