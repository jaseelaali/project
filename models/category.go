package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Category_Name string `json:"category_name" gorm:"not null;primary key;unique"`
}
type SubCategory struct {
	gorm.Model
	//Category_Name    string `json:"category_name" gorm:"foregin key"`
	SubCategory_Name string `json:"subcategory_name" gorm:"not null;primary key" `
}
type Products struct {
	gorm.Model
	Product_Name        string `json:"product_name" binding:"required"`
	//Product_Category    string `json:"product_category" binding:"required"`
	//Product_Subcategory string `json:"product_sub_category" binding:"required"`
	Product_Colour      string `json:"product_colour" binding:"required"`
	Product_Size        int    `json:"product_size" binding:"required"`
	Product_Brand       string `json:"product_brand" binding:"required"`
	Product_Price       int
	Stock               int
}
