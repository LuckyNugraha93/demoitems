package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type TransactionDetail struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	TransactionID    uint32      `json:"transaction_id"`
	Item    Item      `json:"item"`
	ItemID  uint32    `gorm:"not null" json:"item_id"`
	ItemQuantity   int    `gorm:"not null;" json:"item_quantity"`
	ItemPrice  float32    `gorm:"not null;" json:"item_price"`
	ItemCost   float32    `gorm:"not null;" json:"item_cost"`
}

func (td *TransactionDetail) Prepare() {
	td.ID = 0
	td.Item = Item{}
}

func (td *TransactionDetail) Validate() error {

	if td.TransactionID < 1 {
		return errors.New("Required Transaction")
	}
	if td.ItemID < 1 {
		return errors.New("Required Item")
	}
	if i.ItemQuantity < 0 {
		return errors.New("Required Item Quantity")
	}
	if i.ItemPrice < 0 {
		return errors.New("Required Item Price")
	}
	if i.ItemCost < 0 {
		return errors.New("Required Item Cost")
	}
	return nil
}