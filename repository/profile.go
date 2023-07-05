package repository

import (
	"github/jaseelaali/orchid/database"
	"github/jaseelaali/orchid/models"
)

func ViewProfile(user_id int) (models.UserProfile, error) {
	profile := models.UserProfile{}
	err := database.DB.Raw("SELECT * FROM users WHERE id=$1", user_id).Scan(&profile)
	if err != nil {
		return profile, err.Error
	}
	return profile, nil
}
func EditProfileEmail(email string, user_id int) error {
	err := database.DB.Raw("UPDATE users SET email=$1 WHERE id=$2", email, user_id).Scan(&models.User{})
	if err != nil {
		return err.Error
	}
	return nil
}
func EditProfilePhoneNumber(number string, user_id int) error {
	err := database.DB.Raw("UPDATE users SET phone_number=$1 WHERE id=$2", number, user_id).Scan(&models.User{})
	if err != nil {
		return err.Error
	}
	return nil
}
func EditProfileUserName(name string, user_id int) error {
	err := database.DB.Raw("UPDATE users SET user_name=$1 WHERE id=$2", name, user_id).Scan(&models.User{})
	if err != nil {
		return err.Error
	}
	return nil
}
func Deleteprofile(user_id int)error{
	err:=database.DB.Raw("DELETE FROM users WHERE id=$1",user_id).Scan(&models.User{})
	if err!=nil{
		return err.Error
	}
	return nil
}