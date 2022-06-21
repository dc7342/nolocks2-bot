package service

import (
	"nolocks-bot/internal/client"
	"nolocks-bot/internal/entity"
)

type LocationService struct {
	client *client.Client
}

func NewLocationService(client *client.Client) *LocationService {
	return &LocationService{client: client}
}

func (l *LocationService) Add(loc *entity.Location) error {
	err := l.client.Send(loc)
	if err != nil {
		return err
	}

	return nil
}

func (l *LocationService) GetAll() ([]entity.Location, error) {
	// Not implemented yet.
	return nil, nil
}

func (l *LocationService) GetByLocation(loc entity.Location) ([]entity.Location, error) {
	// Not implemented yet.
	return nil, nil
}
