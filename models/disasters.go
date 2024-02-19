package models

import "github.com/uadmin/uadmin"

type Disasters struct {
	uadmin.Model
	Name         string
	Description  string
	AlertLevel   AlertLevel
	AlertLevelID uint
}

func (d *Disasters) String() string {
	return d.Name
}