package models

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type TransactionDetail struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	TransactionID    uint64      `json:"transaction_id"`
	Item    Item      `json:"item"`
	ItemID  uint64    `gorm:"not null" json:"item_id"`
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
	if td.ItemQuantity < 0 {
		return errors.New("Required Item Quantity")
	}
	if td.ItemPrice < 0 {
		return errors.New("Required Item Price")
	}
	if td.ItemCost < 0 {
		return errors.New("Required Item Cost")
	}
	return nil
}

func (t *Transaction) SaveTransactionDetails(db *gorm.DB, transactionID uint64) (*Transaction, error) {
	var err error
	transactionDetails := t.TransactionDetails
	for i, _ := range transactionDetails {
		transactionDetail := TransactionDetail{
			TransactionID: transactionID, 
			Item: transactionDetails[i].Item, 
			ItemID: transactionDetails[i].ItemID,
			ItemPrice: transactionDetails[i].ItemPrice,
			ItemCost: transactionDetails[i].ItemCost,
		}
		err = db.Debug().Model(&TransactionDetail{}).Create(&transactionDetail).Error
		if err != nil {
			return &Transaction{}, err
		}
	}
	return t, nil
}