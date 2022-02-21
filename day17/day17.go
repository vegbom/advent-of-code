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
	// areaXMin := 20
	// areaXMax := 30

	// Actual Puzzle Input:
	// target area: x=138..184, y=-125..-71
	areaYMin := -125
	areaYMax := -71
	areaXMin := 138
	areaXMax := 184

	winner := 0
	validV0s := make([]Coordinate, 0)

	tryInitX := -3500
	tryInitY := -3500

	for yV0 := tryInitY; yV0 < 8000; yV0++ {
		didHitY, stepMin, stepMax, yMax := simulateY(yV0, areaYMin, areaYMax)
		// fmt.Printf("Simulate: y_V0 = %d , Steps: %d-%d, Outcome %v , max Y height: %d\n", yV0, stepMin, stepMax, didHitY, yMax)
		if didHitY && yMax > winner {
			winner = yMax
		}
		if didHitY {
			for xV0 := tryInitX; xV0 < 4500; xV0++ {
				if simulateX(xV0, areaXMin, areaXMax, stepMin, stepMax) {
					validV0s = append(validV0s, Coordinate{xV0, yV0})
					fmt.Printf("Found Valid V0: %d,%d\n", xV0, yV0)
				}
			}
		}
	}

	fmt.Printf("PART 1 - Max Y is: %d\n", winner)
	fmt.Printf("PART 2 - Valid V0 count: %d\n", len(validV0s))
}

func simulateY(yV int, tgtMin int, tgtMax int) (didHit bool, stepMin int, stepMax int, yMax int) {
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
	step := 0
	for inRange(y) == Undershoot {
		step++
		y += yV
		if y > yMax {
			yMax = y
		}
		yV = yVnext(yV)
	}
	stepMin = step

	for inRange(y) == InArea {
		didHit = true
		stepMax = step
		step++
		y += yV
		yV = yVnext(yV)
	}

	return didHit, stepMin, stepMax, yMax
}

func yVnext(yV int) int {
	return yV - 1
}

func simulateX(xV int, tgtMin int, tgtMax int, stepMin int, stepMax int) bool {
	if tgtMax < 0 || tgtMin < 0 {
		log.Fatalf("simulateX This won't work")
	}
	if tgtMax < tgtMin {
		log.Fatalf("simulateX Wrong Input")
	}

	inRange := func(x int) Outcome {
		if x > tgtMax {
			return Overshoot
		} else if x < tgtMin {
			return Undershoot
		}
		return InArea
	}

	x := 0
	step := 0
	for step < stepMax {
		step++
		x += xV
		xV = xVnext(xV)
		if step >= stepMin && inRange(x) == InArea {
			return true
		}
	}

	return inRange(x) == InArea
}

func xVnext(xV int) int {
	if xV > 0 {
		return xV - 1
	} else if xV < 0 {
		return xV + 1
	}
	return 0
}
