package computer

import "fmt"

// Concrete Implementation
type EpsonPrinter struct{}

func (e *EpsonPrinter) PrintFile() {
	fmt.Println("print by Epson Printer")
}
