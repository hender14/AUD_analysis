package gateway

import (
	"github.com/hender14/app/domain"
	"github.com/hender14/app/usecase/port"

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
	context CRUD
}

// NewUserRepository はUserRepositoryを返します．
func NewUserRepository(c CRUD) port.UserRepository {
	return &UserRepository{
		context: c,
	}
}

func (repo *UserRepository) QueryEmail_none(email string) (account domain.SignUser, err error) {
	qparam := &Fsqparam{Collection: "SignUser", Key: "email", Condition: "==", Value: email}
	qrfield, err := repo.context.Fsquery(qparam)
	if err != nil {
		return account, err
	}
	if len(qrfield) != 0 {
		err = errors.New("email had already registered")
		fmt.Printf("email query err: %s/n", err)
		return qrfield[0], err
	}
	return account, err
}

func (repo *UserRepository) QueryEmail(email string) (account domain.SignUser, err error) {
	qparam := &Fsqparam{Collection: "SignUser", Key: "email", Condition: "==", Value: email}
	qrfield, err := repo.context.Fsquery(qparam)
	if err != nil {
		return account, err
	}
	if len(qrfield) == 0 || len(qrfield) == 2 {
		err = errors.New("email had not registered")
		fmt.Printf("email query error: %s/n", err)
		return account, err
	}
	account = qrfield[0]

	if account.ID == "" {
		err = errors.New("id had not registered")
		fmt.Printf("id has not generated: %s/n", err)
		return account, err
	}
	return account, err
}

func (repo *UserRepository) RegisterAccoount(s *domain.SignUser) error {
	err := repo.context.Fscreate(s)
	if err != nil {
		fmt.Printf("create result: %s/n", s)
		fmt.Printf("err result: %s/n", err)
		return err
	}

	return err
}

func (repo *UserRepository) ReadID(id string) (*domain.SignUser, error) {
	account, err := repo.context.Fsread(id)
	if err != nil {
		fmt.Printf("query result: %s/n", account)
		fmt.Printf("query err result: %s/n", err)
		return account, err
	}
	if account.ID == "" {
		err = errors.New("user should have an auto generated ID")
		fmt.Printf("no ID err: %s/n", err)
		return account, err
	}
	return account, err
}

func (repo *UserRepository) DeleteAccount(d *domain.SignUser) error {
	err := repo.context.Fsdelete(d)
	if err != nil {
		return err
	}
	return err
}
