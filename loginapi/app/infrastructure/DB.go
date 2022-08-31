package infrastructure

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

type Awscontext struct {
	Db *dynamo.DB
}

func NewDB() (*Awscontext, error) {
	a := new(Awscontext)
	err := a.NoSqlconnect()
	if err != nil {
		return nil, err
	}
	return a, err
}

func (a *Awscontext) NoSqlconnect() error {
	// firestormの初期化
	cred := credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_ID"), os.Getenv("AWS_ACCESS_KEY"), "") // 最後の引数は[セッショントークン]
	a.Db = dynamo.New(session.New(), &aws.Config{
		Credentials: cred,
		Region:      aws.String(os.Getenv("AWS_REGION")),
	})

	return nil
}
