package computer

import "fmt"

// Abstraction
type MacComputer struct {
	printer Printer
}

func (m *MacComputer) Print() {
	fmt.Println("Println request for Mac")
	m.printer.PrintFile()
}

func (m *MacComputer) SetPrinter(printer Printer) {
	m.printer = printer
}
