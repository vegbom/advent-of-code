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

type Scanner3D struct {
	id         int
	beaconsAbs []Coord3D
	beaconsRel [][]Coord3D
}

func main() {
	list := Loader("3dtest1.txt")
	fmt.Printf("INPUT:\n%v\n", list)

	for i, orientation := range allOrientations3D() {
		// Generate relative coordinates
		a := rotate3D(list[0].beaconsAbs, orientation)
		fmt.Printf("Orientation %d - %v:\n", i, orientation)
		for _, v := range a {
			fmt.Printf("%d,%d,%d\n", v.x, v.y, v.z)
		}
		fmt.Printf("\n")
	}

	return

	// overlapBeacons := make(map[ScannerPair][]int, 0)
	// for i := 0; i < (len(list) - 1); i++ {
	// 	s1 := list[i]

	// 	// Generate relative coordinates
	// 	s1.beaconsRel = getRelativeCoords2D(s1.beaconsAbs)

	// 	for j := i + 1; j < len(list); j++ {
	// 		s2 := list[j]
	// 		if i == j {
	// 			continue
	// 		}

	// 		// Rotate through all the options
	// 		for _, orientation := range allOrientations2D() {
	// 			// Generate relative coordinates
	// 			s2.beaconsRel = getRelativeCoords2D(rotate2D(s2.beaconsAbs, orientation))
	// 			// Count how many same for each of the beacons and if it exceeds threshold
	// 			// If it does, then record the beacon ID?
	// 			THRESHOLD := 3
	// 			overlap := getOverlapBeaconIds(s1.beaconsRel, s2.beaconsRel, THRESHOLD)
	// 			if len(overlap) >= THRESHOLD {
	// 				fmt.Printf("Beacon %d and %d in orientation %v has %d matching\n", i, j, orientation, len(overlap))
	// 				// found the matching orientation
	// 				overlapBeacons[ScannerPair{i, j}] = overlap
	// 				// break
	// 				// break out TODO CHECK IF THIS IS TRUE? It might not be?
	// 			}
	// 		}
	// 	}
	// }

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

func getOverlapBeaconIds(list1 [][]Coord2D, list2 [][]Coord2D, threshold int) (beaconIds []int) {
	beaconIds = make([]int, 0)
	for i, v1 := range list1 {
		for _, v2 := range list2 {
			if getNumOverlapLinks(v1, v2) >= (threshold - 1) {
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
			if b1 == b2 {
				continue
			}
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
			if b1 == b2 {
				continue
			}
			currentBeacon = append(currentBeacon, Coord3D{b2.x - b1.x, b2.y - b1.y, b2.z - b1.z})
		}
		relCoords = append(relCoords, currentBeacon)
	}

	return relCoords
}

func Loader(filename string) []Scanner3D {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	goScanner := bufio.NewScanner(f)

	result := make([]Scanner3D, 0)
	var currentScanner Scanner3D

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
			currentScanner.beaconsAbs = make([]Coord3D, 0)
			scanHeader = false
		} else {
			coordsString := strings.Split(goScanner.Text(), ",")
			currentScanner.beaconsAbs = append(currentScanner.beaconsAbs, Coord3D{MustStringToInt(coordsString[0]), MustStringToInt(coordsString[1]), MustStringToInt(coordsString[2])})
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
