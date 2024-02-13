package models

import "github.com/uadmin/uadmin"

type Enforcers struct {
	uadmin.Model
	User     uadmin.User
	UserID   uint
	Position string
	Active   bool
}
