package main

// Learning objectives: Regex

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type VentLine struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

func swap(a, b int) (int, int) {
	temp := b
	b = a
	a = temp
	return a, b
}

func isNegative(a int) int {
	if a < 0 {
		return -1
	}
	return 1
}

func main() {
	lines, maxX, maxY, err := Loader("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Part 1
	computePoints(true, lines, maxX, maxY) // overlapCount: 5373
	// Part 2
	computePoints(false, lines, maxX, maxY) // overlapCount: 21514
}

func computePoints(isPart1 bool, lines []VentLine, maxX int, maxY int) int {
	// Initialize the 2D grid
	fmt.Printf("Max X: %d, Max Y: %d\n", maxX, maxY)
	// +1 because it's size, not 0-indexed ID
	diagram := make([][]int, maxY+1)
	for i := range diagram {
		diagram[i] = make([]int, maxX+1)
	}

	for _, line := range lines {
		dx := line.x2 - line.x1
		dy := line.y2 - line.y1
		x := line.x1
		y := line.y1

		if dx == 0 {
			// Vertical line
			for y != line.y2+isNegative(dy) {
				diagram[y][x]++
				y += isNegative(dy)
			}
		} else {
			// PART 1 SKIP:
			if dy != 0 && isPart1 {
				continue
			}
			m := dy / dx
			b := line.y1 - m*line.x1
			for x <= line.x2 {
				y = m*x + b
				diagram[y][x]++
				x++
			}
		}
	}

	return completeDiagram(diagram)
}

func completeDiagram(d [][]int) int {
	// loop through diagram to count and optionally print
	// the items
	overlapCount := 0
	for _, vx := range d {
		for _, v := range vx {
			// if v == 0 {
			// 	fmt.Printf(".")
			// } else {
			// 	fmt.Printf("%v", v)
			// }
			if v >= 2 {
				overlapCount++
			}
		}
		// fmt.Printf("\n")
	}
	fmt.Printf("overlapCount: %d\n", overlapCount)
	return overlapCount
}

func MustStringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func Loader(filename string) (lines []VentLine, maxX int, maxY int, err error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, 0, 0, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	lines = make([]VentLine, 0)
	r := regexp.MustCompile(`(?P<x1>\d+),(?P<y1>\d+)\s->\s(?P<x2>\d+),(?P<y2>\d+)`)

	for scanner.Scan() {
		// Could have used reflection here but it's too advanced of a topic for now
		// fmt.Printf("%#v\n", r.FindStringSubmatch(scanner.Text()))
		// fmt.Printf("%#v\n", r.SubexpNames())
		matches := r.FindStringSubmatch(scanner.Text())
		line := VentLine{
			MustStringToInt(matches[1]),
			MustStringToInt(matches[2]),
			MustStringToInt(matches[3]),
			MustStringToInt(matches[4]),
		}

		// Enforce x2 >= x1
		if line.x1 > line.x2 {
			line.x1, line.x2 = swap(line.x1, line.x2)
			line.y1, line.y2 = swap(line.y1, line.y2)
		}

		lines = append(lines, line)

		// Calculate the max size
		isLarger := func(val, max int) int {
			if val > max {
				return val
			}
			return max
		}
		maxX = isLarger(line.x2, maxX)
		maxY = isLarger(line.y1, maxY)
		maxY = isLarger(line.x2, maxY)
	}

	return lines, maxX, maxY, nil
}
