package factory

type Car struct {
	typeCar string
	brand   string
}

func (c *Car) GetTypeCar() string {
	return c.typeCar
}

func (c *Car) GetBrand() string {
	return c.brand
}

func (c *Car) SetTypeCar(typeCar string) {
	c.typeCar = typeCar
}

func (c *Car) SetBrand(brand string) {
	c.brand = brand
}
