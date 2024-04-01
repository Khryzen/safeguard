package models

import (
	"strings"

	"github.com/uadmin/uadmin"
)

type Disasters struct {
	uadmin.Model
	Name        string
	Description string
}

func (d *Disasters) String() string {
	return d.Name
}

func (d *Disasters) Save() {
	d.Name = strings.ToUpper(d.Name)
	uadmin.Save(d)
}
