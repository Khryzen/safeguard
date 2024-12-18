package models

import "github.com/uadmin/uadmin"

type Enforcers struct {
	uadmin.Model
	User           uadmin.User
	UserID         uint
	Position       string
	EnforcerType   EnforcerType
	EnforcerTypeID uint
	Active         bool
}

func (e *Enforcers) String() string {
	uadmin.Preload(e)

	return e.User.FirstName + " " + e.User.LastName
}
