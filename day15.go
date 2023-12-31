package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type HashLens struct {
	label  string
	hash   int
	dash   bool
	length int
}

func day15() (int, int, error) {
	answerP1, err := d15p1()
	if err != nil {
		return 0, 0, err
	}

	answerP2, err := d15p2()
	if err != nil {
		return 0, 0, err
	}

	return answerP1, answerP2, err
}

func d15p1() (int, error) {
	file, err := os.Open("inputs/d15p1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	hashes := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		for _, step := range strings.Split(scanner.Text(), ",") {
			hashes = append(hashes, hashInitSequenceStep(step))
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return sum(hashes), nil
}

func d15p2() (int, error) {
	file, err := os.Open("inputs/d15p2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	boxes := make([][]HashLens, 256)
	lenses := []HashLens{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		for _, step := range strings.Split(scanner.Text(), ",") {
			dashSplit := strings.Split(step, "-")
			equalSplit := strings.Split(step, "=")

			lens := HashLens{}
			if len(dashSplit) == 2 {
				lens.label = dashSplit[0]
				lens.hash = hashInitSequenceStep(dashSplit[0])
				lens.dash = true

				focalLength, err := strconv.Atoi(dashSplit[1])
				if err == nil {
					lens.length = focalLength
				}
			} else if len(equalSplit) > 1 {
				lens.label = equalSplit[0]
				lens.hash = hashInitSequenceStep(equalSplit[0])
				lens.dash = false

				focalLength, err := strconv.Atoi(equalSplit[1])
				if err == nil {
					lens.length = focalLength
				}
			}
			lenses = append(lenses, lens)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for _, lens := range lenses {
		foundLensIdx := -1

		for i, boxLens := range boxes[lens.hash] {
			if boxLens.label == lens.label {
				foundLensIdx = i
				break
			}
		}

		// Add new lens
		if !lens.dash && foundLensIdx == -1 {
			boxes[lens.hash] = append(boxes[lens.hash], lens)
			continue
		}

		// Update lens length
		if !lens.dash && foundLensIdx != -1 {
			boxes[lens.hash][foundLensIdx] = lens
			continue
		}

		// Lens doesn't exist
		if lens.dash && foundLensIdx == -1 {
			continue
		}

		// Remove the lens
		boxes[lens.hash] = append(boxes[lens.hash][:foundLensIdx], boxes[lens.hash][foundLensIdx+1:]...)
	}

	total := 0
	for b, box := range boxes {
		for l, lens := range box {
			focus := b + 1
			focus *= l + 1
			focus *= lens.length
			total += focus
		}
	}

	return total, nil
}

func hashInitSequenceStep(step string) int {
	currentValue := 0
	for _, character := range step {
		currentValue += int(character)
		currentValue *= 17
		currentValue %= 256
	}

	return currentValue
}
