package gateway

import (
	"github.com/hender14/app/domain"
	"github.com/sendgrid/rest"
)

type CRUD interface {
	Awscreate(*domain.SignUser) error
	Awsquery(*Fsqparam) ([]domain.SignUser, error)
	Awsquery_rst(*Fsqparam) ([]domain.ForgotUser, error)
	Awsread(string) (*domain.SignUser, error)
	Awsdelete(*domain.SignUser) error
	Awsdelete_rst(*domain.ForgotUser) error
	Awscreate_rst(*domain.ForgotUser) (*domain.ForgotUser, error)
	Awsupdate(*domain.SignUser) error
}

type MAIL interface {
	Reqconf([]byte) (*rest.Response, error)
}
