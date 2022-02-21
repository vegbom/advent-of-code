package main

import (
	"fmt"
	"log"
)

type Coordinate struct {
	x int
	y int
}

type Outcome int

const (
	Undershoot Outcome = 0
	Overshoot  Outcome = 1
	InArea     Outcome = 2
)

func main() {

	// Example input:
	//target area: x=20..30, y=-10..-5
	// areaYMin := -10
	// areaYMax := -5

	// Actual Puzzle Input:
	// target area: x=138..184, y=-125..-71
	areaYMin := -125
	areaYMax := -71

	winner := 0

	for yV0 := 0; yV0 < 2000; yV0++ {
		didHit, steps, maxY := simulateY(yV0, areaYMax, areaYMin)
		fmt.Printf("Simulate: y_V0 = %d , Steps: %d, Outcome %v , max Y height: %d\n", yV0, steps, didHit, maxY)
		if didHit && maxY > winner {
			winner = maxY
		}
	}

	fmt.Printf("Max Y is: %d\n", winner)
}

func simulateY(yV int, tgtMax int, tgtMin int) (bool, int, int) {
	if tgtMax > 0 || tgtMin > 0 {
		log.Fatalf("This won't work")
	}
	if tgtMax < tgtMin {
		log.Fatalf("Wrong Input")
	}

	inRange := func(y int) Outcome {
		if y > tgtMax {
			return Undershoot
		} else if y < tgtMin {
			return Overshoot
		}
		return InArea
	}

	y := 0
	yMax := 0
	steps := 0
	for inRange(y) == Undershoot {
		steps++
		y += yV
		if y > yMax {
			yMax = y
		}
		yV = yVnext(yV)
	}

	return inRange(y) == InArea, steps, yMax
}

func yVnext(yV int) int {
	return yV - 1
}
