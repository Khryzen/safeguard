package models

import "github.com/uadmin/uadmin"

type Notification struct {
	uadmin.Model
	Name       string
	CellNumber string
	Emergency  string
	Responded  bool
	Enforcer   Enforcers
	EnforcerID uint
}
