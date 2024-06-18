package models

import (
	"fmt"
	"time"

	"github.com/uadmin/uadmin"
)

type Survey struct {
	uadmin.Model
	Usage             int    `uadmin:"read_only;display_name:Usage frequency"`
	Overall           int    `uadmin:"read_only;display_name:Overall Experience"`
	Navigate          int    `uadmin:"read_only;display_name:App Navigation"`
	RTMonitoring      int    `uadmin:"read_only;display_name:Realtime Monitoring"`
	SMS               int    `uadmin:"read_only;display_name:SMS Notifications"`
	IncidentReporting int    `uadmin:"read_only;display_name:Incident Reporting"`
	Chatbot           int    `uadmin:"read_only;display_name:Chatbot Assistance"`
	TrainingModules   int    `uadmin:"read_only;display_name:Provided Training Modules"`
	Comments          string `uadmin:"read_only"`
}

func (s *Survey) String() string {
	return fmt.Sprint(time.Now().Unix())
}
