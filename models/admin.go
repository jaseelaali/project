package models

import "gorm.io/gorm"

// type SuperAdmin struct {
// 	gorm.Model
// 	Admin_Name string
// 	Email      string
// 	Password   string
// }
type User struct {
	gorm.Model
	User_Name    string `json:"user_name" gorm:"not null"`
	Email        string `json:"email" binding:"required,email" gorm:"not null;unique"`
	Phone_Number string `json:"phone_number" binding:"required,numeric,len=10" gorm:"not null;unique"`
	Password     string `json:"password" binding:"required,min=6" gorm:"not null;unique"`
	Status       string `json:"status"`
}
type UserProfile struct {
	User_Name    string
	Email        string
	Phone_Number string
}
type Admin struct {
	gorm.Model
	Admin_Name string
	Email      string
	Password   string
}
