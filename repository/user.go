package repository

import (
	"errors"
	"fmt"
	"github/jaseelaali/orchid/database"
	"github/jaseelaali/orchid/models"
	utils "github/jaseelaali/orchid/utils"

	"gorm.io/gorm"
	//"golang.org/x/crypto/bcrypt"
)

func CreateUser(newUser models.User) error {
	user := models.User{}
	result := database.DB.Raw("INSERT INTO users(user_name,email,phone_number,password,status) VALUES($1,$2,$3,$4,$5);",
		newUser.User_Name, newUser.Email, newUser.Phone_Number, newUser.Password, newUser.Status).Scan(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil

}


func View(page, perPage int) ([]models.UserResponses, utils.MetaData, error) {
	users := []models.UserResponses{}
	var totalRecords int64
	err := database.DB.Raw("SELECT COUNT(id) FROM users;").Scan(&totalRecords).Error
	if err != nil {
		return users, utils.MetaData{}, err
	}
	metaData, offset, err := utils.ComputeMetaData(page, perPage, int(totalRecords))
	if err != nil {
		return users, metaData, err
	}

	err = database.DB.Raw("SELECT * FROM users OFFSET ? LIMIT ?;", offset, perPage).Scan(&users).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return users, metaData, errors.New("Record not found")
		}
		return users, metaData, err
	}

	return users, metaData, nil
}

func BlockUser(user_id int) error {
	//user_id = user_id
	var status string
	database.DB.Raw("SELECT status FROM users WHERE id=$1;", user_id).Scan(&status)

	fmt.Println(user_id)

	if status == "blocked" {
		return errors.New("Selected user is already blocked")
	}
	err := database.DB.Raw("UPDATE users SET status=$1 WHERE id=$2;", "blocked", user_id).Scan(&models.User{})
	if err.Error != nil {
		return err.Error
	}
	return nil
}
func UnBlockUser(user_id int) error {
	var status string
	database.DB.Raw("SELECT status FROM users WHERE id=$1;", user_id).Scan(&status)
	if status == "active" {
		return errors.New("selected user already active")
	}
	err := database.DB.Raw("UPDATE users SET status=$1 WHERE id=$2", "active", user_id).Scan(&models.User{})
	if err.Error != nil {
		return err.Error
	}
	return nil
}
func BlocUsers(page, perPage int) ([]models.UserResponses, utils.MetaData, error) {
	users := []models.UserResponses{}
	var totalRecords int64
	err := database.DB.Raw("SELECT COUNT(id) FROM users;").Scan(&totalRecords).Error
	if err != nil {
		return users, utils.MetaData{}, err
	}
	metaData, offset, err := utils.ComputeMetaData(page, perPage, int(totalRecords))
	if err != nil {
		return users, metaData, err
	}
	err = database.DB.Raw("SELECT * FROM users WHERE status='blocked' OFFSET ? LIMIT ?;", offset, perPage).Scan(&users).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return users, metaData, errors.New("Record not found")
		}
		return users, metaData, err
	}
	return users, metaData, nil
}
func ActiveUser(page, perPage int) ([]models.UserResponses, utils.MetaData, error) {
	users := []models.UserResponses{}
	var totalRecords int64
	err := database.DB.Raw("SELECT COUNT(id) FROM users;").Scan(&totalRecords).Error
	if err != nil {
		return users, utils.MetaData{}, err
	}
	metaData, offset, err := utils.ComputeMetaData(page, perPage, int(totalRecords))
	if err != nil {
		return users, metaData, err
	}
	err = database.DB.Raw("SELECT * FROM users WHERE status='active'OFFSET ? LIMIT ?;", offset, perPage).Scan(&users).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return users, metaData, errors.New("Record not found")
		}
		return users, metaData, err
	}
	return users, metaData, nil
}

// func NewPassword(password string, email string) error {
// 	newpassword := password
// 	user := email
// 	fmt.Println("++++++")
// 	fmt.Println(password)
// 	fmt.Println("+++++")
// 	//fmt.Println(Id)
// 	fmt.Println("++++++")
// 	Password, err := (bcrypt.GenerateFromPassword([]byte(newpassword), 11))
// 	database.DB.Raw("UPDATE password SET password=$1 WHERE id=$2 ", Password, user).Scan(&models.User{})
// 	if err != nil {
// 		return err
// 	}
// 	return nil

// }
func UserById(id int) ([]models.UserResponses, error) {
	body := []models.UserResponses{}
	fmt.Println("...................", id)
	result := database.DB.Raw("SELECT * FROM users WHERE id=$1;", id).Scan(&body)
	if result.Error != nil {
		return nil, result.Error
	}
	fmt.Println(body)
	return body, nil

}
func UserByName(name string) ([]models.UserResponses, error) {
	body := []models.UserResponses{}
	result := database.DB.Raw("SELECT * FROM users WHERE user_name=$1;", name).Scan(&body)
	if result.Error != nil {
		return nil, result.Error
	}
	return body, nil
}
func UserByEmail(Email string) ([]models.UserResponses, error) {
	body := []models.UserResponses{}
	result := database.DB.Raw("SELECT * FROM users WHERE email=$1;", Email).Scan(&body)
	if result.Error != nil {
		return nil, result.Error
	}
	return body, nil
}
func UserByNumber(number string) ([]models.UserResponses, error) {
	body := []models.UserResponses{}
	result := database.DB.Raw("SELECT * FROM users WHERE email=$1;", number).Scan(&body)
	if result.Error != nil {
		return nil, result.Error
	}
	return body, nil
}
