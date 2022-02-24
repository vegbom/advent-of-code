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

type Coord2D struct {
	x int
	y int
}

type Scanner2D struct {
	id      int
	beacons []Coord2D
}

func main() {
	list := Loader("2dtest.txt")
	fmt.Printf("%v", list)
}

func Loader(filename string) []Scanner2D {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	goScanner := bufio.NewScanner(f)

	result := make([]Scanner2D, 0)
	var currentScanner Scanner2D

	scanHeader := true
	headerExp := regexp.MustCompile(`scanner (?P<id>\d+)`)
	// beaconsExp2D := regexp.MustCompile(`^(?P<x>-?[0-9]\d+),(?P<y>-?[0-9]\d+)`)
	// beaconsExp3D := regexp.MustCompile(`^(?P<x>-?[0-9]\d+),(?P<y>-?[0-9]\d+),(?P<z>-?[0-9]\d+)`)

	for goScanner.Scan() {
		if strings.TrimSpace(goScanner.Text()) == "" {
			result = append(result, currentScanner)
			scanHeader = true
			continue
		}

		if scanHeader {
			fmt.Printf("INSTR: %v\n", goScanner.Text())
			matches := headerExp.FindStringSubmatch(goScanner.Text())
			// instrs = append(instrs, Instruction{matches[r.SubexpIndex("dir")], MustStringToInt(matches[r.SubexpIndex("pos")])})
			currentScanner.id = MustStringToInt(matches[headerExp.SubexpIndex("id")])
			currentScanner.beacons = make([]Coord2D, 0)
			scanHeader = false
		} else {
			coordsString := strings.Split(goScanner.Text(), ",")
			currentScanner.beacons = append(currentScanner.beacons, Coord2D{MustStringToInt(coordsString[0]), MustStringToInt(coordsString[1])})
		}
	}
	result = append(result, currentScanner)

	return result
}

func MustStringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}
