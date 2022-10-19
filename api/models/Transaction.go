package models

import (
	"errors"
	"html"
	"strings"
	"time"
	"math"
	"strconv"

	"github.com/jinzhu/gorm"
)

type Transaction struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Number     string    `gorm:"size:255;not null;unique" json:"number"`
	Date   time.Time    `gorm:"not null;" json:"date"`
	TransactionDetails []TransactionDetail
	PriceTotal    float64  `gorm:"not null" json:"price_total"`
	CostTotal  float64    `gorm:"not null" json:"cost_total"`
	DeletedOn time.Time
}

func (t *Transaction) Prepare() {
	t.ID = 0
	t.PriceTotal = math.Round(t.PriceTotal)
	t.CostTotal = math.Round(t.CostTotal)
}

func (t *Transaction) Validate() error {

	if t.Date.IsZero() {
		return errors.New("Required Date")
	}
	if len(t.TransactionDetails) == 0 {
		return errors.New("Required Transaction Details")
	}
	return nil
}

func (t *Transaction) SaveTransaction(db *gorm.DB) (*Transaction, error) {
	var err error

	var lastTransaction Transaction
	err = db.Debug().Model(&Transaction{}).Last(&lastTransaction).Error
	if err != nil {
		return &Transaction{}, err
	}
	var tNumber = ""
	if lastTransaction.ID!=0 {
		tNumber =  strconv.FormatUint(lastTransaction.ID, 10)
	} else {
		tNumber = "1"
	}
	currentTime := time.Now()
	if len(tNumber)>1{
		t.Number = currentTime.Format("202210") + html.EscapeString(strings.TrimSpace("-"+tNumber))
	}else{
		t.Number = currentTime.Format("202210") + html.EscapeString(strings.TrimSpace("0"+tNumber))
	}
	
	err = db.Debug().Model(&Transaction{}).Create(&t).Error
	if err != nil {
		return &Transaction{}, err
	}

	for i, _ := range t.TransactionDetails {
		transactionDetail := TransactionDetail{
			TransactionID: t.ID, 
			Item: t.TransactionDetails[i].Item, 
			ItemID: t.TransactionDetails[i].ItemID,
			ItemPrice: t.TransactionDetails[i].ItemPrice,
			ItemCost: t.TransactionDetails[i].ItemCost,
		}

		err = db.Debug().Model(&TransactionDetail{}).Create(&transactionDetail).Error
		if err != nil {
			return &Transaction{}, err
		}
	}
	return t, nil
}

func (t *Transaction) FindAllTransactions(db *gorm.DB) (*[]Transaction, error) {
	var err error
	transactions := []Transaction{}
	transactionDetails := []TransactionDetail{}
	err = db.Debug().Model(&Transaction{}).Limit(100).Find(&transactions).Error
	if err != nil {
		return &[]Transaction{}, err
	}
	if len(transactions) > 0 {
		for i, _ := range transactions {
			err = db.Debug().Model(&TransactionDetail{}).Where("transaction_id = ?", transactions[i].ID).Find(&transactionDetails).Error
			if err != nil {
				return &[]Transaction{}, err
			}
			for j, _ := range transactionDetails {
				item := Item{}
				err = db.Debug().Model(&Item{}).Where("id = ?", transactionDetails[j].ItemID).Find(&item).Error
				if err != nil {
					return &[]Transaction{}, err
				}
				transactionDetails[j].Item = item
			}
			transactions[i].TransactionDetails = transactionDetails
		}
	}
	return &transactions, nil
}

func (t *Transaction) FindTransactionByID(db *gorm.DB, pid uint64) (*Transaction, error) {
	var err error
	transactionDetails := []TransactionDetail{}
	err = db.Debug().Model(&Transaction{}).Where("id = ?", pid).Take(&t).Error
	if err != nil {
		return &Transaction{}, err
	}
	if t.ID != 0 {
		err = db.Debug().Model(&TransactionDetail{}).Where("transaction_id = ?", t.ID).Find(&transactionDetails).Error
		if err != nil {
			return &Transaction{}, err
		}
		for j, _ := range transactionDetails {
			item := Item{}
			err = db.Debug().Model(&Item{}).Where("id = ?", transactionDetails[j].ItemID).Find(&item).Error
			if err != nil {
				return &Transaction{}, err
			}
			transactionDetails[j].Item = item
		}
		t.TransactionDetails = transactionDetails
	}
	return t, nil
}

func (t *Transaction) UpdateATransaction(db *gorm.DB) (*Transaction, error) {

	var err error

	err = db.Debug().Model(&Transaction{}).Where("id = ?", t.ID).Updates(Transaction{Date: t.Date}).Error
	if err != nil {
		return &Transaction{}, err
	}
	if t.ID != 0 {
		for i, _ := range t.TransactionDetails {
			err := db.Debug().Model(&Item{}).Where("id = ?", t.TransactionDetails[i].ItemID).Take(&t.TransactionDetails[i].Item).Error
			if err != nil {
				return &Transaction{}, err
			}
		}
	}
	return t, nil
}

func (t *Transaction) DeleteATransaction(db *gorm.DB, pid uint64) (*Transaction, error) {

	var err error

	err = db.Debug().Model(&Transaction{}).Where("id = ?", pid).Updates(Transaction{DeletedOn: time.Now()}).Error
	if err != nil {
		return &Transaction{}, err
	}
	return t, nil
}