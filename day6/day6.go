package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"vegbom/day6/fish"
)

const VERBOSE = false

func main() {
	fishTimers, err := Loader("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Initial State: %v\n", fishTimers)

	initialFishes := make([]fish.Fish, 0)

	for _, days := range fishTimers {
		initialFishes = append(initialFishes, fish.New(days))
	}

	// Part 1
	// SimulateFishByStructs(initialFishes, 18)
	SimulateFishByStructs(initialFishes, 80) // 362346
	SimluateFishByBuckets(initialFishes, 80) // Answer should match

	// Part 2
	SimluateFishByBuckets(initialFishes, 256)
}

func SimluateFishByBuckets(initialFishes []fish.Fish, days int) int {
	ageBuckets := make([]int, 9)
	for _, f := range initialFishes {
		ageBuckets[f.DaysLeftToSpawn]++
	}

	for d := 0; d < days; d++ {
		ageBuckets = append(ageBuckets[1:], ageBuckets[0])
		ageBuckets[6] += ageBuckets[8]
	}

	sum := 0
	for i := 0; i < 9; i++ {
		sum += ageBuckets[i]
	}
	fmt.Printf("BY BUKETS Day %d  Count: %d \n", days, sum)
	return sum
}

func SimulateFishByStructs(initialFishes []fish.Fish, days int) int {
	fishes := make([]fish.Fish, len(initialFishes))
	copy(fishes, initialFishes)

	for i := 1; i <= days; i++ {
		for fishID, currentFish := range fishes {
			if currentFish.LiveAnotherDay() {
				// Spawn new fish
				fishes = append(fishes, fish.New())
			}
			fishes[fishID] = currentFish
		}
		if VERBOSE {
			fmt.Printf("After %3d day(s): Count: %d || ", i, len(fishes))
			for _, currentFish := range fishes {
				fmt.Printf("%d,", currentFish.DaysLeftToSpawn)
			}
			fmt.Printf("\n")
		}
	}

	fmt.Printf("After %d Days there are total %d fish\n", days, len(fishes))
	return len(fishes)
}

func Loader(filename string) (fishTimers []int, err error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	fishTimers = nil

	for scanner.Scan() {
		for _, n := range strings.Split(scanner.Text(), ",") {
			i, err := strconv.Atoi(n)
			if err != nil {
				return nil, err
			}
			fishTimers = append(fishTimers, i)
		}
	}

	return fishTimers, err
}
