package gateway

import (
	"errors"
	"fmt"
	"time"

	"github.com/hender14/app/domain"
	"github.com/hender14/app/usecase/port"
)

type Fsqparam struct {
	Collection string
	Key        string
	Condition  string
	Value      string
}

type UserRepository struct {
	context CRUD
	config  MAIL
}

func NewUserRepository(c CRUD, s MAIL) port.UserRepository {
	return &UserRepository{
		context: c,
		config:  s,
	}
}

func (repo *UserRepository) QueryEmail_none(email string) (account *domain.SignUser, err error) {
	qparam := &Fsqparam{Collection: "SignUser", Key: "email", Condition: "=", Value: email}
	qrfield, err := repo.context.Awsquery(qparam)
	if err != nil {
		return account, err
	}
	if len(qrfield) != 0 {
		err = errors.New("email had already registered")
		fmt.Printf("email query err: %s/n", err)
		return &qrfield[0], err
	}
	return account, err
}

func (repo *UserRepository) QueryEmail(email string) (account *domain.SignUser, err error) {
	qparam := &Fsqparam{Collection: "SignUser", Key: "email", Condition: "=", Value: email}
	qrfield, err := repo.context.Awsquery(qparam)
	if err != nil {
		return account, err
	}
	if len(qrfield) == 0 || len(qrfield) == 2 {
		err = errors.New("email had not registered")
		fmt.Printf("email query error: %s/n", err)
		return account, err
	}
	account = &qrfield[0]

	if account.ID == "" {
		err = errors.New("id had not registered")
		fmt.Printf("id has not generated: %s/n", err)
		return account, err
	}
	return account, err
}

func (repo *UserRepository) RegisterAccoount(s *domain.SignUser) error {
	err := repo.context.Awscreate(s)
	if err != nil {
		fmt.Printf("create result: %s/n", s)
		fmt.Printf("err result: %s/n", err)
		return err
	}

	return err
}

func (repo *UserRepository) ReadID(id string) (*domain.SignUser, error) {
	account, err := repo.context.Awsread(id)
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
	err := repo.context.Awsdelete(d)
	if err != nil {
		return err
	}
	return err
}

func (repo *UserRepository) ResetAccount(email string, token string) (*domain.ForgotUser, error) {
	f := new(domain.ForgotUser)
	f.Email = email
	f.Token = token
	f.Year = time.Now()

	f, err := repo.context.Awscreate_rst(f)
	if err != nil {
		fmt.Printf("create result: %s/n", f)
		fmt.Printf("err result: %s/n", err)
		return f, err
	}
	return f, err
}

func (repo *UserRepository) QueryToken(token string) (account *domain.ForgotUser, err error) {
	qparam := &Fsqparam{Collection: "ForgotUser", Key: "token", Condition: "=", Value: token}
	qrfield, err := repo.context.Awsquery_rst(qparam)
	if err != nil {
		return account, err
	}
	if len(qrfield) == 0 || len(qrfield) == 2 {
		err = errors.New("reset had not registered")
		fmt.Printf("token query error: %s/n", err)
		return account, err
	}
	account = &qrfield[0]
	return account, err
}

func (repo *UserRepository) UpdateAccoount(account *domain.SignUser, password string) (*domain.SignUser, error) {
	// password encode
	passwordb, err := domain.EncodePassword(password)
	if err != nil {
		fmt.Printf("password encode err: %s\n", err)
		return account, err
	}

	account.Password = passwordb
	account.Year = time.Now()

	err = repo.context.Awsupdate(account)
	if err != nil {
		fmt.Printf("update result: %s/n", account)
		fmt.Printf("err result: %s/n", err)
		return account, err
	}
	return account, err
}

func (repo *UserRepository) DeleteAccoount(account *domain.ForgotUser) (err error) {
	err = repo.context.Awsdelete_rst(account)
	if err != nil {
		return err
	}
	return err
}
