package repository

import (
	"github/jaseelaali/orchid/database"
	"github/jaseelaali/orchid/models"
)

func OrderStatus(user_id int, payment_id string) error {
	var product_id []int
	database.DB.Raw("SELECT product_id FROM ordered_products WHERE user_id=$1;", user_id).Scan(&product_id)
	for i := range product_id {
		var productname string
		var quantity, product_price int
		Result := database.DB.Raw("INSERT INTO order_statuses(user_id, product_id) VALUES($1, $2);", user_id, product_id[i])
		if Result.Error != nil {
			return Result.Error
		}
		Result = database.DB.Raw("SELECT product_name FROM products WHERE id=$1;", product_id[i]).Scan(&productname)
		if Result.Error != nil {
			return Result.Error
		}
		Result = database.DB.Raw("SELECT quantity FROM ordered_products WHERE user_id=$1 AND product_id=$2;", user_id, product_id[i]).Scan(&quantity)
		if Result.Error != nil {
			return Result.Error
		}
		Result = database.DB.Raw("SELECT product_price FROM products WHERE id=$1;", product_id[i]).Scan(&product_price)
		if Result.Error != nil {
			return Result.Error
		}

		Result = database.DB.Raw("INSERT INTO order_statuses(user_id, product_name, quantity, product_price, product_id,delivery,payment_id) VALUES ($1, $2, $3, $4, $5,'not done',$6)", user_id, productname, quantity, product_price, product_id[i], payment_id).Scan(&models.OrderStatus{})
		if Result.Error != nil {
			return Result.Error
		}
	}
	return nil
}
func ClearCart(user_id int) error {
	result := database.DB.Raw("DElETE FROM cart_items WHERE user_id=$1;", user_id).Scan(&models.CartItem{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func OrderViewUpdation(user_id int) error {
	var product_id []int
	database.DB.Raw("SELECT product_id FROM cart_items WHERE user_id=$1;", user_id).Scan(&product_id)
	for i := range product_id {
		var productname string
		var quantity, product_price int
		Result := database.DB.Raw("INSERT INTO ordered_products(user_id, product_id) VALUES($1, $2);", user_id, product_id[i]).Scan(&models.OrderedProduct{})
		if Result.Error != nil {
			return Result.Error
		}
		Result = database.DB.Raw("SELECT product_name FROM products WHERE id=$1;", product_id[i]).Scan(&productname)
		if Result.Error != nil {
			return Result.Error
		}
		Result = database.DB.Raw("SELECT quantity FROM cart_items WHERE user_id=$1 AND product_id=$2;", user_id, product_id[i]).Scan(&quantity)
		if Result.Error != nil {
			return Result.Error
		}
		Result = database.DB.Raw("SELECT product_price FROM products WHERE id=$1;", product_id[i]).Scan(&product_price)
		if Result.Error != nil {
			return Result.Error
		}
		Result = database.DB.Raw("UPDATE ordered_products  SET product=$1, quantity=$2, price=$3,total=$4, payment=$5 WHERE user_id=$6 AND product_id=$7  AND  payment is null;", productname, quantity, product_price, quantity*product_price, "not done", user_id, product_id[i]).Scan(&models.OrderedProduct{})
	}
	return nil
}
