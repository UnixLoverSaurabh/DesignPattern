package client

import "DesignPattern/AdapterDesign/example2/answer/model"

type Client struct {
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) InsertSquareUsbInComputer(computer model.Computer) {
	computer.InsertInSquarePort()
}
