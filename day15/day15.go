package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/RyanCarrier/dijkstra"
)

type Coordinate struct {
	row int
	col int
}

type RiskMap [][]int
type BigMap [][]RiskMap

func main() {
	riskmap := Loader("input.txt")
	Part1(riskmap)
	Part2(riskmap)
}

func Part1(riskmap [][]int) {
	mapSz := Coordinate{len(riskmap), len(riskmap[0])}
	graph := dijkstra.NewGraph()
	vertexId := 0

	for i := range riskmap {
		for j := range riskmap[i] {
			graph.AddVertex(vertexId)

			if j > 0 {
				graph.AddArc(vertexId-1, vertexId, int64(riskmap[i][j]))
				graph.AddArc(vertexId, vertexId-1, int64(riskmap[i][j-1]))
			}
			if i > 0 {
				graph.AddArc(vertexId-mapSz.row, vertexId, int64(riskmap[i][j]))
				graph.AddArc(vertexId, vertexId-mapSz.row, int64(riskmap[i-1][j]))
			}
			// fmt.Printf("%d ", riskmap[i][j])
			// fmt.Printf("[%d ↓ %d -> %d]", vertexId-mapSz.row, vertexId-1, vertexId)
			vertexId++
		}
		// fmt.Printf("\n")
	}

	best, err := graph.Shortest(0, vertexId-1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Part 1: Shortest distance ", best.Distance, " following path ", best.Path) // 366
}

func Part2(riskmap RiskMap) {
	mapSz := Coordinate{len(riskmap), len(riskmap[0])}
	bigMap := make([][]RiskMap, 0)

	// Horizontally
	rowData := make([]RiskMap, 0)
	rowData = append(rowData, riskmap)
	for i := 1; i < 5; i++ {
		rowData = append(rowData, Expand(rowData[i-1]))
	}
	bigMap = append(bigMap, rowData)

	// Vertically
	for i := 1; i < 5; i++ {
		rowData := make([]RiskMap, 0)
		rowData = append(rowData, Expand(bigMap[i-1][0]))
		rowData = append(rowData, Expand(bigMap[i-1][1]))
		rowData = append(rowData, Expand(bigMap[i-1][2]))
		rowData = append(rowData, Expand(bigMap[i-1][3]))
		rowData = append(rowData, Expand(bigMap[i-1][4]))
		bigMap = append(bigMap, rowData)
	}
	// Note there is duplicate work being done here and can be simplified in the future

	// Convert to a conventional RiskMap
	// for i := 0; i < mapSz.col*5; i++ {
	// 	for j := 0; j < mapSz.row*5; j++ {
	// 		fmt.Printf("%d", getFromBigMap(i, j))
	// 	}
	// 	fmt.Printf("\n")
	// }

	graph := dijkstra.NewGraph()
	bigMapSz := Coordinate{mapSz.row * 5, mapSz.col * 5}
	getFromBigMap := func(i, j int) int {
		return bigMap[i/mapSz.row][j/mapSz.col][i%mapSz.row][j%mapSz.col]
	}
	coord2vertexId := func(c Coordinate) int {
		return c.row*bigMapSz.col + c.col
	}
	for i := 0; i < bigMapSz.row; i++ {
		for j := 0; j < bigMapSz.col; j++ {
			graph.AddVertex(coord2vertexId(Coordinate{i, j}))
			if j > 0 {
				graph.AddArc(coord2vertexId(Coordinate{i, j - 1}), coord2vertexId(Coordinate{i, j}), int64(getFromBigMap(i, j)))
				graph.AddArc(coord2vertexId(Coordinate{i, j}), coord2vertexId(Coordinate{i, j - 1}), int64(getFromBigMap(i, j-1)))
			}
			if i > 0 {
				graph.AddArc(coord2vertexId(Coordinate{i - 1, j}), coord2vertexId(Coordinate{i, j}), int64(getFromBigMap(i, j)))
				graph.AddArc(coord2vertexId(Coordinate{i, j}), coord2vertexId(Coordinate{i - 1, j}), int64(getFromBigMap(i-1, j)))
			}
		}
	}

	best, err := graph.Shortest(0, coord2vertexId(Coordinate{bigMapSz.row - 1, bigMapSz.col - 1}))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Part 2 - Shortest distance ", best.Distance) // 2829
}

func Expand(source RiskMap) RiskMap {
	result := make(RiskMap, 0)

	for i := range source {
		rowData := make([]int, 0)

		for j := range source[i] {
			rowData = append(rowData, Increment(source[i][j]))
		}

		result = append(result, rowData)
	}

	return result
}

func Increment(v int) int {
	v++
	if v > 9 {
		return 1
	}
	return v
}

func Loader(filename string) (riskmap RiskMap) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	riskmap = make([][]int, 0)

	for scanner.Scan() {
		rowData := make([]int, 0)
		for _, positionStr := range scanner.Text() {
			i, err := strconv.Atoi(string(positionStr))
			if err != nil {
				log.Fatal(err)
			}
			rowData = append(rowData, i)
		}

		riskmap = append(riskmap, rowData)
	}

	return riskmap
}
