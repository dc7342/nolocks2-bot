package service

import "github.com/dc7342/nolocks2/internal/entity"

type Location interface {
	GetByLocation(loc entity.Location) ([]entity.Location, error)
	GetAll() ([]entity.Location, error)
	Add(comment *entity.Location) error
}
