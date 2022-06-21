package service

import "nolocks-bot/internal/client"

type Service struct {
	Location
}

func NewService(client *client.Client) *Service {
	return &Service{Location: NewLocationService(client)}
}
