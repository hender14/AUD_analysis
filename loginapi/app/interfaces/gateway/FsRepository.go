package gateway

import (
	"app/domain"
)

type Fsc interface {
	Fscreate(*domain.SignUser) error
	Fsquery(*Fsqparam) ([]domain.SignUser, error)
}
