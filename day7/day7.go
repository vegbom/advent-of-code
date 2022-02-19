package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	crabPositions, err := Loader("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	p1 := Part1(crabPositions)
	fmt.Printf("Part 1 (by using median stat): Fuel Use: %d\n", p1)
	p2 := Part2(crabPositions)
	// 96798312 is wrong, because average position is 474.523 which rounds up to 475
	// However, actually by the brute force method, position 474 is the optimal method
	// Not sure why this is happening
	fmt.Printf("Part 2 (by using average stat): Fuel Use: %d\n", p2)

	// Brute Force Part 2
	sort.Ints(crabPositions)
	minFuel := math.MaxInt
	minPos := 0
	for i := crabPositions[0]; i < crabPositions[len(crabPositions)-1]; i++ {
		f := GetFuelUsagePart2(crabPositions, i)
		if f < minFuel {
			minFuel = f
			minPos = i
		}
	}
	fmt.Printf("Part 2 (by brute force) Minimum Pos: %d Fuel: %d\n", minPos, minFuel)
}

func Part1(crabPositions []int) int {
	crabs := make([]int, len(crabPositions))
	copy(crabs, crabPositions)

	// Get the median
	sort.Ints(crabs)
	var median int
	n := len(crabs)
	// fmt.Printf("%v  n=%d\n", crabs, n)
	if n%2 == 1 {
		median = crabs[(n+1)/2]
	} else {
		for offset := 1; offset < n/2; offset++ {
			if crabs[n/2-offset] != crabs[n/2-(offset-1)] {
				median = crabs[n/2+(offset-1)]
				break
			}
			if crabs[n/2+offset] != crabs[n/2+(offset-1)] {
				median = crabs[n/2-(offset-1)]
				break
			}
			if offset == n/2-1 {
				median = crabs[n/2]
			}
		}
	}
	fmt.Printf("Median=%d\n", median)
	return GetFuelUsagePart1(crabs, median)
}

func Part2(crabPositions []int) int {
	// Get the average
	var sum int
	for _, v := range crabPositions {
		sum += v
	}
	average := float64(sum) / float64(len(crabPositions))

	fmt.Printf("Average=%f Rounded: %f\n", average, math.Round(average))

	/*
		Step 1: Solve for the average value
		Step 2: Going from float to int we need to compare floor and ceil for the
			optimum value because you might be rounding in the wrong direction.
			Let's say you have to pick one of the discrete options and you are somewhere in between,
			you need to just compare and pick the optimum one because we don't know which way it's going
			or the movement of the curve at that point
			Need to check it, no way to predict?
	*/

	return GetFuelUsagePart2(crabPositions, int(math.Round(average)))
}

func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func GetFuelUsagePart1(crabPositions []int, finalPosition int) int {
	sum := 0
	for _, pos := range crabPositions {
		fuel := AbsInt(pos - finalPosition)
		sum += fuel
		// fmt.Printf("Move from %d to %d: %d fuel\n", pos, finalPosition, fuel)
	}
	fmt.Printf("Move to: %d Total Fuel: %d\n", finalPosition, sum)
	return sum
}

func GetFuelUsagePart2(crabPositions []int, finalPosition int) int {
	sum := 0.0
	for _, pos := range crabPositions {
		x := AbsInt(pos - finalPosition)
		fuel := 0.5*math.Pow(float64(x), 2) + 0.5*float64(x)
		sum += fuel
		// fmt.Printf("Move from %d to %d: %f fuel\n", pos, finalPosition, fuel)
	}
	// fmt.Printf("Move to %d: Total Fuel: %0.3f\n", finalPosition, sum)
	return int(sum)
}

func Loader(filename string) (crabPositions []int, err error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	crabPositions = nil

	for scanner.Scan() {
		for _, n := range strings.Split(scanner.Text(), ",") {
			i, err := strconv.Atoi(n)
			if err != nil {
				return nil, err
			}
			crabPositions = append(crabPositions, i)
		}
	}

	return crabPositions, err
}
