package seed

import (
	"../models"
	"github.com/jinzhu/gorm"
)

func SeedData(db *gorm.DB) {
	shop1 := models.Shop{Name: "Wacky Wizardry"}
	shop2 := models.Shop{Name: "Fancy Frills"}
	shop3 := models.Shop{Name: "Thrasher Thrills"}
	shop4 := models.Shop{Name: "Doug's Dirt Shop"}
	shop5 := models.Shop{Name: "Smiles For Kyle"}
	product1_1 := models.Product{ShopID: 1, Name: "Wooden Rod", Value: 1000, Quantity: 24}
	product1_2 := models.Product{ShopID: 1, Name: "Wizard Hat", Value: 2500, Quantity: 21}
	product1_3 := models.Product{ShopID: 1, Name: "Pure Magic", Value: 9999, Quantity: 2}
	product2_1 := models.Product{ShopID: 2, Name: "Thick Dress", Value: 100, Quantity: 56}
	product2_2 := models.Product{ShopID: 2, Name: "Thin Dress", Value: 20, Quantity: 256}
	product3_1 := models.Product{ShopID: 3, Name: "Gun", Value: 10000, Quantity: 1}
	product3_2 := models.Product{ShopID: 3, Name: "Fireworks", Value: 100, Quantity: 10}
	product3_3 := models.Product{ShopID: 3, Name: "A bomb", Value: 900000, Quantity: 1}
	product4_1 := models.Product{ShopID: 4, Name: "Fine soil", Value: 10, Quantity: 1029}
	product4_2 := models.Product{ShopID: 4, Name: "Coarse soil", Value: 13, Quantity: 10129}
	product4_3 := models.Product{ShopID: 4, Name: "Wet soil", Value: 60, Quantity: 102}
	product4_4 := models.Product{ShopID: 4, Name: "Dry soil", Value: 13, Quantity: 1022219}
	product5_1 := models.Product{ShopID: 5, Name: "One wide smile", Value: 1, Quantity: 111111}

	order1_1 := models.Order{ShopID: 1}
	lineItem1_1 := models.LineItem{OrderID: 1, ProductID: 1, Quantity: 2}
	lineItem1_2 := models.LineItem{OrderID: 1, ProductID: 2, Quantity: 20}
	db.Create(&shop1)
	db.Create(&shop2)
	db.Create(&shop3)
	db.Create(&shop4)
	db.Create(&shop5)
	db.Create(&product1_1)
	db.Create(&product1_2)
	db.Create(&product1_3)
	db.Create(&product2_1)
	db.Create(&product2_2)
	db.Create(&product3_1)
	db.Create(&product3_2)
	db.Create(&product3_3)
	db.Create(&product4_1)
	db.Create(&product4_2)
	db.Create(&product4_3)
	db.Create(&product4_4)
	db.Create(&product5_1)
	db.Create(&order1_1)
	db.Create(&lineItem1_1)
	db.Create(&lineItem1_2)
}

func DropAndReseedData(db *gorm.DB) {
	if db.HasTable("shops") {
		db.DropTable(&models.Shop{})
	}
	if db.HasTable("products") {
		db.DropTable(&models.Product{})
	}
	if db.HasTable("orders") {
		db.DropTable(&models.Order{})
	}
	if db.HasTable("line_items") {
		db.DropTable(&models.LineItem{})
	}
	db.AutoMigrate(&models.Shop{}, &models.Product{}, &models.Order{}, &models.LineItem{})
	SeedData(db)
}
