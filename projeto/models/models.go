package models

import (
	"errors"
	"time"
  "github.com/paulolucaspires/Trabalho-PAW/projeto/models
)

var ErrNoRecord = errors.New("models: no matching record Found")

type projeto struct{
  ID int
  Title string
  Content string
  Created time.Time
  Expires time.Time
}