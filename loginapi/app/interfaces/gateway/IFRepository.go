package gateway

import (
	"github.com/hender14/app/domain"
	"github.com/sendgrid/rest"
)

type CRUD interface {
	Fscreate(*domain.SignUser) error
	Fsquery(*Fsqparam) ([]domain.SignUser, error)
	Fsquery_rst(*Fsqparam) ([]domain.ForgotUser, error)
	Fsread(string) (*domain.SignUser, error)
	Fsdelete(*domain.SignUser) error
	Fsdelete_rst(*domain.ForgotUser) error
	Fscreate_rst(*domain.ForgotUser) (*domain.ForgotUser, error)
	Fsupdate(*domain.SignUser) error
}

type MAIL interface {
	Reqconf([]byte) (*rest.Response, error)
}
