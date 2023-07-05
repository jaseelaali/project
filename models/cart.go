package models

import (
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	User_id     int
	Cart_id     int `json:"cart_id"`
	Product_id  int
	Total_price int
}
type CartItem struct {
	CartI_Id            int `json:"carti_id" gorm:"unique;primarykey"`
	User_id             int `json:"user_id"`
	Product_Name        string
	Product_Id          int
	Quantity            int
	Product_Price       int
	Product_Total_price int
	Coupen              string
}
