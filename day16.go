package main

import (
	"bufio"
	"errors"
	"io"
	"os"
)

func day16() (int, int, error) {
	answerP1, err := d16p1()
	if err != nil {
		return 0, 0, err
	}

	answerP2, err := d16p2()
	if err != nil {
		return 0, 0, err
	}

	return answerP1, answerP2, err
}

func d16p1() (int, error) {
	file, err := os.Open("inputs/d16p1.txt")
	if err != nil {
		return 0, errors.New("could not read file input")
	}

	defer file.Close()
	reader := bufio.NewReader(file)

	grid := []string{}

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

		grid = append(grid, string(line))
	}

	return countEnergisedTiles(grid, Coord{-1, 0}, Coord{0, 0}), nil
}

func d16p2() (int, error) {
	file, err := os.Open("inputs/d16p2.txt")
	if err != nil {
		return 0, errors.New("could not read file input")
	}

	defer file.Close()
	reader := bufio.NewReader(file)

	grid := []string{}

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

		grid = append(grid, string(line))
	}

	energisedTiles := []int{}
	for y := range grid {
		energisedTiles = append(energisedTiles, countEnergisedTiles(grid, Coord{-1, y}, Coord{0, y}))
		energisedTiles = append(energisedTiles, countEnergisedTiles(grid, Coord{len(grid[0]), y}, Coord{len(grid[0]) - 1, y}))
	}
	for x := range grid[0] {
		energisedTiles = append(energisedTiles, countEnergisedTiles(grid, Coord{x, -1}, Coord{x, 0}))
		energisedTiles = append(energisedTiles, countEnergisedTiles(grid, Coord{x, len(grid)}, Coord{x, len(grid) - 1}))
	}

	maxTiles := 0
	for _, tileCount := range energisedTiles {
		if tileCount > maxTiles {
			maxTiles = tileCount
		}
	}

	return maxTiles, nil
}

func countEnergisedTiles(grid []string, start, next Coord) int {
	cache := make(map[struct{ prev, current Coord }][]Coord) // Speed up by reusing between runs
	energisedTiles := traverseLightGrid(grid, start, next, []struct{ prev, current Coord }{}, cache)

	uniqueTiles := make(map[Coord]bool)
	for _, tile := range energisedTiles {
		uniqueTiles[tile] = true
	}

	return len(uniqueTiles)
}

func traverseLightGrid(grid []string, prev, current Coord, visited []struct{ prev, current Coord }, cache map[struct{ prev, current Coord }][]Coord) []Coord {
	cacheKey := struct{ prev, current Coord }{prev, current}
	tiles, ok := cache[cacheKey]
	if ok {
		return tiles
	}

	if tileVisited(prev, current, visited) || tileOutOfBounds(grid, current) {
		return []Coord{}
	}

	visited = append(visited, struct{ prev, current Coord }{prev, current})
	tiles = []Coord{current}
	char := grid[current.y][current.x]

	// Vertical split
	if char == '|' && prev.x != current.x {
		tiles = append(tiles, traverseLightGrid(grid, current, Coord{current.x, current.y - 1}, visited, cache)...)
		tiles = append(tiles, traverseLightGrid(grid, current, Coord{current.x, current.y + 1}, visited, cache)...)

		cacheTiles := make([]Coord, len(tiles))
		copy(cacheTiles, tiles)
		cache[cacheKey] = tiles

		return tiles
	}

	// Horizontal split
	if char == '-' && prev.y != current.y {
		tiles = append(tiles, traverseLightGrid(grid, current, Coord{current.x - 1, current.y}, visited, cache)...)
		tiles = append(tiles, traverseLightGrid(grid, current, Coord{current.x + 1, current.y}, visited, cache)...)

		cacheTiles := make([]Coord, len(tiles))
		copy(cacheTiles, tiles)
		cache[cacheKey] = tiles

		return tiles
	}

	// Right angle /
	if char == '/' {
		newCoord := Coord{
			x: current.x - (current.y - prev.y),
			y: current.y - (current.x - prev.x),
		}
		tiles = append(tiles, traverseLightGrid(grid, current, newCoord, visited, cache)...)

		cacheTiles := make([]Coord, len(tiles))
		copy(cacheTiles, tiles)
		cache[cacheKey] = tiles

		return tiles
	}

	// Left angle \
	if char == '\\' {
		newCoord := Coord{
			x: current.x + (current.y - prev.y),
			y: current.y + (current.x - prev.x),
		}
		tiles = append(tiles, traverseLightGrid(grid, current, newCoord, visited, cache)...)

		cacheTiles := make([]Coord, len(tiles))
		copy(cacheTiles, tiles)
		cache[cacheKey] = tiles

		return tiles
	}

	// Pass through
	newCoord := Coord{
		x: current.x + current.x - prev.x,
		y: current.y + current.y - prev.y,
	}

	tiles = append(tiles, traverseLightGrid(grid, current, newCoord, visited, cache)...)

	cacheTiles := make([]Coord, len(tiles))
	copy(cacheTiles, tiles)
	cache[cacheKey] = tiles

	return tiles
}

func tileVisited(prev, current Coord, visited []struct{ prev, current Coord }) bool {
	for _, tile := range visited {
		if tile.prev == prev && tile.current == current {
			return true
		}
	}

	return false
}

func tileOutOfBounds(grid []string, coord Coord) bool {
	if coord.y < 0 || coord.y >= len(grid) {
		return true
	}

	if coord.x < 0 || coord.x >= len(grid[0]) {
		return true
	}

	return false
}
