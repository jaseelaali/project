package repository

import (
	//"fmt"

	"errors"
	"fmt"
	"github/jaseelaali/orchid/database"
	"github/jaseelaali/orchid/models"
	"github/jaseelaali/orchid/utils"

	"gorm.io/gorm"
)

func Addcart(P_Id, Quantity, U_Id int) error {
	var productname string
	var productprice int
	err := database.DB.Raw("SELECT product_name FROM products WHERE id=$1;", P_Id).Scan(&productname)
	err = database.DB.Raw("SELECT product_price FROM products WHERE id=$1;", P_Id).Scan(&productprice)
	//fmt.Println("************* %v"productname)

	err = database.DB.Create(&models.CartItem{
		User_id:             U_Id,
		Product_Name:        productname,
		Product_Id:          P_Id,
		Quantity:            Quantity,
		Product_Price:       productprice,
		Product_Total_price: (productprice * Quantity),
	}).Where("user_id", U_Id)
	if err != nil {
		return err.Error
	}
	return nil
}
func ADDcart(newquantity, P_Id, Quantity, U_Id int) error {
	var productname string
	var productprice int
	err := database.DB.Raw("SELECT product_name FROM products WHERE id=$1;", P_Id).Scan(&productname)
	err = database.DB.Raw("SELECT product_price FROM products WHERE id=$1;", P_Id).Scan(&productprice)
	//fmt.Println("************* %v"productname)
	var quantity int
	quantity = (newquantity + Quantity)
	price := ((newquantity + Quantity) * productprice)
	//fmt.Println("****pazhe:%v****puth:%v***quq:%v", newquantity, Quantity, quantity)
	err = database.DB.Raw("UPDATE cart_items SET quantity=$1 WHERE user_id=$2 AND product_id=$3;", quantity, U_Id, P_Id).Scan(&models.CartItem{})
	err = database.DB.Raw("UPDATE cart_items SET product_total_price=$1 WHERE product_id=$2  AND user_id=$3;", price, P_Id, U_Id).Scan(&models.CartItem{})
	if err != nil {
		return err.Error
	}
	return nil
}

func Viewcart(ID, page, perPage int) ([]models.CartItem, utils.MetaData, error) {
	Id := ID
	cart := []models.CartItem{}

	var totalRecords int64
	err := database.DB.Raw("SELECT COUNT(user_id) FROM cart_items  WHERE  user_id=?;", Id).Scan(&totalRecords).Error
	if err != nil {
		return cart, utils.MetaData{}, err
	}
	metaData, offset, err := utils.ComputeMetaData(page, perPage, int(totalRecords))
	if err != nil {
		return cart, metaData, err
	}
	//fmt.Println("************* %v", Id)
	err = database.DB.Raw("SELECT * FROM cart_items WHERE  user_id=? OFFSET ? LIMIT ?;", Id, offset, perPage).Scan(&cart).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return cart, metaData, errors.New("Record not found")
		}
		return cart, metaData, err
	}
	return cart, metaData, nil
}
func Deleteitem(Id, Quantity, U_Id int) error {
	var OldQuantity, Price int
	err := database.DB.Raw("SELECT quantity FROM cart_items WHERE user_id=$1 AND product_id=$2;", U_Id, Id).Scan(&OldQuantity)
	err = database.DB.Raw("SELECT product_price FROM cart_items WHERE user_id=$1 AND product_id=$2;", U_Id, Id).Scan(&Price)
	if OldQuantity > 0 {
		quantity := (OldQuantity - Quantity)
		if quantity > 0 {
			price := Price * quantity
			//fmt.Println("****old:%v**kal:%v**new:%v***price:%v", OldQuantity, Quantity, quantity, price)
			err = database.DB.Raw("UPDATE cart_items SET quantity=$1 WHERE user_id=$2 AND product_id=$3;", quantity, U_Id, Id).Scan(&models.CartItem{})
			if err.Error != nil {
				return err.Error
			}
			err = database.DB.Raw("UPDATE cart_items SET product_total_price=$1 WHERE user_id=$2 AND product_id=$3;", price, U_Id, Id).Scan(&models.CartItem{})
			if err.Error != nil {
				return err.Error
			}
		} else if quantity == 0 {
			err = database.DB.Raw("DELETE  FROM cart_items WHERE user_id=$1 AND product_id=$2;", U_Id, Id).Scan(&models.CartItem{})
			if err.Error != nil {
				return err.Error
			}
		} else {
			return errors.New("check you products count")
		}
		if err.Error != nil {
			return err.Error
		}
		return nil
	} else {
		return errors.New("no item present")
	}
}

/*
	func ADCart(product_id, user_id int) error {
		result := database.DB.Raw("SELECT * from carts where user_id=$1;", user_id).Scan(&models.Cart{})
		if result.Error != nil {
			//cart_items := &models.CartItem{}
			var price int
			result = database.DB.Raw("SELECT product_total_price from cart_items where user_id=$1 AND product_id=$2;", user_id, product_id).Scan(&price)
			fmt.Println("******")
			fmt.Println("\\\\\\\\\\\\", price)
		}
		return nil
	}
*/

func SumCart(user_id int) (int, error) {

	var sum int
	result := database.DB.Raw("SELECT SUM(product_total_price) FROM cart_items WHERE user_id=$1;", user_id).Scan(&sum)
	fmt.Println("\n\nthe sum is :", sum, user_id)
	if result != nil {
		return sum, result.Error
	}
	return sum, nil
}
func Sum(user_id int) (int, error) {
	var sum int
	result := database.DB.Raw("SELECT total_amount FROM orders WHERE user_id=$1;", user_id).Scan(&sum)
	fmt.Println("\n\the sum is :", sum, user_id)
	if result != nil {
		return sum, result.Error
	}
	return sum, nil
}
