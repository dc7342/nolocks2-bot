package client

import "github.com/dc7342/nolocks2/internal/entity"

type Client struct {
	HTTPClient
}

func NewClient(conf entity.NoLocksConfig) *Client {
	client, err := NewHTTPClient(conf)
	if err != nil {
		panic(err)
	}
	return &Client{HTTPClient: *client}
}
