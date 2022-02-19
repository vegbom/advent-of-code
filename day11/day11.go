package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
	octopusmap, err := Loader("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	octopusmapSz := Coordinate{len(octopusmap), len(octopusmap[0])}

	if octopusmapSz.row != 10 || octopusmapSz.col != 10 {
		log.Fatal("Map size incorrect")
	}

	Part1(octopusmap, octopusmapSz)
	Part2(octopusmap, octopusmapSz)
}

func Part1(input [][]int, size Coordinate) int {
	octopusmap := copyOctopusmap(input, size)
	flashSum := 0
	verbose := true
	for step := 1; step <= 100; step++ {
		if verbose {
			fmt.Printf("STEP: %d\n", step)
		}
		energyOut, flashedOut := Simulate(octopusmap, size)
		flashSum += CountNumFlashes(flashedOut)
		if verbose {
			Visualize(energyOut, flashedOut)
		}
		octopusmap = energyOut
	}

	fmt.Printf("\nPART 1: flashSum: %d\n", flashSum)
	return flashSum
}

func Part2(octopusmap [][]int, size Coordinate) int {
	map1 := copyOctopusmap(octopusmap, size)
	step := 0
	for {
		step++
		energyOut, flashedOut := Simulate(map1, size)
		if CountNumFlashes(flashedOut) == size.row*size.col {
			fmt.Printf("PART 2: All Octopuses Flash at Step %d\n", step)
			// Visualize(energyOut, flashedOut)
			return step
		}
		map1 = energyOut
	}
}

func copyOctopusmap(octopusmap [][]int, size Coordinate) [][]int {
	newMap := make([][]int, size.row)
	for i := range newMap {
		newMap[i] = make([]int, size.col)
		copy(newMap[i], octopusmap[i])
	}
	return newMap
}

func Visualize(energy [][]int, flashed [][]bool) {
	for i := range energy {
		lineOutput := ""
		for j := range energy[i] {
			if flashed[i][j] {
				lineOutput = fmt.Sprintf("%s%s%v%s", lineOutput, Green, energy[i][j], Reset)
			} else {
				lineOutput = fmt.Sprintf("%s%v", lineOutput, energy[i][j])
			}
		}
		fmt.Printf("%s\n", lineOutput)
	}
}

func CountNumFlashes(flashed [][]bool) (numFlashes int) {
	for i := range flashed {
		for j := range flashed[i] {
			if flashed[i][j] {
				numFlashes++
			}
		}
	}
	return numFlashes
}

func Simulate(energy [][]int, energySz Coordinate) ([][]int, [][]bool) {
	// Pass 1 - Increase every octopus by 1
	for i := range energy {
		for j := range energy[i] {
			energy[i][j]++
		}
	}

	// Pass 2 -
	// While there is an octopus level greather than 9 and has not been flashed, flash it
	flashed := make([][]bool, energySz.row)
	for i := range flashed {
		flashed[i] = make([]bool, energySz.col)
	}

	for i := range energy {
		for j := range energy[i] {
			energy, flashed = flash(Coordinate{i, j}, energySz, energy, flashed)
		}
	}

	return energy, flashed
}

func flash(current Coordinate, energySz Coordinate, energy [][]int, flashed [][]bool) ([][]int, [][]bool) {
	// Enforce state transitions
	if flashed[current.row][current.col] {
		energy[current.row][current.col] = 0
		return energy, flashed
	}
	if energy[current.row][current.col] <= 9 {
		return energy, flashed
	}

	flashed[current.row][current.col] = true
	energy[current.row][current.col] = 0

	for _, adj := range GetAdjacentCoords(current, energySz) {
		energy[adj.row][adj.col]++
		energy, flashed = flash(adj, energySz, energy, flashed)
	}
	return energy, flashed
}

func GetAdjacentCoords(current Coordinate, mapSz Coordinate) []Coordinate {
	candidates := make([]Coordinate, 0)
	candidates = append(candidates, Coordinate{current.row + 1, current.col - 1})
	candidates = append(candidates, Coordinate{current.row + 1, current.col})
	candidates = append(candidates, Coordinate{current.row + 1, current.col + 1})
	candidates = append(candidates, Coordinate{current.row - 1, current.col - 1})
	candidates = append(candidates, Coordinate{current.row - 1, current.col})
	candidates = append(candidates, Coordinate{current.row - 1, current.col + 1})
	candidates = append(candidates, Coordinate{current.row, current.col + 1})
	candidates = append(candidates, Coordinate{current.row, current.col - 1})

	adjacents := make([]Coordinate, 0)
	for _, v := range candidates {
		if v.row >= 0 && v.row < mapSz.row && v.col >= 0 && v.col < mapSz.col {
			adjacents = append(adjacents, v)
		}
	}

	return adjacents
}

func Loader(filename string) (octopusmap [][]int, err error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	octopusmap = make([][]int, 0)

	for scanner.Scan() {
		row_data := make([]int, 0)
		for _, positionStr := range scanner.Text() {
			i, err := strconv.Atoi(string(positionStr))
			if err != nil {
				return nil, err
			}
			row_data = append(row_data, i)
		}

		octopusmap = append(octopusmap, row_data)
	}

	return octopusmap, err
}
