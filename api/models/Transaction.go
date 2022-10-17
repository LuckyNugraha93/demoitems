package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Transaction struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Number     string    `gorm:"size:255;not null;unique" json:"number"`
	Date   time.Time    `gorm:"not null;" json:"date"`
	TransactionDetails []TransactionDetail
	PriceTotal    User  `gorm:"not null" json:"price_total"`
	CostTotal  uint32    `gorm:"not null" json:"cost_total"`
	DeletedAt time.Time `gorm:"index" json:"deleted_at"`
}

func (t *Transaction) Prepare() {
	t.ID = 0
	t.PriceTotal = math.Round(t.PriceTotal)
	t.CostTotal = math.Round(t.CostTotal)
}

func (t *Transaction) Validate() error {

	if t.Number == "" {
		return errors.New("Required Number")
	}
	if t.Date == nil {
		return errors.New("Required Date")
	}
	if t.TransactionDetails == 0 {
		return errors.New("Required Transaction Details")
	}
	return nil
}

func (t *Transaction) SaveTransaction(db *gorm.DB) (*Transaction, error) {
	var err error, errDetail error
	for i, _ := range t.TransactionDetails {
		errDetail = db.Debug().Model(&TransactionDetail{}).Create(&t.TransactionDetails[i]).Error
		if errDetail != nil {
			return &Transaction{}, errDetail
		}
	}
	var tNumber = ""
	if len([]rune(t.Number)) < 2 {
		for i, _ := (2-len([]rune(t.Number))) {
			tNumber = tNumber + "0"
		}
		tNumber = tNumber + t.Number;
	}
	currentTime := time.Now()
	t.Number = currentTime.Format("20220101") + html.EscapeString(strings.TrimSpace(tNumber)
	err = db.Debug().Model(&Transaction{}).Create(&t).Error
	if err != nil {
		return &Transaction{}, err
	}
	return t, nil
}

func (t *Transaction) FindAllTransactions(db *gorm.DB) (*[]Transaction, error) {
	var err error
	transactions := []Transaction{}
	err = db.Debug().Model(&Transaction{}).Limit(100).Find(&transactions).Error
	if err != nil {
		return &[]Transaction{}, err
	}
	if len(transactions) > 0 {
		for i, _ := range transactions {
			for j, _ := range transactions[i].TransactionDetails {
				err := db.Debug().Model(&Item{}).Where("id = ?", transactions[i].TransactionDetails[j].ItemID).Take(&transactions[i].TransactionDetails[j].Item).Error
				if err != nil {
					return &[]Transaction{}, err
				}
			}
		}
	}
	return &transactions, nil
}

func (t *Transaction) FindTransactionByID(db *gorm.DB, pid uint64) (*Transaction, error) {
	var err error
	err = db.Debug().Model(&Transaction{}).Where("id = ?", pid).Take(&p).Error
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

func (t *Transaction) UpdateATransaction(db *gorm.DB) (*Transaction, error) {

	var err error

	err = db.Debug().Model(&Transaction{}).Where("id = ?", t.ID).Updates(Transaction{Date: p.Date}).Error
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

	err = db.Debug().Model(&Transaction{}).Where("id = ?", pid).Updates(Transaction{DeletedAt: time.Now()}).Error
	if err != nil {
		return &Transaction{}, err
	}
	return t, nil
}