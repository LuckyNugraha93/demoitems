package seed

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/LuckyNugraha93/demoitems/api/models"
)

var users = []models.User{
	models.User{
		Username: "user",
		Password: "password",
	},
}

var items = []models.Item{
	models.Item{
		Name: "item1",
		Price: 100000,
		Cost: 95000,
	},
	models.Item{
		Name: "item2",
		Price: 50000,
		Cost: 45000,
	},
}

var transactions = []models.Transaction{
	models.Transaction{
		ID: 1,
		Number: "202210-1",
		Date: time.Now(),
		PriceTotal: 150000,
		CostTotal: 140000,
	},
}

var transactionDetails = []models.TransactionDetail{
	models.TransactionDetail{
		ID: 1,
		TransactionID: 1,
		ItemQuantity: 1,
		ItemPrice: 100000,
		ItemCost: 95000,
	},
	models.TransactionDetail{
		ID: 2,
		TransactionID: 1,
		ItemQuantity: 1,
		ItemPrice: 50000,
		ItemCost: 45000,
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.User{},&models.TransactionDetail{},&models.Item{},&models.Transaction{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{},&models.Item{},&models.Transaction{},&models.TransactionDetail{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = db.Debug().Model(&models.TransactionDetail{}).AddForeignKey("item_id", "items(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}

		err = db.Debug().Model(&models.Transaction{}).Create(&transactions[i]).Error
		if err != nil {
			log.Fatalf("cannot seed transactions table: %v", err)
		}
	}

	for i, _ := range items {
		err = db.Debug().Model(&models.Item{}).Create(&items[i]).Error
		if err != nil {
			log.Fatalf("cannot seed items table: %v", err)
		}
		transactionDetails[i].ItemID = items[i].ID

		err = db.Debug().Model(&models.TransactionDetail{}).Create(&transactionDetails[i]).Error
		if err != nil {
			log.Fatalf("cannot seed transaction_details table: %v", err)
		}
	}
}