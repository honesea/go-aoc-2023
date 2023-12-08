package main

import (
	"bufio"
	"errors"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

type CamelGameHand struct {
	cards []string
	bid   int
}

var cardRanks = map[string]int{
	"2": 1,
	"3": 2,
	"4": 3,
	"5": 4,
	"6": 5,
	"7": 6,
	"8": 7,
	"9": 8,
	"T": 9,
	"J": 10,
	"Q": 11,
	"K": 12,
	"A": 13,
}

var jokerCardRanks = map[string]int{
	"J": 0,
	"2": 1,
	"3": 2,
	"4": 3,
	"5": 4,
	"6": 5,
	"7": 6,
	"8": 7,
	"9": 8,
	"T": 9,
	"Q": 11,
	"K": 12,
	"A": 13,
}

func day7() (int, int, error) {
	answerP1, err := d7p1()
	if err != nil {
		return 0, 0, err
	}

	answerP2, err := d7p2()
	if err != nil {
		return 0, 0, err
	}

	return answerP1, answerP2, err
}

func d7p1() (int, error) {
	file, err := os.Open("inputs/d7p1.txt")
	if err != nil {
		return 0, errors.New("could not read file input")
	}

	defer file.Close()
	reader := bufio.NewReader(file)

	hands := []CamelGameHand{}

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

		handStr := strings.Split(string(line), " ")
		if len(handStr) != 2 {
			continue
		}

		cards := strings.Split(handStr[0], "")
		bid, err := strconv.Atoi(handStr[1])
		if err != nil {
			continue
		}

		hand := CamelGameHand{
			cards: cards,
			bid:   bid,
		}
		hands = append(hands, hand)
	}

	sort.Slice(hands, func(i, j int) bool {
		a := hands[i]
		b := hands[j]

		return compareCamelGameHand(a, b)
	})

	handsTotal := 0
	for i := 1; i <= len(hands); i++ {
		handsTotal += i * hands[i-1].bid
	}

	return handsTotal, nil
}

func d7p2() (int, error) {
	file, err := os.Open("inputs/d7p2.txt")
	if err != nil {
		return 0, errors.New("could not read file input")
	}

	defer file.Close()
	reader := bufio.NewReader(file)

	hands := []CamelGameHand{}

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

		handStr := strings.Split(string(line), " ")
		if len(handStr) != 2 {
			continue
		}

		cards := strings.Split(handStr[0], "")
		bid, err := strconv.Atoi(handStr[1])
		if err != nil {
			continue
		}

		hand := CamelGameHand{
			cards: cards,
			bid:   bid,
		}
		hands = append(hands, hand)
	}

	sort.Slice(hands, func(i, j int) bool {
		a := hands[i]
		b := hands[j]

		return compareJokerCamelGameHand(a, b)
	})

	handsTotal := 0
	for i := 1; i <= len(hands); i++ {
		handsTotal += i * hands[i-1].bid
	}

	return handsTotal, nil
}

func compareCamelGameHand(a, b CamelGameHand) bool {
	aType := getHandType(a.cards)
	bType := getHandType(b.cards)

	if aType != bType {
		return aType < bType
	}

	// Tie break
	for i := 0; i < len(a.cards); i++ {
		aCardRank := cardRanks[a.cards[i]]
		bCardRank := cardRanks[b.cards[i]]

		if aCardRank != bCardRank {
			return aCardRank < bCardRank
		}
	}

	// Absolute fallback but shouldn't even hit here
	return a.bid < b.bid
}

func compareJokerCamelGameHand(a, b CamelGameHand) bool {
	aType := getHighestRankWithJoker(a.cards)
	bType := getHighestRankWithJoker(b.cards)

	if aType != bType {
		return aType < bType
	}

	// Tie break
	for i := 0; i < len(a.cards); i++ {
		aCardRank := jokerCardRanks[a.cards[i]]
		bCardRank := jokerCardRanks[b.cards[i]]

		if aCardRank != bCardRank {
			return aCardRank < bCardRank
		}
	}

	// Absolute fallback but shouldn't even hit here
	return a.bid < b.bid
}

// Five of a kind = 6, High card = 0
func getHandType(cards []string) int {
	cardCounts := map[string]int{}
	for _, card := range cards {
		cardCount, ok := cardCounts[card]
		if !ok {
			cardCounts[card] = 1
		} else {
			cardCounts[card] = cardCount + 1
		}
	}

	countsList := []int{}
	for _, count := range cardCounts {
		countsList = append(countsList, count)
	}

	// 5 of a kind
	if contains(countsList, 5) {
		return 6
	}

	// 4 of a kind
	if contains(countsList, 4) {
		return 5
	}

	// Full house
	if contains(countsList, 3) && contains(countsList, 2) {
		return 4
	}

	// 3 of a kind
	if contains(countsList, 3) {
		return 3
	}

	pairs := 0
	for _, count := range countsList {
		if count == 2 {
			pairs++
		}
	}

	// 2 pair
	if pairs == 2 {
		return 2
	}

	// 1 pair
	if pairs == 1 {
		return 1
	}

	// High card
	return 0
}

func getHighestRankWithJoker(cards []string) int {
	possibleHandRanks := []int{}

	// Try all possible
	for _, card := range cards {
		if card == "J" {
			continue
		}

		cardsCopy := []string{}
		cardsCopy = append(cardsCopy, cards...)
		for i := 0; i < len(cardsCopy); i++ {
			cardsCopy[i] = strings.Replace(cardsCopy[i], "J", card, -1)
		}

		possibleHandRanks = append(possibleHandRanks, getHandType(cardsCopy))
	}

	// Default to base hand
	highestRank := getHandType(cards)
	for _, rank := range possibleHandRanks {
		if rank > highestRank {
			highestRank = rank
		}
	}

	return highestRank
}
