package vehicle

import "errors"

const (
	CruiseMotorbikeType = 1
	SportMotorbikeType  = 2
)

type MotorbikeFactory struct {
}

func (m *MotorbikeFactory) GetVehicle(v int) (Vehicle, error) {
	switch v {
	case CruiseMotorbikeType:
		return new(CruiseMotorbike), nil
	case SportMotorbikeType:
		return new(SportMotorbike), nil
	}

	return nil, errors.New("Type motorbike %d is not recognized")
}
