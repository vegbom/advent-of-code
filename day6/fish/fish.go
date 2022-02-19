package fish

type Fish struct {
	DaysLeftToSpawn int
}

func New(days ...int) Fish {
	f := Fish{8}
	if len(days) > 0 {
		f.DaysLeftToSpawn = days[0]
	}
	// fmt.Printf("New Fish Created %d \n", f.DaysLeftToSpawn)
	return f
}

func (f *Fish) LiveAnotherDay() bool {
	if f.DaysLeftToSpawn == 0 {
		f.DaysLeftToSpawn = 6
		return true
	}

	f.DaysLeftToSpawn--
	return false
}
