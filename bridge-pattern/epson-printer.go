package computer

import "fmt"

type EpsonPrinter struct{}

func (e *EpsonPrinter) PrintFile() {
	fmt.Println("print by Epson Printer")
}
