package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	User_Id int
	//Order_Id       int
	Address_Id      int
	TotalCartAmount int
	Coupen          string
	Coupen_Name     string
	Discount        int
	Total_Amount    int
	Payment_Method  string
	Payment_Status  string
	Payment_Id      string
	Order_Status    string
}
type OrderStatus struct {
	gorm.Model
	User_id       int
	Product_id    int
	order_id      int
	Product_name  string
	Quantity      int
	Product_Price int
	Payment_Id    string
	Delivery      string
}
type OrderedProduct struct {
	gorm.Model
	User_id    int
	Product_id int
	Product    string
	Quantity   int
	Price      int
	Total      int
	Payment    string
}
