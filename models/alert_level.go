package models

import (
	"strings"

	"github.com/uadmin/uadmin"
)

type AlertLevel struct {
	uadmin.Model
	Name  string
	Level int
}

func (a *AlertLevel) Save() {
	a.Name = strings.ToUpper(a.Name)

	uadmin.Save(a)
}
