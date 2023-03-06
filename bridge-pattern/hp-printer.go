package computer

import "fmt"

type HpPrinter struct {
}

func (h *HpPrinter) PrintFile() {
	fmt.Println("Print by HP Printer")
}
