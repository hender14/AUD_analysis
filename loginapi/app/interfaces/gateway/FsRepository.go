package gateway

import (
	"app/domain"
)

type CRUD interface {
	Fscreate(*domain.SignUser) error
	Fsquery(*Fsqparam) ([]domain.SignUser, error)
}
