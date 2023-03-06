package main

import (
	"fmt"
	"pattern/factory-method"
)

func main() {
	bmw, _ := factory.GetCar("BMW")
	Println(bmw)
}

func Println(c factory.ICar) {
	fmt.Printf("Type car: %s\n", c.GetTypeCar())
	fmt.Printf("Brand: %s\n", c.GetBrand())
}
