package models

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	User_id    int    `json:"user_id" gorm:"FOREGIN KEY"`
	House_name string `json:"house_name" binding:"required"`
	Place      string `json:"place" binding:"required"`
	District   string
	State      string
	Pin_Number int
}
