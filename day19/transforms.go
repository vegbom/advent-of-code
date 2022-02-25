package main

import (
	"log"
	"math"
)

type Orientation struct {
	axis     int
	rotation int
	flip     bool
}

// func allOrientations2D() (orientations []Orientation) {
// 	orientations = make([]Orientation, 0)
// 	for rotation := 0; rotation <= 3; rotation++ {
// 		orientations = append(orientations, Orientation{rotation, false})
// 		orientations = append(orientations, Orientation{rotation, true})
// 	}
// 	return orientations
// }

func allOrientations3D() (orientations []Orientation) {
	orientations = make([]Orientation, 0)
	for axis := 0; axis < 3; axis++ {
		for rotation := 0; rotation < 4; rotation++ {
			orientations = append(orientations, Orientation{axis, rotation, false})
			orientations = append(orientations, Orientation{axis, rotation, true})
			// orientations = append(orientations, Orientation{rotation, true})
		}
	}
	return orientations
}

func rotate3D(beacons []Coord3D, orientation Orientation) (final []Coord3D) {
	// Define orientation as follows:
	// orientation / 3 -> gives the axis.
	//	0 = N/A, 1 = X, 2 = Y, 3 = Z
	//	so it's either steps of 0, 4, 8, or 12.
	// orientation % 4 -> gives the angle.
	//	0 = 0deg, 1 = 90deg CW, 2 = 180deg CW, 3 = 270deg CW
	// Cannot use integer to represent flip as in the case with 2D, because
	// you can still flip while having no rotation

	a := 0.0 // angle
	switch orientation.rotation {
	case 1:
		a = math.Pi / 2
	case 2:
		a = math.Pi
	case 3:
		a = (math.Pi * 3) / 2
	}

	// Rotation Matrix
	r := [3][3]float64{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}
	switch orientation.axis {
	case 0:
		// X
		r = [3][3]float64{
			{1, 0, 0},
			{0, math.Cos(a), -math.Sin(a)},
			{0, math.Sin(a), math.Cos(a)},
		}
	case 1:
		// Y
		r = [3][3]float64{
			{math.Cos(a), 0, math.Sin(a)},
			{0, 1, 0},
			{-math.Sin(a), 0, math.Cos(a)},
		}
	case 2:
		// Z
		r = [3][3]float64{
			{math.Cos(a), -math.Sin(a), 0},
			{math.Sin(a), math.Cos(a), 0},
			{0, 0, 1},
		}
	default:
		log.Fatal("Wrong axis")
	}

	// Multiply matrix
	final = make([]Coord3D, 0)
	for _, v := range beacons {
		if orientation.flip {
			switch orientation.axis {
			case 0:
				v.x *= -1
			case 1:
				v.y *= -1
			case 2:
				v.z *= -1
			}
		}
		// WARNING: If you instead write
		// xNew := r[0][0]*float64(v.x) + r[0][1]*float64(v.y) + r[0][2]*float64(v.z)
		// that will cause incorrect result due to floating point errors and rounding!
		xNew := int(r[0][0])*v.x + int(r[0][1])*v.y + int(r[0][2])*v.z
		yNew := int(r[1][0])*v.x + int(r[1][1])*v.y + int(r[1][2])*v.z
		zNew := int(r[2][0])*v.x + int(r[2][1])*v.y + int(r[2][2])*v.z
		final = append(final, Coord3D{int(xNew), int(yNew), int(zNew)})
	}
	return final
}

func rotate2D(beacons []Coord2D, orientation Orientation) (final []Coord2D) {
	// Define orientation as follows:
	// 0 = 0deg, 1 = 90deg CW, 2 = 180deg CW, 3 = 270deg CW
	// Cannot use integer to represent flip as in the case with 2D, because
	// you can still flip while having no rotation

	// Rotation Matrix
	r := [2][2]int{
		{1, 0},
		{0, 1},
	}
	switch orientation.rotation {
	case 1:
		r = [2][2]int{{0, -1}, {1, 0}}
	case 2:
		r = [2][2]int{{-1, 0}, {0, -1}}
	case 3:
		r = [2][2]int{{0, 1}, {-1, 0}}
	}

	// Multiply matrix
	final = make([]Coord2D, 0)
	for _, v := range beacons {
		xNew := r[0][0]*v.x + r[0][1]*v.y
		yNew := r[1][0]*v.x + r[1][1]*v.y
		if orientation.flip {
			xNew *= -1
			yNew *= -1
		}
		final = append(final, Coord2D{xNew, yNew})
	}

	return final
}
