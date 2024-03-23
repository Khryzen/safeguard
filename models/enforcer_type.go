package models

import "github.com/uadmin/uadmin"

type EnforcerType struct {
	uadmin.Model
	Designation string
	Role        string
	Active      bool
}
