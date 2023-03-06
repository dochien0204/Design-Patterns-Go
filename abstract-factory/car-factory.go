package vehicle

import "errors"

const (
	LuxuryCar   = 1
	FamiliarCar = 2
)

type CarFactory struct {
}

func (c *CarFactory) GetVehicle(v int) (Vehicle, error) {
	switch v {
	case LuxuryCar:
		return new(Luxury), nil
	case FamiliarCar:
		return new(Familiar), nil
	}

	return nil, errors.New("Vehicle type %d is not recognized")
}
