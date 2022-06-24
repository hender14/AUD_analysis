package infrastructure

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"

	firebase "firebase.google.com/go"
	"github.com/jschoedt/go-firestorm"
	"google.golang.org/api/option"
)

type Fsc struct {
	Ctx context.Context
	Fsc *firestorm.FSClient
}

func NewDB() (*Fsc, error) {
	f := new(Fsc)
	err := f.NoSqlconnect()
	if err != nil {
		return nil, err
	}
	return f, err
}

func (f *Fsc) NoSqlconnect() error {
	// firestormの初期化
	f.Ctx = context.Background()
	keyfile := Loadenv()
	sa := option.WithCredentialsJSON(keyfile)
	conf := &firebase.Config{ProjectID: os.Getenv("GCP_PROJECT_ID")}
	app, err := firebase.NewApp(f.Ctx, conf, sa)
	if err != nil {
		fmt.Printf("firebase auth has problem:%s\n", err)
		return err
	}
	client, err := app.Firestore(f.Ctx)
	if err != nil {
		fmt.Printf("firestore auth has problem: %s\n", err)
		return err
	}
	f.Fsc = firestorm.New(client, "ID", "")
	return err
}

func Loadenv() []byte {
	// sEnc := base64.StdEncoding.EncodeToString([]byte(JSON))
	// fmt.Println(sEnc)
	sEnc := os.Getenv("GCP_KEYFILE_JSON")
	sDec, _ := base64.StdEncoding.DecodeString(sEnc)

	return sDec
}
