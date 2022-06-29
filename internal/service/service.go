package service

import "github.com/dc7342/nolocks2/internal/client"

type Service struct {
	Location
}

func NewService(client *client.Client) *Service {
	return &Service{Location: NewLocationService(client)}
}
