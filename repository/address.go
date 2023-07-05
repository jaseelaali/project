package repository

import (
	"errors"
	"fmt"
	"github/jaseelaali/orchid/database"
	"github/jaseelaali/orchid/models"
)

func AddAddress(address models.Address, user_id int) error {
	fmt.Printf("*****%v*****++", user_id)
	Address := models.Address{}
	err := database.DB.Raw("INSERT INTO addresses(user_id,house_name,place,district,state,pin_number)VALUES($1,$2,$3,$4,$5,$6);", user_id, address.House_name, address.Place,
		address.District, address.State, address.Pin_Number).Scan(&Address).Error
	if err != nil {
		return err
	}
	return nil
}
func EditHousename(name string, Id int) error {
	err := database.DB.Raw("UPDATE addresses SET house_name=$1 WHERE id =$2;", name, Id).Scan(&models.Address{})
	return err.Error
}
func EditPlace(place string, Id int) error {
	err := database.DB.Raw("UPDATE addresses SET place=$1 WHERE id =$2;", place, Id).Scan(&models.Address{})
	return err.Error
}
func EditDistrict(district string, Id int) error {
	err := database.DB.Raw("UPDATE addresses SET district=$1 WHERE id =$2;", district, Id).Scan(&models.Address{})
	return err.Error
}
func EditState(state string, Id int) error {
	err := database.DB.Raw("UPDATE addresses SET state=$1 WHERE id =$2;", state, Id).Scan(&models.Address{})
	return err.Error
}
func EditPin(pin, Id int) error {
	err := database.DB.Raw("UPDATE addresses SET pin_number=$1 WHERE id =$2;", pin, Id).Scan(&models.Address{})
	return err.Error
}

func Deleteaddress(User_id, Address_id int) error {
	var exist int
	err := database.DB.Raw("SELECT COUNT(id) FROM addresses WHERE user_id=$1 AND id=$2;", User_id, Address_id).Scan(&exist)
	if exist == 0 {
		return errors.New("this address not occured")
	}
	err = database.DB.Raw("DELETE FROM addresses WHERE user_id=$1 AND id=$2", User_id, Address_id).Scan(&models.Address{})
	if err != nil {
		return err.Error
	}
	return nil
}
func Viewaddress(user_id int) ([]models.Address, error) {
	address := []models.Address{}
	fmt.Println("*********", address)

	err := database.DB.Raw("SELECT * FROM addresses WHERE user_id=$1", user_id).Scan(&address)
	fmt.Println("*********", address)
	if err != nil {
		return address, err.Error
	}

	if address == nil {
		return nil, errors.New("add your address")
	}
	return address, nil
}
