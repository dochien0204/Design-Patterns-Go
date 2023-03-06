package computer

import "fmt"

type WinComputer struct {
	printer Printer
}

func (w *WinComputer) Print() {
	fmt.Println("Print request for Win")
	w.printer.PrintFile()
}

func (w *WinComputer) SetPrinter(printer Printer) {
	w.printer = printer
}
