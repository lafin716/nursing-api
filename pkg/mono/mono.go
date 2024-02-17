package mono

import "nursing_api/pkg/mono/date"

type Client struct {
	Date *date.DateClient
}

func NewMono() *Client {
	return &Client{
		Date: date.New(),
	}
}
