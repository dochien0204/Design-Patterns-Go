package factory

type ICar interface {
	GetTypeCar() string
	GetBrand() string
	SetTypeCar(typeCar string)
	SetBrand(brand string)
}
