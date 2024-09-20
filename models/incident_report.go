package models

import (
	"net/smtp"
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

	report := IncidentReport{}
	uadmin.Get(&report, "id = ?", i.ID)
	uadmin.Preload(&report)
	uadmin.Preload(&report.Disasters)
	// apiKey := uadmin.GetSetting("SMSAPI").(string)
	sms_users := []SMSSubscriber{}
	uadmin.All(&sms_users)
	o := sms.Operator{
		Type:     "SEMAPHORE",
		Username: "71174e4de05abb5b79ac73a25d76fe9b",
		// SenderName: "eSafeguard",
	}

	o.Message = report.IncidentDateStr + "\n" + report.Disasters.Name + "\n" + report.AlertLevel.Name + " at " + report.Location + "\n" + report.Remarks
	o.From = "Safeguard 24/7"
	for j := range sms_users {
		cell := ValidateCellPhone(sms_users[j].MobileNumber)
		if cell != "" {
			o.Phone = cell
			o.SendSMS()
		}
		if sms_users[j].EmailAddress != "" {
			// SendEmail(sms_users[j].EmailAddress, o.Message)
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

func SendEmail(recipient string, body string) {
	from := "esafeguard2024@gmail.com"
	password := "shcfingudhxjtkfp"
	to := []string{recipient}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	subject := "Subject: Disaster Notification\r\n"
	message := []byte(subject + body)

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		uadmin.Trail(uadmin.ERROR, "Unable to send email. %s", err.Error())
		return
	}
	uadmin.Trail(uadmin.INFO, "Email sent successfully")
}
