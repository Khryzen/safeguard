package models

import "github.com/uadmin/uadmin"

type Item struct {
	uadmin.Model
	Name  string
	Image string
}
