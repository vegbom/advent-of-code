package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Coordinate struct {
	x int
	y int
}

type Instruction struct {
	dir string
	pos int
}

func main() {
	dots, instructions, err := Loader("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	_ = instructions

	// Part 1:
	// Visualize(dots)
	d1 := Fold(dots, instructions[0])
	// Visualize(d1)
	fmt.Printf("Size of d1 (visible dots):%d\n", len(d1)) // Part 1 answer: 847

	// Part 2:
	for _, instr := range instructions {
		dots = Fold(dots, instr)
	}
	Visualize(dots)
	fmt.Printf("Size of dots (visible dots):%d\n", len(dots))
}

func Fold(dots []Coordinate, inst Instruction) []Coordinate {
	// xSz, ySz := GetGridSize(dots)

	fmt.Printf("Folding with %v\n", inst)

	dotsNew := make([]Coordinate, 0)

	for _, dot := range dots {
		// Need to fold
		foldFrom := 0
		if inst.dir == "y" {
			foldFrom = dot.y
		} else {
			foldFrom = dot.x
		}
		distToFold := foldFrom - inst.pos

		if distToFold < 0 {
			// Unaffected. Copy directly.
		} else if distToFold == 0 {
			// on the fold line, do not append
			continue
		} else {
			// folding needed
			foldTo := inst.pos - distToFold
			if inst.dir == "y" {
				dot.y = foldTo
			} else {
				dot.x = foldTo
			}
		}

		if !duplicateExists(dot, dotsNew) {
			dotsNew = append(dotsNew, dot)
		}
	}

	return dotsNew
}

func duplicateExists(dot Coordinate, dots []Coordinate) bool {
	for _, v := range dots {
		if v.x == dot.x && v.y == dot.y {
			return true
		}
	}
	return false
}

func GetGridSize(dots []Coordinate) (xSz int, ySz int) {
	for _, v := range dots {
		if v.x > xSz {
			xSz = v.x
		}
		if v.y > ySz {
			ySz = v.y
		}
	}
	return xSz + 1, ySz + 1
}

func Visualize(dots []Coordinate) {
	xSz, ySz := GetGridSize(dots)

	for y := 0; y < ySz; y++ {
		for x := 0; x < xSz; x++ {
			found := false
			for _, v := range dots {
				if v.x == x && v.y == y {
					fmt.Print("#")
					found = true
					break
				}
			}
			if !found {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
}

func Loader(filename string) (dots []Coordinate, instrs []Instruction, err error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()

	dots = make([]Coordinate, 0)
	scanner := bufio.NewScanner(f)

	instrs = make([]Instruction, 0)

	scanInstructions := false
	r := regexp.MustCompile(`fold along (?P<dir>[xy])=(?P<pos>[[:digit:]]+)`)

	for scanner.Scan() {
		if strings.TrimSpace(scanner.Text()) == "" {
			scanInstructions = true
			continue
		}

		if !scanInstructions {
			coordsString := strings.Split(scanner.Text(), ",")
			dots = append(dots, Coordinate{MustStringToInt(coordsString[0]), MustStringToInt(coordsString[1])})
		} else {
			fmt.Printf("INSTR: %v\n", scanner.Text())
			matches := r.FindStringSubmatch(scanner.Text())
			instrs = append(instrs, Instruction{matches[r.SubexpIndex("dir")], MustStringToInt(matches[r.SubexpIndex("pos")])})
		}

	}

	return dots, instrs, err
}

func MustStringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}
