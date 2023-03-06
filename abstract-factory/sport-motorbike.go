package vehicle

type SportMotorbike struct{}

func (s *SportMotorbike) GetWheels() int {
	return 2
}

func (s *SportMotorbike) GetSeats() int {
	return 2
}

func (s *SportMotorbike) GetType() int {
	return SportMotorbikeType
}
