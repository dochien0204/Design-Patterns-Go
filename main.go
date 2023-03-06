package main

import (
	"fmt"
	v "pattern/abstract-factory"
	b "pattern/bridge-pattern"
	"pattern/factory-method"
	counter "pattern/singleton"

	a "pattern/adapter"
)

func main() {

	//Factory-method Design Pattern
	bmw, _ := factory.GetCar("BMW")
	Println(bmw)

	//Singleton Pattern using sync
	// counter.GetInstance().Increase()
	// counter.GetInstance().Increase()
	// fmt.Println("Counter:", counter.GetInstance().Get())

	for i := 0; i < 30; i++ {
		counter.GetInstance().Increase()
	}
	fmt.Println(counter.GetInstance().Get())

	//abstract-factory
	carFactory, _ := v.GetVehicleFactory(v.CarFactoryType)

	luxuryCar, _ := carFactory.GetVehicle(v.LuxuryCar)

	fmt.Println("Luxury car")
	PrintDetail(luxuryCar)

	motorbikeFactory, _ := v.GetVehicleFactory(v.MotorbikeFactoryType)
	sportMotorbike, _ := motorbikeFactory.GetVehicle(v.SportMotorbikeType)
	fmt.Println("Sport Motorbike")
	PrintDetail(sportMotorbike)

	//bridge pattern
	fmt.Println()
	fmt.Println("Bridge Pattern")
	hpPrinter := new(b.HpPrinter)
	epsonPrinter := new(b.EpsonPrinter)

	macComputer := new(b.MacComputer)
	winComputer := new(b.WinComputer)

	macComputer.SetPrinter(hpPrinter)
	macComputer.Print()
	fmt.Println()

	winComputer.SetPrinter(epsonPrinter)
	winComputer.Print()
	fmt.Println()

	//adapter
	client := &a.Client{}
	mac := &a.MacMachine{}

	client.InsertLightningConnectorIntoComputer(mac)

	windowMachine := &a.Window{}
	windowAdapter := &a.WindowAdapter{
		WindowMachine: windowMachine,
	}

	client.InsertLightningConnectorIntoComputer(windowAdapter)

}

func Println(c factory.ICar) {
	fmt.Printf("Type car: %s\n", c.GetTypeCar())
	fmt.Printf("Brand: %s\n", c.GetBrand())
}

func PrintDetail(v v.Vehicle) {
	fmt.Printf("Seats : %d\n", v.GetSeats())
	fmt.Printf("Wheels : %d\n\n", v.GetWheels())
}
