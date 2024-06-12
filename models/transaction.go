package models

import (
	"time"

	"github.com/uadmin/uadmin"
)

type Transaction struct {
	uadmin.Model
	Code            string
	Item            Item
	ItemID          uint
	Quantity        int
	TransactionType TransactionType
	Enforcer        Enforcers
	EnforcerID      uint
	CreatedAt       *time.Time	
}

type TransactionType int

func (TransactionType) InventoryIn() TransactionType {
	return 1
}

func (TransactionType) InventoryOut() TransactionType {
	return 2
}
