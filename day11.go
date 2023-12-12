package main

import (
	"bufio"
	"errors"
	"io"
	"math"
	"os"
	"strings"
)

func day11() (int, int, error) {
	answerP1, err := d11p1()
	if err != nil {
		return 0, 0, err
	}

	answerP2, err := d11p2()
	if err != nil {
		return 0, 0, err
	}

	return answerP1, answerP2, err
}

func d11p1() (int, error) {
	file, err := os.Open("inputs/d11p1.txt")
	if err != nil {
		return 0, errors.New("could not read file input")
	}

	defer file.Close()
	reader := bufio.NewReader(file)

	galaxyMap := [][]string{}

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

		galaxyLine := strings.Split(string(line), "")
		galaxyMap = append(galaxyMap, galaxyLine)
	}

	galaxies := []Coord{}
	for y := range galaxyMap {
		for x := range galaxyMap[y] {
			if galaxyMap[y][x] == "#" {
				galaxies = append(galaxies, Coord{x: x, y: y})
			}
		}
	}

	galaxies = scaleGalaxy(2, galaxyMap, galaxies)

	total := 0
	for i, galaxy1 := range galaxies {
		for j, galaxy2 := range galaxies {
			if galaxy1 == galaxy2 || i > j {
				continue
			}

			horizontal := math.Abs(float64(galaxy1.x - galaxy2.x))
			vertical := math.Abs(float64(galaxy1.y - galaxy2.y))
			total += int(horizontal) + int(vertical)
		}
	}

	return total, nil
}

func d11p2() (int, error) {
	file, err := os.Open("inputs/d11p2.txt")
	if err != nil {
		return 0, errors.New("could not read file input")
	}

	defer file.Close()
	reader := bufio.NewReader(file)

	galaxyMap := [][]string{}

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

		galaxyLine := strings.Split(string(line), "")
		galaxyMap = append(galaxyMap, galaxyLine)
	}

	galaxies := []Coord{}
	for y := range galaxyMap {
		for x := range galaxyMap[y] {
			if galaxyMap[y][x] == "#" {
				galaxies = append(galaxies, Coord{x: x, y: y})
			}
		}
	}

	galaxies = scaleGalaxy(1000000, galaxyMap, galaxies)

	total := 0
	for i, galaxy1 := range galaxies {
		for j, galaxy2 := range galaxies {
			if galaxy1 == galaxy2 || i > j {
				continue
			}

			horizontal := math.Abs(float64(galaxy1.x - galaxy2.x))
			vertical := math.Abs(float64(galaxy1.y - galaxy2.y))
			total += int(horizontal) + int(vertical)
		}
	}

	return total, nil
}

func scaleGalaxy(scale int, galaxyMap [][]string, galaxies []Coord) []Coord {
	expandedGalaxies := make([]Coord, len(galaxies))
	copy(expandedGalaxies, galaxies)

	// Expand rows
	for y := 0; y < len(galaxyMap); y++ {
		empty := true

		for _, character := range galaxyMap[y] {
			if character == "#" {
				empty = false
				break
			}
		}

		if !empty {
			continue
		}

		for i := 0; i < len(galaxies); i++ {
			if galaxies[i].y > y {
				expandedGalaxies[i] = Coord{
					x: expandedGalaxies[i].x,
					y: expandedGalaxies[i].y + (scale - 1),
				}
			}
		}
	}

	// Expand cols

	for x := 0; x < len(galaxyMap[0]); x++ {
		empty := true

		for y := 0; y < len(galaxyMap); y++ {
			if galaxyMap[y][x] == "#" {
				empty = false
				break
			}
		}

		if !empty {
			continue
		}

		for i := 0; i < len(galaxies); i++ {
			if galaxies[i].x > x {
				expandedGalaxies[i] = Coord{
					x: expandedGalaxies[i].x + (scale - 1),
					y: expandedGalaxies[i].y,
				}
			}
		}
	}

	return expandedGalaxies
}
