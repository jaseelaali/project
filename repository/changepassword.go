package repository

import (
	"github/jaseelaali/orchid/database"
	"github/jaseelaali/orchid/models"

	"golang.org/x/crypto/bcrypt"
)

/*import (
	"math/rand"
	"strconv"
)*/

func Forget(emai, phonenumber, password string) error {
	newpassword, err := (bcrypt.GenerateFromPassword([]byte(password), 11))
	if err != nil {
		return err
	}
	err = database.DB.Raw("update users set password=$1 where email=$2 and phone_number =$2; ", newpassword, emai, phonenumber).Scan(&models.User{}).Error
	return err
}
