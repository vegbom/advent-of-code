package main

import (
	"bufio"
	"day12/graph"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

var idStart int
var idEnd int
var isPart1 bool

func main() {
	caves, err := Loader("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	idStart, _ = caves.GetId("start")
	idEnd, _ = caves.GetId("end")

	initVisit := make([]int, 0)

	isPart1 = true
	routesPart1 := Explore(initVisit, idStart, caves)
	fmt.Printf("PART 1 Total Number of Routes: %d\n", len(routesPart1))

	isPart1 = false
	routesPart2 := Explore(initVisit, idStart, caves)
	fmt.Printf("PART 2 Total Number of Routes: %d\n", len(routesPart2))

	// for _, route := range routesPart1 {
	// 	for _, n := range route {
	// 		fmt.Printf("%s,", caves.GetName(n))
	// 	}
	// 	fmt.Printf("\n")
	// }
}

func Explore(visited []int, current int, caves *graph.Graph) [][]int {

	if current == idStart && len(visited) > 0 {
		return nil
	}

	// PART 2: Enforce a single small cave can be visited at most twice,
	// the remaining small caves can be visited at most once
	currentVisitCount := 0
	if caves.IsSmall(current) {
		for _, n := range visited {
			if n == current {
				currentVisitCount++
			}
		}
	}
	if currentVisitCount == 1 {
		if isPart1 {
			// PART 1: Enforce small cave only visit once
			return nil
		}
		// This is second visit
		smallCavesVisited := make([]int, 0)
		for _, v := range visited {
			if caves.IsSmall(v) && v != current {
				smallCavesVisited = append(smallCavesVisited, v)
			}
		}
		sort.Ints(smallCavesVisited)
		for i := 1; i < len(smallCavesVisited); i++ {
			if smallCavesVisited[i-1] == smallCavesVisited[i] {
				return nil
			}
		}
	} else if currentVisitCount > 1 {
		// the second visit already occured
		return nil
	}

	newRoutes := make([][]int, 0)
	visited = append(visited, current)

	if current == idEnd {
		return append(newRoutes, visited)
	}

	for _, next := range caves.Neighbours(current) {
		visitNew := make([]int, len(visited))
		copy(visitNew, visited)
		result := Explore(visitNew, next, caves)
		if len(result) > 0 {
			newRoutes = append(newRoutes, result...)
		}
	}
	return newRoutes
}

func Loader(filename string) (caves *graph.Graph, err error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	caves = graph.New()
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		names := strings.Split(scanner.Text(), "-")
		id1, err1 := caves.GetId(names[0])
		if err1 != nil {
			id1 = caves.AddNode(names[0])
		}
		id2, err2 := caves.GetId(names[1])
		if err2 != nil {
			id2 = caves.AddNode(names[1])
		}
		caves.AddEdge(id1, id2)
	}

	return caves, err
}
