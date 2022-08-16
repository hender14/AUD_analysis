package port

import "github.com/hender14/app/domain"

type UserRepository interface {
	QueryEmail(*domain.InUser) (err error)
	RegisterAccoount(*domain.SignUser) (err error)
}

type UserInputPort interface {
	Sign(input *domain.InUser) (*domain.SignUser, error)
}
