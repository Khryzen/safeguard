package models

import (
	"time"

	"github.com/uadmin/uadmin"
)

type IncidentReport struct {
	uadmin.Model
	IncidentDate time.Time
	Disasters    Disasters
	DisastersID  uint
	AlertLevel   AlertLevel
	AlertLevelID uint
	Captain      Enforcers
	CaptainID    uint
	Responded    bool
	Remarks      string
	Location     string
	Casualties   int
}
