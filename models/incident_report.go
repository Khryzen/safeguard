package models

import (
	"time"

	"github.com/uadmin/uadmin"
)

type IncidentReport struct {
	uadmin.Model
	IncidentDate    time.Time
	IncidentDateStr string
	Disasters       Disasters
	DisastersID     uint
	AlertLevel      AlertLevel
	AlertLevelID    uint
	Captain         Enforcers
	CaptainID       uint
	Responded       bool
	Remarks         string
	Location        string
	Casualties      int
}

func (i *IncidentReport) Save() {
	formattedTime := i.IncidentDate.Format("January 2, 2006 3:04 PM")
	i.IncidentDateStr = formattedTime

	uadmin.Save(i)
}
