package port

import "app/domain"

type UserRepository interface {
	QueryEmail(*domain.InUser) (err error)
	RegisterAccoount(*domain.SignUser) (err error)
}

type UserInputPort interface {
	Sign(input *domain.InUser) (interface{}, error)
}
