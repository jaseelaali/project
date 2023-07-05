package repository

import (
	"errors"
	"fmt"
	"github/jaseelaali/orchid/database"
	"github/jaseelaali/orchid/models"
	"github/jaseelaali/orchid/utils"

	"gorm.io/gorm"
)

func Add_Order(user_id, Address_id int) (int, error) {
	var Payment []models.Order
	var cart int
	result := database.DB.Raw("SELECT payment_status  FROM orders WHERE user_id=$1;", user_id).Scan(&Payment)
	if result.Error != nil {
		return 0, result.Error
	}
	for i := range Payment {
		//payment := fmt.Sprint(Payment)
		if Payment[i].Payment_Status == "not done" {

			return 0, errors.New("one order is  pending")
		}
	}
	result = database.DB.Raw("SELECT SUM(product_total_price) FROM cart_items WHERE user_id=$1 ;", user_id).Scan(&cart)
	if result.Error != nil {
		return 0, result.Error
	}
	//updated
	var order_id int
	result = database.DB.Raw("INSERT INTO orders (user_id,address_id,total_cart_amount,total_amount,payment_status)VALUES(?,?,?,?,?);", user_id, Address_id, cart, cart, "not done").Scan(&models.Order{})
	//result = database.DB.Raw(" SELECT id FROM orders WHERE user_id=$1 AND order_status is null;").Scan(&order_id)
	fmt.Printf("......order..........%v", order_id)
	//result = database.DB.Raw("INSERT INTO orders (total_cart_amount)VALUES(?) WHERE user_id=$2 AND address_id=$3;", cart, user_id, Address_id).Scan(&models.Order{})
	//result = database.DB.Raw("UPDATE Orders SET total_cart_amount=$1 WHERE user_id=$2 AND address_id=$3;", cart, user_id, Address_id).Scan(&models.Order{})
	if result.Error != nil {
		return 0, result.Error
	}
	return cart, nil

}
func OrderUpdation(payment_id string, user_id int) error {
	result := database.DB.Raw("UPDATE orders SET payment_method=$1,payment_status=$2,payment_id=$3,order_status=$4 WHERE user_id=$5 AND order_status is null;", "Razor_pay", "success", payment_id, "success", user_id).Scan(&models.Order{})
	if result.Error != nil {
		return result.Error
	}

	result = database.DB.Raw("UPDATE ordered_products SET payment='done' WHERE user_id=$1 AND payment =$2;", user_id, "not done").Scan(&models.OrderedProduct{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func Show_Order(user_id, page, perPage int) ([]models.OrderedProduct, utils.MetaData, error) {
	value := []models.OrderedProduct{}
	var totalRecords int64
	err := database.DB.Raw("SELECT COUNT(id) FROM ordered_products  WHERE  user_id=?;", user_id).Scan(&totalRecords).Error
	if err != nil {
		return value, utils.MetaData{}, err
	}
	metaData, offset, err := utils.ComputeMetaData(page, perPage, int(totalRecords))
	if err != nil {
		return value, metaData, err
	}
	err = database.DB.Raw("SELECT * FROM ordered_products WHERE user_id=$1 and payment=$2 OFFSET $3 LIMIT $4;", user_id, "not done", offset, perPage).Scan(&value).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return value, metaData, errors.New("Record not found")
		}
		return value, metaData, err
	}
	return value, metaData, nil
}

func Cancel_Order(user_id int) error {
	result := database.DB.Raw("DELETE from ordered_products WHERE user_id=$1 and payment=$2;", user_id, "not done").Scan(&models.OrderedProduct{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
