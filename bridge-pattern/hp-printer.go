package computer

import "fmt"

// Concrete Implementations
type HpPrinter struct {
}

func (h *HpPrinter) PrintFile() {
	fmt.Println("Print by HP Printer")
}
