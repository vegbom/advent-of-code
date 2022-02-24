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

type ScannerPair struct {
	s1 int
	s2 int
}

type Coord3D struct {
	x int
	y int
	z int
}

type Scanner2D struct {
	id         int
	beaconsAbs []Coord2D
	beaconsRel [][]Coord2D
}

func main() {
	list := Loader("2dtest.txt")
	fmt.Printf("%v\n", list)

	overlapBeacons := make(map[ScannerPair][]int, 0)
	for i := 0; i < (len(list) - 1); i++ {
		s1 := list[i]

		// Generate relative coordinates
		s1.beaconsRel = getRelativeCoords2D(s1.beaconsAbs)

		for j := i + 1; j < len(list); j++ {
			s2 := list[j]
			if i == j {
				continue
			}

			// Rotate through all the options
			for rotation := 0; rotation <= 3; rotation++ {
				// Generate relative coordinates
				noflip, flipped = getRelativeCoords2D(rotate2D(s2.beaconsAbs, rotation))

				// Count how many same for each of the beacons and if it exceeds threashold
				// If it does, then record the beacon ID?
				THREASHOLD := 2
				overlap := getOverlapBeaconIds(s1.beaconsRel, s2.beaconsRel, THREASHOLD)
				if len(overlap) > THREASHOLD {
					fmt.Printf("Beacon %d and %d in rotation %d has matching: %v\n", i, j, rotation, trimBeaconListForPrinting(s1.beaconsAbs, overlap))
					// found the matching orientation
					overlapBeacons[ScannerPair{i, j}] = overlap
					break
					// break out TODO CHECK IF THIS IS TRUE? It might not be?
				}
			}
		}
	}

	// try to match
	// n := rotate2D(list[0], 0, false)
	// fmt.Printf("%v\n", n)
	// n = rotate2D(n, 1, false)
	// fmt.Printf("%v\n", n)
	// n = rotate2D(n, 3, false)
	// fmt.Printf("%v\n", n)
	// n = rotate2D(n, 1, false)
	// n = rotate2D(n, 1, true)

	// fmt.Printf("Relatives: %v\n", getRelativeCoords2D(n.beacons))
	// fmt.Printf("Relatives: %v\n", getRelativeCoords2D(list[1].beacons))

}

func getOverlapBeaconIds(list1 [][]Coord2D, list2 [][]Coord2D, threashold int) (beaconIds []int) {
	beaconIds = make([]int, 0)
	for i, v1 := range list1 {
		for _, v2 := range list2 {
			if getNumOverlapLinks(v1, v2) >= threashold {
				beaconIds = append(beaconIds, i)
			}
		}
	}
	return beaconIds
}

func getNumOverlapLinks(beacon1 []Coord2D, beacon2 []Coord2D) (count int) {
	for _, link1 := range beacon1 {
		for _, link2 := range beacon2 {
			if link1 == link2 {
				count++
			}
		}
	}
	return count
}

func trimBeaconListForPrinting(beaconsAbs []Coord2D, beaconIds []int) []Coord2D {
	result := make([]Coord2D, 0)
	for _, id := range beaconIds {
		result = append(result, beaconsAbs[id])
	}
	return result
}

func getRelativeCoords2D(beaconsAbs []Coord2D) [][]Coord2D {
	relCoords := make([][]Coord2D, 0)
	for _, b1 := range beaconsAbs {
		currentBeacon := make([]Coord2D, 0)
		for _, b2 := range beaconsAbs {
			currentBeacon = append(currentBeacon, Coord2D{b2.x - b1.x, b2.y - b1.y})
		}
		relCoords = append(relCoords, currentBeacon)
	}

	return relCoords
}

func getRelativeCoords(beacons []Coord3D) [][]Coord3D {
	relCoords := make([][]Coord3D, 0)
	for _, b1 := range beacons {
		currentBeacon := make([]Coord3D, 0)
		for _, b2 := range beacons {
			currentBeacon = append(currentBeacon, Coord3D{b2.x - b1.x, b2.y - b1.y, b2.z - b1.z})
		}
		relCoords = append(relCoords, currentBeacon)
	}

	return relCoords
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
			currentScanner.id = MustStringToInt(matches[headerExp.SubexpIndex("id")])
			currentScanner.beaconsAbs = make([]Coord2D, 0)
			scanHeader = false
		} else {
			coordsString := strings.Split(goScanner.Text(), ",")
			currentScanner.beaconsAbs = append(currentScanner.beaconsAbs, Coord2D{MustStringToInt(coordsString[0]), MustStringToInt(coordsString[1])})
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
