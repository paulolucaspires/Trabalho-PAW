package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: no matching record Found")

type projeto struct{
  fullname string
  email string
  telefone int
  website string
  mensagem string
}