package interactor

import (
	"app/domain"
	"app/usecase/port"
)

type UserInteractor struct {
	User port.UserRepository
}

// NewUserInputPort はUserInputPortを取得します．
func NewUserInputPort(userRepository port.UserRepository) port.UserInputPort {
	return &UserInteractor{
		User: userRepository,
	}
}

// user register
func (interactor *UserInteractor) Sign(input *domain.InUser) (interface{}, error) {
	s := new(domain.SignUser)

	// password check
	err := domain.CheckPassword(input)
	if err != nil {
		return nil, err
	}

	// query the entity
	err = interactor.User.QueryEmail(input)
	if err != nil {
		return input.Email, err
	}

	// create account info
	err = s.CreateAccoount(input)
	if err != nil {
		return s, err
	}

	// user register
	err = interactor.User.RegisterAccoount(s)
	if err != nil {
		return s, err
	}
	return s, nil
}
