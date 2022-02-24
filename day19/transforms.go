package main

func rotate3D(beacons []Coord3D, orientation int, flip bool) []Coord3D {
	// Define orientation as follows:
	// orientation / 3 -> gives the axis.
	//	0 = N/A, 1 = X, 2 = Y, 3 = Z
	//	so it's either steps of 0, 4, 8, or 12.
	// orientation % 4 -> gives the angle.
	//	0 = 0deg, 1 = 90deg CW, 2 = 180deg CW, 3 = 270deg CW
	// Cannot use integer to represent flip as in the case with 2D, because
	// you can still flip while having no rotation

}

func rotate2D(beacons []Coord2D, orientation int) (finalNoFlip []Coord2D, finalFlip []Coord2D) {
	// Define orientation as follows:
	// 0 = 0deg, 1 = 90deg CW, 2 = 180deg CW, 3 = 270deg CW
	// Cannot use integer to represent flip as in the case with 2D, because
	// you can still flip while having no rotation

	// Rotation Matrix
	r := [2][2]int{
		{1, 0},
		{0, 1},
	}
	switch orientation {
	case 1:
		r = [2][2]int{{0, -1}, {1, 0}}
	case 2:
		r = [2][2]int{{-1, 0}, {0, -1}}
	case 3:
		r = [2][2]int{{0, 1}, {-1, 0}}
	}

	// Multiply matrix
	finalNoFlip = make([]Coord2D, 0)
	finalFlip = make([]Coord2D, 0)
	for _, v := range beacons {
		xNew := r[0][0]*v.x + r[0][1]*v.y
		yNew := r[1][0]*v.x + r[1][1]*v.y
		finalNoFlip = append(finalNoFlip, Coord2D{xNew, yNew})
		xNew *= -1
		yNew *= -1
		finalFlip = append(finalFlip, Coord2D{xNew, yNew})
	}

	return finalNoFlip, finalFlip
}
