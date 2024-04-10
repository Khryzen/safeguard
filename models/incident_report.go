package models

import (
	"fmt"
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

	apiKey := uadmin.GetSetting("SMSAPI").(string)
	sms_users := []SMSSubscriber{}
	uadmin.All(&sms_users)

	for j := range sms_users {
		sms := SMSRequest{
			From:    "Bongabong MDRRMO",
			To:      sms_users[j].MobileNumber,
			Message: i.IncidentDateStr + "\n" + i.Disasters.Name + "\n" + i.AlertLevel.Name + " at " + i.Location + "\n" + i.Remarks,
		}

		err := sendSMS(apiKey, sms)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("SMS sent successfully!")
	}
}
