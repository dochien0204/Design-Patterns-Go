package adapter

import "fmt"

type MacMachine struct {
}

func (m *MacMachine) InsertIntoLightningPort() {
	fmt.Println("Lightning connector is plugged into mac machine")
}
