package gcp

import (
	"context"
	"encoding/base64"
	"os"
	"errors"
	"fmt"

	"google.golang.org/api/option"
	firebase "firebase.google.com/go"
	"github.com/jschoedt/go-firestorm"

	"github.com/hender14/app/model"
)

var fsc model.Fsc
// var client *firestore.Client

func NoSqlconnect() (err error) {
  // firestormの初期化
  fsc.Ctx = context.Background()
  keyfile := Loadenv()
  sa := option.WithCredentialsJSON(keyfile)
  conf := &firebase.Config{ProjectID: os.Getenv("GCP_PROJECT_ID")}
  app, err := firebase.NewApp(fsc.Ctx, conf, sa)
  if err != nil {
    fmt.Printf("firebase auth has problem:%s\n", err)
    return err
  }
  client, err := app.Firestore(fsc.Ctx)
  if err != nil {
    fmt.Printf("firestore auth has problem: %s\n", err)
    return err
  }
  fsc.Fsc = firestorm.New(client, "ID", "")
  return nil
}

// Create the entity
func Fscreate (user *model.SignUser) (err error) {
  if err := fsc.Fsc.NewRequest().CreateEntities(fsc.Ctx, user)(); err != nil {
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
func Fscreate_rst (user *model.PasswordReset) (err error) {
  if err := fsc.Fsc.NewRequest().CreateEntities(fsc.Ctx, user)(); err != nil {
    fmt.Printf("user was not found by search: %s\n", err)
    return err
  }
  if user.ID == "" {
    err = errors.New("user should have an auto generated ID")
    fmt.Println(err)
    return err
  }
  return
}

// Read the entity by ID
func Fsread (id string) (rditem *model.SignUser, err error) {
  rditem = &model.SignUser{ID: id}
  if _, err := fsc.Fsc.NewRequest().GetEntities(fsc.Ctx, rditem)(); err != nil {
    fmt.Printf("read result: %s\n", rditem)
    fmt.Printf("user was not found by search: %s\n", err)
    return nil, err
  }
  return rditem, err
}

// Update the entity
func Fsupdate (user *model.SignUser) (err error) {
  fsc.Fsc.NewRequest().UpdateEntities(fsc.Ctx, user)()
  upitem := &model.SignUser{ID: user.ID}
  if _, err := fsc.Fsc.NewRequest().GetEntities(fsc.Ctx, upitem)(); err != nil {
    fmt.Printf("read result: %s\n", upitem)
    fmt.Printf("user was not found by search: %s\n", err)
  }
//   if upitem.ID != user.ID || upitem.Make != user.Make {
//     fmt.Printf("update not reflected : upitem: %v user: %v\n", upitem, user)
// }
  return err
}

// query the entity
func Fsquery (param *model.Fsqparam ) (qritem []model.SignUser, err error) {
  query := fsc.Fsc.Client.Collection(param.Collection).Where(param.Key, param.Condition, param.Value)
  qritem = make([]model.SignUser, 0)
  if err := fsc.Fsc.NewRequest().QueryEntities(fsc.Ctx, query, &qritem)(); err != nil {
    fmt.Printf("query result: %s/n", qritem)
		fmt.Printf("query err result: %s/n", err)
  }
  return qritem, err
}

func Fsquery_rst (param *model.Fsqparam ) (qritem []model.PasswordReset, err error) {
  query := fsc.Fsc.Client.Collection(param.Collection).Where(param.Key, param.Condition, param.Value)
  qritem = make([]model.PasswordReset, 0)
  if err := fsc.Fsc.NewRequest().QueryEntities(fsc.Ctx, query, &qritem)(); err != nil {
    fmt.Printf("user was not found by search: %s\n", err)
  }
  return qritem, err
}

// Delete the entity
func Fsdelete (user *model.SignUser) (err error) {
  fsc.Fsc.NewRequest().DeleteEntities(fsc.Ctx, user)()
  dlitem := &model.SignUser{ID: user.ID}
  if _, err := fsc.Fsc.NewRequest().GetEntities(fsc.Ctx, dlitem)(); err == nil {
    fmt.Printf("Delete has not completed: %s\n", user.ID)
  }
  return err
}

func Fsdelete_rst (user *model.PasswordReset) (err error) {
  fsc.Fsc.NewRequest().DeleteEntities(fsc.Ctx, user)()
  dlitem := &model.PasswordReset{ID:user.ID}
  if _, err := fsc.Fsc.NewRequest().GetEntities(fsc.Ctx, dlitem)(); err == nil {
    fmt.Printf("Delete has not completed: %s\n", user.ID)
  }
  return err
}

func Loadenv () ([]byte) {
  // sEnc := base64.StdEncoding.EncodeToString([]byte(JSON))
  // fmt.Println(sEnc)
  sEnc := os.Getenv("GCP_KEYFILE_JSON")
  sDec, _ := base64.StdEncoding.DecodeString(sEnc)

  return sDec
}