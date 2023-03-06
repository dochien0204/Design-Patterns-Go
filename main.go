package main

import (
	"fmt"
	"pattern/factory-method"
	counter "pattern/singleton"
)

func main() {

	//Factory-method Design Pattern
	bmw, _ := factory.GetCar("BMW")
	Println(bmw)

	//Singleton Pattern
	counter.GetInstance().Increase()
	counter.GetInstance().Increase()
	fmt.Println("Counter:", counter.GetInstance().Get())

}

func Println(c factory.ICar) {
	fmt.Printf("Type car: %s\n", c.GetTypeCar())
	fmt.Printf("Brand: %s\n", c.GetBrand())
}
