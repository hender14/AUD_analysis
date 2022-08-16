package port

import "github.com/hender14/app/domain"

type UserRepository interface {
	QueryEmail(*domain.InUser) (err error)
	RegisterAccoount(*domain.SignUser) (err error)
	ReadID(*domain.SignUser) error
	DeleteAccount(*domain.SignUser) error
}

type UserInputPort interface {
	Sign(input *domain.InUser) (*domain.SignUser, error)
	Delete(input *domain.SignUser) (*domain.SignUser, error)
}
