package vehicle

type Familiar struct {
}

func (f *Familiar) GetWheels() int {
	return 4
}

func (f *Familiar) GetSeats() int {
	return 4
}

func (f *Familiar) GetDoors() int {
	return 4
}
