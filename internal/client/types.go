package client

import "github.com/dc7342/nolocks2/internal/entity"

type JWT interface {
	Get(conf entity.NoLocksConfig) (string, error)
	getToken(conf entity.NoLocksConfig) (string, error)
	refresh(conf entity.NoLocksConfig) (string, error)
	verify(conf entity.NoLocksConfig) (bool, error)
}

type HTTP interface {
	Send(loc *entity.Location) error
	Get(loc entity.Location) ([]entity.Location, error)
	GetAll() ([]entity.Location, error)
}
