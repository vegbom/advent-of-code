package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

const Reset = "\033[0m"
const Red = "\033[31m"
const Green = "\033[32m"
const Yellow = "\033[33m"

type Coordinate struct {
	row int
	col int
}

func main() {
	heightmap, err := Loader("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	mapSz := Coordinate{len(heightmap), len(heightmap[0])}

	visited := make([][]bool, mapSz.row)
	for i := range visited {
		visited[i] = make([]bool, mapSz.col)
	}

	// fmt.Printf("%v\n", heightmap)
	// fmt.Printf("%v\n", visited)

	// Altitude is from 0 to 9
	basinSizes := make([]int, 0)
	minAlts := make([]int, 0)
	minAltCoords := make([]Coordinate, 0)
	for altitude := 0; altitude <= 9; altitude++ {
		// find a point which is this altitude
		for i := range heightmap {
			for j, val := range heightmap[i] {
				if val == altitude && !visited[i][j] {
					bsz := 0
					visited, bsz = visitAdjacentAndMark(Coordinate{i, j}, mapSz, heightmap, visited, bsz)
					minAlts = append(minAlts, altitude)
					minAltCoords = append(minAltCoords, Coordinate{i, j})
					basinSizes = append(basinSizes, bsz)
				}
			}
		}
	}

	fmt.Printf("minAlts=%v\n", minAlts)
	fmt.Printf("minAltCoords=%v\n", minAltCoords)

	riskLevelSum := 0
	for _, val := range minAlts {
		riskLevelSum += val + 1
	}
	fmt.Printf("riskLevelSum=%v\n", riskLevelSum)

	sort.Ints(basinSizes)
	// assuming we have at least three basins!
	fmt.Printf("basinSizes=%v\n", basinSizes)
	bszMultiply := basinSizes[len(basinSizes)-1] * basinSizes[len(basinSizes)-2] * basinSizes[len(basinSizes)-3]
	fmt.Printf("basinSize SUM: %v\n", bszMultiply)
}

func visitAdjacentAndMark(current Coordinate, mapSz Coordinate, heightmap [][]int, visited [][]bool, basinSize int) ([][]bool, int) {
	if visited[current.row][current.col] {
		return visited, basinSize
	}
	visited[current.row][current.col] = true
	if heightmap[current.row][current.col] != 9 {
		basinSize++
	}
	for _, adj := range getAdjacentCoords(current, mapSz) {
		if heightmap[current.row][current.col] <= heightmap[adj.row][adj.col] {
			visited, basinSize = visitAdjacentAndMark(adj, mapSz, heightmap, visited, basinSize)
		}
	}
	return visited, basinSize
}

func getAdjacentCoords(current Coordinate, mapSz Coordinate) []Coordinate {
	adjacents := make([]Coordinate, 0)

	if current.row+1 < mapSz.row {
		adjacents = append(adjacents, Coordinate{current.row + 1, current.col})
	}
	if current.row-1 >= 0 {
		adjacents = append(adjacents, Coordinate{current.row - 1, current.col})
	}
	if current.col+1 < mapSz.col {
		adjacents = append(adjacents, Coordinate{current.row, current.col + 1})
	}
	if current.col-1 >= 0 {
		adjacents = append(adjacents, Coordinate{current.row, current.col - 1})
	}

	return adjacents
}

func Loader(filename string) (heightmap [][]int, err error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	heightmap = make([][]int, 0)

	for scanner.Scan() {
		row_data := make([]int, 0)
		for _, positionStr := range scanner.Text() {
			i, err := strconv.Atoi(string(positionStr))
			if err != nil {
				return nil, err
			}
			row_data = append(row_data, i)
		}

		heightmap = append(heightmap, row_data)
	}

	return heightmap, err
}
