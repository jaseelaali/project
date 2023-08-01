package repository

import (
	"errors"
	"fmt"
	"github/jaseelaali/orchid/database"
	"github/jaseelaali/orchid/models"

	"golang.org/x/crypto/bcrypt"
)

/*import (
	"math/rand"
	"strconv"
)*/

func Forget(email, phonenumber, password string) error {
	fmt.Println("********:", email, phonenumber)
	newpassword, err := (bcrypt.GenerateFromPassword([]byte(password), 11))
	if err != nil {
		return err
	}
	var num string
	err = database.DB.Raw("select phone_number from users where email=$1;",email).Scan(&num).Error
	if err!=nil{
		return err
	}
	if num != phonenumber{
		return errors.New("information are not matched")
	}
	err = database.DB.Raw("update users set password=$1 where email=$2 and phone_number =$3; ", newpassword, email, phonenumber).Scan(&models.User{}).Error
	return err
}
