package models

import (
	"strings"

	"github.com/uadmin/uadmin"
)

type Item struct {
	uadmin.Model
	Name  string
	Image string
}

func (i *Item) Save() {
	i.Name = strings.ToUpper(i.Name)
	uadmin.Save(i)
}
