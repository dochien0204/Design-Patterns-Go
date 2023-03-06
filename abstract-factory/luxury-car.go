package vehicle

type Luxury struct{}

func (l *Luxury) GetWheels() int {
	return 6
}

func (l *Luxury) GetSeats() int {
	return 6
}

func (l *Luxury) GetDoors() int {
	return 4
}
