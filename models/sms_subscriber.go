package models

import "github.com/uadmin/uadmin"

type SMSSubscriber struct {
	uadmin.Model
	FirstName    string
	LastName     string
	Address      string
	MobileNumber string
	EmailAddress string
}
