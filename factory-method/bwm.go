package factory

type BMWCar struct {
	Car
}

func NewBMWCar() ICar {
	return &BMWCar{
		Car: Car{
			typeCar: "sport cars",
			brand:   "BMW",
		},
	}
}
