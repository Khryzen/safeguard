package models

import "github.com/uadmin/uadmin"

type InventoryLine struct {
	uadmin.Model
	Item      Item
	ItemID    uint
	Remaining int
}
