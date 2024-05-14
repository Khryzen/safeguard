package models

import (
	"strconv"
	"strings"
	"time"

	"github.com/mbdeguzman/safeguard/sms"
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

	// apiKey := uadmin.GetSetting("SMSAPI").(string)
	sms_users := []SMSSubscriber{}
	uadmin.All(&sms_users)

	// for j := range sms_users {
	// 	sms := SMSRequest{
	// 		From:    "Bongabong MDRRMO",
	// 		To:      sms_users[j].MobileNumber,
	// 		Message: i.IncidentDateStr + "\n" + i.Disasters.Name + "\n" + i.AlertLevel.Name + " at " + i.Location + "\n" + i.Remarks,
	// 	}

	// 	err := sendSMS(apiKey, sms)
	// 	if err != nil {
	// 		fmt.Println("Error:", err)
	// 		return
	// 	}
	// 	fmt.Println("SMS sent successfully!")
	// }

	o := sms.Operator{
		Type:     "SEMAPHORE",
		Username: "71174e4de05abb5b79ac73a25d76fe9b",
		// SenderName: "Safeguard",
	}

	o.Message = i.IncidentDateStr + "\n" + i.Disasters.Name + "\n" + i.AlertLevel.Name + " at " + i.Location + "\n" + i.Remarks
	o.From = "Safeguard 24/7"
	for j := range sms_users {
		cell := ValidateCellPhone(sms_users[j].MobileNumber)
		if cell != "" {
			o.Phone = cell
			o.SendSMS()
		}
	}
}

// ValidateCellPhone !
func ValidateCellPhone(cell string) string {
	if strings.HasPrefix(cell, "+63") && len(cell) == 13 {
		cell = strings.Replace(cell, "+63", "0", 1)
	} else if strings.HasPrefix(cell, "0063") && len(cell) == 13 {
		cell = strings.Replace(cell, "0063", "0", 1)
	} else if strings.HasPrefix(cell, "09") && len(cell) == 11 {
	} else {
		return ""
	}

	_, err := strconv.ParseInt(cell, 10, 64)
	if err != nil {
		return ""
	}
	return cell
}
