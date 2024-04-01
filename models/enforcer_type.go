package models

import (
	"strings"

	"github.com/uadmin/uadmin"
)

type EnforcerType struct {
	uadmin.Model
	Designation string
	Role        string
	Active      bool
}

func (e *EnforcerType) Save() {
	e.Designation = strings.ToUpper(e.Designation)
	e.Role = strings.ToUpper(e.Role)
	uadmin.Save(e)
}
