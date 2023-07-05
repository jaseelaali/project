package models

type UserResponses struct {
	ID           uint   `json:"id"`
	User_Name    string `json:"user_name"`
	Email        string `json:"email"`
	Phone_Number string `json:"phone_number"`
}
