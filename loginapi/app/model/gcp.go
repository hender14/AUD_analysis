package model

import (
	"context"
	"github.com/jschoedt/go-firestorm"
)

type Fsc struct {
  Ctx     context.Context
  Fsc      *firestorm.FSClient
}

type  Fsqparam struct {
	Collection  string
	Key         string
	Condition   string
  Value       string
}