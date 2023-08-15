package client

import "DesignPattern/AdapterDesign/example2/question/model"

type Client struct {
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) InsertSquareUsbInComputer(mac *model.Mac) {
	mac.InsertInSquarePort()
}
