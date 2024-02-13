package models

import "github.com/uadmin/uadmin"

type Protocols struct {
	uadmin.Model
	Name        string
	Code        string
	Disasters   Disasters
	DisastersID uint
}
