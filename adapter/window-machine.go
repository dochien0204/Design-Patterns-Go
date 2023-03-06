package adapter

import "fmt"

type Window struct {
}

func (w *Window) InsertIntoUSBPort() {
	fmt.Println("USB connector is plugged into windows machine")
}
