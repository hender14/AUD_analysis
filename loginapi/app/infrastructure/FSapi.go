package infrastructure

import (
	"errors"
	"fmt"

	"github.com/hender14/app/domain"
	"github.com/hender14/app/interfaces/gateway"
)

// Create the entity
func (f *Fscontext) Fscreate(user *domain.SignUser) (err error) {
	if err := f.Fsc.NewRequest().CreateEntities(f.Ctx, user)(); err != nil {
		fmt.Printf("create entity has problem: %s\n", err)
		return err
	}
	if user.ID == "" {
		err = errors.New("user should have an auto generated ID")
		fmt.Println(err)
	}
	return err
}

// Create the entity(while reset)
func (f *Fscontext) Fscreate_rst(user *domain.ForgotUser) (*domain.ForgotUser, error) {
	var err error

	if err := f.Fsc.NewRequest().CreateEntities(f.Ctx, user)(); err != nil {
		fmt.Printf("user was not found by search: %s\n", err)
		return user, err
	}
	if user.ID == "" {
		err := errors.New("user should have an auto generated ID")
		fmt.Println(err)
		return user, err
	}
	return user, err
}

// Read the entity by ID
func (f *Fscontext) Fsread(id string) (rditem *domain.SignUser, err error) {
	rditem = &domain.SignUser{ID: id}
	if _, err := f.Fsc.NewRequest().GetEntities(f.Ctx, rditem)(); err != nil {
		fmt.Printf("read result: %s\n", rditem)
		fmt.Printf("user was not found by search: %s\n", err)
		return rditem, err
	}
	return rditem, err
}

// Update the entity
func (f *Fscontext) Fsupdate(user *domain.SignUser) (err error) {
	f.Fsc.NewRequest().UpdateEntities(f.Ctx, user)()
	upitem := &domain.SignUser{ID: user.ID}
	if _, err := f.Fsc.NewRequest().GetEntities(f.Ctx, upitem)(); err != nil {
		fmt.Printf("read result: %s\n", upitem)
		fmt.Printf("user was not found by search: %s\n", err)
	}
	//   if upitem.ID != user.ID || upitem.Make != user.Make {
	//     fmt.Printf("update not reflected : upitem: %v user: %v\n", upitem, user)
	// }
	return err
}

// query the entity
func (f *Fscontext) Fsquery(param *gateway.Fsqparam) (qritem []domain.SignUser, err error) {
	query := f.Fsc.Client.Collection(param.Collection).Where(param.Key, param.Condition, param.Value)
	qritem = make([]domain.SignUser, 0)
	if err := f.Fsc.NewRequest().QueryEntities(f.Ctx, query, &qritem)(); err != nil {
		fmt.Printf("query result: %s/n", qritem)
		fmt.Printf("query err result: %s/n", err)
	}
	return qritem, err
}

func (f *Fscontext) Fsquery_rst(param *gateway.Fsqparam) (qritem []domain.ForgotUser, err error) {
	query := f.Fsc.Client.Collection(param.Collection).Where(param.Key, param.Condition, param.Value)
	qritem = make([]domain.ForgotUser, 0)
	if err := f.Fsc.NewRequest().QueryEntities(f.Ctx, query, &qritem)(); err != nil {
		fmt.Printf("user was not found by search: %s\n", err)
	}
	return qritem, err
}

// Delete the entity
func (f *Fscontext) Fsdelete(user *domain.SignUser) (err error) {
	f.Fsc.NewRequest().DeleteEntities(f.Ctx, user)()
	dlitem := &domain.SignUser{ID: user.ID}
	if _, err := f.Fsc.NewRequest().GetEntities(f.Ctx, dlitem)(); err == nil {
		fmt.Printf("Delete has not completed: %s\n", user.ID)
	}
	return err
}

func (f *Fscontext) Fsdelete_rst(user *domain.ForgotUser) (err error) {
	f.Fsc.NewRequest().DeleteEntities(f.Ctx, user)()
	dlitem := &domain.ForgotUser{ID: user.ID}
	if _, err := f.Fsc.NewRequest().GetEntities(f.Ctx, dlitem)(); err == nil {
		fmt.Printf("Delete has not completed: %s\n", user.ID)
	}
	return err
}
