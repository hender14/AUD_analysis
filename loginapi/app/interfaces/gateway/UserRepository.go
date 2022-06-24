package gateway

import (
	"app/domain"

	"errors"
	"fmt"
)

type Fsqparam struct {
	Collection string
	Key        string
	Condition  string
	Value      string
}

type UserRepository struct {
	Fsc Fsc
}

func (repo *UserRepository) QueryEmail(s *domain.InUser) error {
	qparam := &Fsqparam{Collection: "SignUser", Key: "email", Condition: "==", Value: s.Email}
	qrfield, err := repo.Fsc.Fsquery(qparam)
	if err != nil {
		return err
	}
	if len(qrfield) != 0 {
		err = errors.New("email had already registered")
		fmt.Printf("email query err: %s/n", err)
		return err
	}
	return err
}

func (repo *UserRepository) RegisterAccoount(s *domain.SignUser) error {
	err := repo.Fsc.Fscreate(s)
	if err != nil {
		fmt.Printf("create result: %s/n", s)
		fmt.Printf("err result: %s/n", err)
		return err
	}

	return err
}
