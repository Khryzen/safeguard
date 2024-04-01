package models

import "github.com/uadmin/uadmin"

type AlertLevel struct {
	uadmin.Model
	Name  string
	Level int
}
