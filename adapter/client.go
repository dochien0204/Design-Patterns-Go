package adapter

import "fmt"

type Client struct {
}

func (c *Client) InsertLightningConnectorIntoComputer(computer Computer) {
	fmt.Println("Client inserts Lightning connector into computer ")
	computer.InsertIntoLightningPort()
}
