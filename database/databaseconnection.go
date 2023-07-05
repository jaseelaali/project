package database

import (
	"fmt"
	"github/jaseelaali/orchid/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseConnection() {
	// load database
	if err := godotenv.Load(); err != nil {
		fmt.Println("error in loading env file")
	}
	// connect database

	dsn := os.Getenv("DB")
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("error in connecting database")
	}
	log.Println("Successfully connected to database")
	//sync database
	err = DB.AutoMigrate(
		//&models.SuperAdmin{},
		&models.Admin{},
		&models.User{},
		&models.UserResponses{},
		&models.Category{},
		&models.SubCategory{},
		&models.Products{},
		&models.Cart{},
		&models.CartItem{},
		&models.Address{},
		&models.Payment{},
		&models.Order{},
		&models.OrderStatus{},
		&models.Coupen{},
		&models.OrderedProduct{},
		&models.WishList{},
		&models.ViewWishList{},
		&models.Wallet{},
	)
	if err != nil {
		log.Println("error in syncing the database")
	}

}
