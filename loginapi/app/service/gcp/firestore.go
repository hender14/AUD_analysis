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

	"app/model"
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
func Fscreate (user *model.User) (err error) {
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
func Fsread (user *model.User) (rditem *model.User, err error) {
  rditem = &model.User{ID:user.ID}
  if _, err := fsc.Fsc.NewRequest().GetEntities(fsc.Ctx, rditem)(); err != nil {
    fmt.Printf("user was not found by search: %s\n", err)
  }
  return rditem, err
}

// Update the entity
func Fsupdate (user *model.User) (upitem *model.User, err error) {
  fsc.Fsc.NewRequest().UpdateEntities(fsc.Ctx, user)()
  upitem = &model.User{ID:user.ID}
  if _, err := fsc.Fsc.NewRequest().GetEntities(fsc.Ctx, upitem)(); err != nil {
    fmt.Printf("user was not found by search: %s\n", err)
}
//   if upitem.ID != user.ID || upitem.Make != user.Make {
//     fmt.Printf("update not reflected : upitem: %v user: %v\n", upitem, user)
// }
  return upitem, err
}

// query the entity
func Fsquery (param *model.Fsqparam ) (qritem []model.User, err error) {
  query := fsc.Fsc.Client.Collection(param.Collection).Where(param.Key, param.Condition, param.Value)
  qritem = make([]model.User, 0)
  if err := fsc.Fsc.NewRequest().QueryEntities(fsc.Ctx, query, &qritem)(); err != nil {
      fmt.Printf("user was not found by search: %s\n", err)
  }
  return qritem, err
}

// Delete the entity
func Fsdelete (user *model.User) (err error) {
  fsc.Fsc.NewRequest().DeleteEntities(fsc.Ctx, user)()
  dlitem := &model.User{ID:user.ID}
  if _, err := fsc.Fsc.NewRequest().GetEntities(fsc.Ctx, dlitem)(); err != nil {
      fmt.Println("Delete has not completed")
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