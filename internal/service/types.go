package service

import "nolocks-bot/internal/entity"

type Location interface {
	GetByLocation(loc entity.Location) ([]entity.Location, error)
	GetAll() ([]entity.Location, error)
	Add(comment *entity.Location) error
}
