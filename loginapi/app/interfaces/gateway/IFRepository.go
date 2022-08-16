package gateway

import (
	"github.com/hender14/app/domain"
)

type CRUD interface {
	Fscreate(*domain.SignUser) error
	Fsquery(*Fsqparam) ([]domain.SignUser, error)
	Fsread(string) (*domain.SignUser, error)
	Fsdelete(*domain.SignUser) error
}
