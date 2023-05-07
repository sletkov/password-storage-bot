package storage

import "errors"

type Storage interface {
	Set() (int, error)
	Get() (int, error)
	Delete() (int, error)
}

var ErrNoService = errors.New("there is no service")
