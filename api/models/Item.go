package models

import (
	"errors"
	"html"
	"strings"
	"time"
	"math"

	"github.com/jinzhu/gorm"
)

type Item struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Name     string    `gorm:"size:255;not null;unique" json:"name"`
	Price   float64    `gorm:"not null;" json:"price"`
	Cost   float64    `gorm:"not null;" json:"cost"`
	DeletedOn time.Time
}

func (i *Item) Prepare() {
	i.ID = 0
	i.Name = html.EscapeString(strings.TrimSpace(i.Name))
	i.Price = math.Round(i.Price)
	i.Cost = math.Round(i.Cost)
}

func (i *Item) Validate() error {

	if i.Name == "" {
		return errors.New("Required Name")
	}
	if i.Price < 0 {
		return errors.New("Required Price")
	}
	if i.Cost < 0 {
		return errors.New("Required Cost")
	}
	return nil
}

func (i *Item) SaveItem(db *gorm.DB) (*Item, error) {
	var err error
	err = db.Debug().Model(&Item{}).Create(&i).Error
	if err != nil {
		return &Item{}, err
	}
	return i, nil
}

func (i *Item) FindAllItems(db *gorm.DB) (*[]Item, error) {
	var err error
	items := []Item{}
	err = db.Debug().Model(&Item{}).Limit(100).Find(&items).Error
	if err != nil {
		return &[]Item{}, err
	}
	return &items, nil
}

func (i *Item) FindItemByID(db *gorm.DB, pid uint64) (*Item, error) {
	var err error
	err = db.Debug().Model(Item{}).Where("id = ?", pid).Take(&i).Error
	if err != nil {
		return &Item{}, err
	}
	return i, nil
}

func (i *Item) UpdateAnItem(db *gorm.DB, pid uint64) (*Item, error) {

	var err error
	i.ID = pid
	err = db.Debug().Model(&Item{}).Where("id = ?", i.ID).Updates(Item{Name: i.Name, Cost: i.Cost, Price: i.Price}).Error
	if err != nil {
		return &Item{}, err
	}
	return i, nil
}

func (i *Item) DeleteAnItem(db *gorm.DB, pid uint64) (*Item, error) {

	var err error
	i.ID = pid
	err = db.Debug().Model(&Item{}).Where("id = ?", i.ID).Updates(Item{DeletedOn: time.Now()}).Error
	if err != nil {
		return &Item{}, err
	}
	return i, nil
}