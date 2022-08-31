package infrastructure

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/hender14/app/domain"
	"github.com/hender14/app/interfaces/gateway"
)

type Idcreate struct {
	Id    string `dynamo:"id"`
	Count int    `dynamo:"count"`
}

// Create the entity
func (a *Awscontext) Awscreate(user *domain.SignUser) (err error) {
	if user.ID, err = a.getId(); err != nil {
		fmt.Printf("create entity has problem: %s\n", err)
		return err
	}
	if err := a.Db.Table("SignUser").Put(user).Run(); err != nil {
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
func (a *Awscontext) Awscreate_rst(user *domain.ForgotUser) (*domain.ForgotUser, error) {
	var err error
	if user.ID, err = a.getId(); err != nil {
		fmt.Printf("create entity has problem: %s\n", err)
		return user, err
	}
	if err := a.Db.Table("ForgotUser").Put(user).Run(); err != nil {
		fmt.Printf("create entity has problem: %s\n", err)
		return user, err
	}
	if user.ID == "" {
		err = errors.New("user should have an auto generated ID")
		fmt.Println(err)
	}
	return user, err
}

func (a *Awscontext) getId() (count string, err error) {
	var user *Idcreate
	// if err := a.Db.Table("Count").Get("id", "get").One(user); err != nil {
	if err := a.Db.Table("Count").Get("id", "get").One(&user); err != nil {
		fmt.Printf("create ID problem: %s\n", err)
		return strconv.Itoa(user.Count), err
	}
	user.Count += 1

	if err := a.Db.Table("Count").Put(&user).Run(); err != nil {
		fmt.Printf("read result: %v\n", user)
		fmt.Printf("user was not found by search: %s\n", err)
		return strconv.Itoa(user.Count), err
	}
	count = strconv.Itoa(user.Count)
	return count, err
}

// Read the entity by ID
func (a *Awscontext) Awsread(id string) (rditem *domain.SignUser, err error) {
	rditem = &domain.SignUser{ID: id}
	if err := a.Db.Table("SignUser").Get("id", id).One(&rditem); err != nil {
		fmt.Printf("read result: %s\n", rditem)
		fmt.Printf("user was not found by search: %s\n", err)
		return rditem, err
	}
	return rditem, err
}

// Update the entity
func (a *Awscontext) Awsupdate(user *domain.SignUser) (err error) {
	if err := a.Db.Table("SignUser").Get("id", user.ID).One(&user); err != nil {
		fmt.Printf("read result: %s\n", user)
		fmt.Printf("user was not found by search: %s\n", err)
		return err
	}
	if err := a.Db.Table("SignUser").Get("id", user.ID).All(&user); err != nil {
		// if _, err := a.Awsc.NewRequest().GetEntities(f.Ctx, upitem)(); err != nil {
		fmt.Printf("read result: %s\n", user)
		fmt.Printf("user was not found by search: %s\n", err)
	}
	return err
}

// query the entity
func (a *Awscontext) Awsquery(param *gateway.Fsqparam) (qritem []domain.SignUser, err error) {
	if err = a.Db.Table(param.Collection).Scan().Filter("'"+param.Key+"'"+param.Condition+"?", param.Value).All(&qritem); err != nil {
		fmt.Printf("query result: %s/n", qritem)
		fmt.Printf("query err result: %s/n", err)
		return qritem, err
	}
	return qritem, err
}

func (a *Awscontext) Awsquery_rst(param *gateway.Fsqparam) (qritem []domain.ForgotUser, err error) {
	if err = a.Db.Table(param.Collection).Scan().Filter("'"+param.Key+"'"+param.Condition+"?", param.Value).All(&qritem); err != nil {
		fmt.Printf("query result: %s/n", qritem)
		fmt.Printf("query err result: %s/n", err)
		return qritem, err
	}
	return qritem, err
}

// Delete the entity
func (a *Awscontext) Awsdelete(user *domain.SignUser) (err error) {
	a.Db.Table("SignUser").Delete("id", user.ID).Run()
	dlitem := &domain.SignUser{ID: user.ID}
	if err := a.Db.Table("SignUser").Get("id", user.ID).All(&dlitem); err != nil {
		fmt.Printf("Delete has not completed: %s\n", user.ID)
	}
	return err
}

func (a *Awscontext) Awsdelete_rst(user *domain.ForgotUser) (err error) {
	a.Db.Table("ForgotUser").Delete("id", user.ID).Run()
	dlitem := &domain.ForgotUser{ID: user.ID}
	if err := a.Db.Table("ForgotUser").Get("id", user.ID).All(&dlitem); err != nil {
		fmt.Printf("Delete has not completed: %s\n", user.ID)
	}
	return err
}
