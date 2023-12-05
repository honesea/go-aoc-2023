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

type ScratchCard struct {
	id             int
	winningNumbers []int
	gameNumbers    []int
	matches        int
}

func day4() (int, int, error) {
	answerP1, err := d4p1()
	if err != nil {
		return 0, 0, err
	}

	answerP2, err := d4p2()
	if err != nil {
		return 0, 0, err
	}

	return answerP1, answerP2, err
}

func d4p1() (int, error) {
	file, err := os.Open("inputs/d4p1.txt")
	if err != nil {
		return 0, errors.New("could not read file input")
	}

	defer file.Close()
	reader := bufio.NewReader(file)

	scratchCards := []ScratchCard{}

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

		scratchCard, err := parseScratchCard(line)
		if err != nil {
			return 0, err
		}

		scratchCards = append(scratchCards, scratchCard)
	}

	scores := []int{}
	for _, card := range scratchCards {
		if card.matches == 0 {
			continue
		}

		if card.matches == 1 {
			scores = append(scores, 1)
			continue
		}

		scores = append(scores, int(math.Pow(2, float64(card.matches-1))))
	}

	return sum(scores), nil
}

func d4p2() (int, error) {
	file, err := os.Open("inputs/d4p2.txt")
	if err != nil {
		return 0, errors.New("could not read file input")
	}

	defer file.Close()
	reader := bufio.NewReader(file)

	scratchCards := map[int][]ScratchCard{}

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

		scratchCard, err := parseScratchCard(line)
		if err != nil {
			return 0, err
		}

		scratchCards[scratchCard.id] = []ScratchCard{scratchCard}
	}

	for id := 1; id <= len(scratchCards); id++ {
		cards := scratchCards[id]

		for _, card := range cards {
			if card.matches == 0 {
				continue
			}

			for i := id + 1; i <= id+card.matches; i++ {
				cardCopy := scratchCards[i][0]
				scratchCards[i] = append(scratchCards[i], cardCopy)
			}
		}
	}

	numCards := []int{}
	for id := 1; id <= len(scratchCards); id++ {
		cards := scratchCards[id]
		numCards = append(numCards, len(cards))
	}

	return sum(numCards), nil
}

func parseScratchCard(line []byte) (ScratchCard, error) {
	cardStr := strings.Split(string(line), ":")
	if len(cardStr) != 2 {
		return ScratchCard{}, errors.New("game line formatted incorrectly")
	}

	cardIdStr := strings.Split(cardStr[0], " ")
	if len(cardIdStr) < 2 {
		return ScratchCard{}, errors.New("game line formatted incorrectly")
	}

	cardId, err := strconv.Atoi(cardIdStr[len(cardIdStr)-1])
	if err != nil {
		return ScratchCard{}, errors.New("game line formatted incorrectly")
	}

	scratchCard := ScratchCard{
		id:             cardId,
		winningNumbers: []int{},
		gameNumbers:    []int{},
		matches:        0,
	}

	gameStr := strings.Split(cardStr[1], "|")
	if len(gameStr) != 2 {
		return ScratchCard{}, errors.New("game line formatted incorrectly")
	}

	scratchCard.winningNumbers = parseGameNumbers(gameStr[0])
	scratchCard.gameNumbers = parseGameNumbers(gameStr[1])

	for _, winningNumber := range scratchCard.winningNumbers {
		for _, gameNumber := range scratchCard.gameNumbers {
			if winningNumber == gameNumber {
				scratchCard.matches++
				break
			}
		}
	}

	return scratchCard, nil
}

func parseGameNumbers(game string) []int {
	numbers := []int{}

	for _, numberStr := range strings.Split(game, " ") {
		number, err := strconv.Atoi(numberStr)
		if err == nil {
			numbers = append(numbers, number)
		}
	}

	return numbers
}
