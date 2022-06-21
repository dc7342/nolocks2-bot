package client

import "nolocks-bot/internal/entity"

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
