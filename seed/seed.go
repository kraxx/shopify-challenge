/*
** A demo needs sample data!
** We include an option to reseed, if things ever get messy.
 */

package seed

import (
	"github.com/jinzhu/gorm"
	"github.com/kraxx/shopify-challenge/models"
)

var shops = []models.Shop{
	models.Shop{Name: "Wacky Wizardry"},
	models.Shop{Name: "Fancy Frills"},
	models.Shop{Name: "Thrasher Thrills"},
	models.Shop{Name: "Doug's Dirt Shop"},
	models.Shop{Name: "Smiles For Kyle"},
}

var products = []models.Product{
	models.Product{ShopID: 1, Name: "Wooden Rod", Value: 1000, Quantity: 24},
	models.Product{ShopID: 1, Name: "Wizard Hat", Value: 2500, Quantity: 21},
	models.Product{ShopID: 1, Name: "Magic Book", Value: 9999, Quantity: 2},
	models.Product{ShopID: 1, Name: "Cheap Rune", Value: 450, Quantity: 202},
	models.Product{ShopID: 1, Name: "Flying Broom", Value: 860, Quantity: 14},
	models.Product{ShopID: 2, Name: "Thick Dress", Value: 100, Quantity: 56},
	models.Product{ShopID: 2, Name: "Cute Skirt", Value: 20, Quantity: 256},
	models.Product{ShopID: 2, Name: "Night Gown", Value: 35, Quantity: 67},
	models.Product{ShopID: 2, Name: "Generic Cloth", Value: 5, Quantity: 21},
	models.Product{ShopID: 3, Name: "Razor", Value: 7, Quantity: 601},
	models.Product{ShopID: 3, Name: "Fireworks", Value: 100, Quantity: 43},
	models.Product{ShopID: 3, Name: "Lazer Beams", Value: 10000, Quantity: 1000},
	models.Product{ShopID: 3, Name: "A Bomb", Value: 900000, Quantity: 1},
	models.Product{ShopID: 4, Name: "Fine Soil", Value: 10, Quantity: 501229},
	models.Product{ShopID: 4, Name: "Coarse Dirt", Value: 13, Quantity: 30129},
	models.Product{ShopID: 4, Name: "Wet Pile", Value: 60, Quantity: 28322},
	models.Product{ShopID: 4, Name: "Dry Dirt", Value: 13, Quantity: 7222219},
	models.Product{ShopID: 5, Name: "One Wide Smile", Value: 1, Quantity: 111111},
}

var orders = []models.Order{
	models.Order{ShopID: 1},
	models.Order{ShopID: 1},
	models.Order{ShopID: 1},
	models.Order{ShopID: 2},
	models.Order{ShopID: 2},
	models.Order{ShopID: 2},
	models.Order{ShopID: 3},
	models.Order{ShopID: 3},
	models.Order{ShopID: 4},
	models.Order{ShopID: 5},
}

var lineItems = []models.LineItem{
	models.LineItem{OrderID: 1, ProductID: 1, Quantity: 2},
	models.LineItem{OrderID: 1, ProductID: 2, Quantity: 12},
	models.LineItem{OrderID: 2, ProductID: 1, Quantity: 5},
	models.LineItem{OrderID: 2, ProductID: 2, Quantity: 1},
	models.LineItem{OrderID: 2, ProductID: 4, Quantity: 1},
	models.LineItem{OrderID: 3, ProductID: 1, Quantity: 1},
	models.LineItem{OrderID: 3, ProductID: 2, Quantity: 1},
	models.LineItem{OrderID: 3, ProductID: 3, Quantity: 1},
	models.LineItem{OrderID: 3, ProductID: 4, Quantity: 1},
	models.LineItem{OrderID: 3, ProductID: 5, Quantity: 1},
	models.LineItem{OrderID: 4, ProductID: 6, Quantity: 2},
	models.LineItem{OrderID: 4, ProductID: 7, Quantity: 5},
	models.LineItem{OrderID: 5, ProductID: 6, Quantity: 12},
	models.LineItem{OrderID: 5, ProductID: 7, Quantity: 155},
	models.LineItem{OrderID: 5, ProductID: 8, Quantity: 60},
	models.LineItem{OrderID: 5, ProductID: 9, Quantity: 20},
	models.LineItem{OrderID: 5, ProductID: 10, Quantity: 12},
	models.LineItem{OrderID: 6, ProductID: 9, Quantity: 32},
	models.LineItem{OrderID: 7, ProductID: 11, Quantity: 320},
	models.LineItem{OrderID: 7, ProductID: 12, Quantity: 1},
	models.LineItem{OrderID: 8, ProductID: 11, Quantity: 100},
	models.LineItem{OrderID: 8, ProductID: 12, Quantity: 638},
	models.LineItem{OrderID: 8, ProductID: 13, Quantity: 1},
	models.LineItem{OrderID: 8, ProductID: 14, Quantity: 60},
	models.LineItem{OrderID: 9, ProductID: 14, Quantity: 9999},
	models.LineItem{OrderID: 9, ProductID: 15, Quantity: 9999},
	models.LineItem{OrderID: 9, ProductID: 16, Quantity: 9999},
	models.LineItem{OrderID: 9, ProductID: 17, Quantity: 9999},
	models.LineItem{OrderID: 10, ProductID: 18, Quantity: 1},
}

func SeedData(db *gorm.DB) {
	for _, shop := range shops {
		db.Create(&shop)
	}
	for _, product := range products {
		db.Create(&product)
	}
	for _, order := range orders {
		db.Create(&order)
	}
	for _, lineItem := range lineItems {
		db.Create(&lineItem)
	}
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
