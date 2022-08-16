package gateway

import (
	"github.com/hender14/app/domain"
)

type CRUD interface {
	Fscreate(*domain.SignUser) error
	Fsquery(*Fsqparam) ([]domain.SignUser, error)
}
