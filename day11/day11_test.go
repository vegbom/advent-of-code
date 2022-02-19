package main

import (
	"fmt"
	"testing"
)

func TestGetAdjacentCoords(t *testing.T) {
	mapSz := Coordinate{10, 10}

	coord1 := Coordinate{0, 0}
	result := GetAdjacentCoords(coord1, mapSz)
	fmt.Printf("%v\n", result)

	coord1 = Coordinate{9, 9}
	result = GetAdjacentCoords(coord1, mapSz)
	fmt.Printf("%v\n", result)

	coord1 = Coordinate{5, 5}
	result = GetAdjacentCoords(coord1, mapSz)
	fmt.Printf("%v\n", result)

	coord1 = Coordinate{0, 9}
	result = GetAdjacentCoords(coord1, mapSz)
	fmt.Printf("%v\n", result)
}
