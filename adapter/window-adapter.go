package adapter

import "fmt"

type WindowAdapter struct {
	WindowMachine *Window
}

func (w *WindowAdapter) InsertIntoLightningPort() {
	fmt.Println("Adapter convert lightning signal to USB")
	w.WindowMachine.InsertIntoUSBPort()
}
