package repository

import (
	"errors"
	"github/jaseelaali/orchid/database"
	"github/jaseelaali/orchid/models"
)

func AddWishlist(user_id, product_id int) error {
	var Count int
	database.DB.Raw("SELECT COUNT(id) FROM products WHERE id=$1;", product_id).Scan(&Count)
	if Count == 0 {
		return errors.New("this product is not available in the product-list")
	}
	var count int
	database.DB.Raw("SELECT COUNT(product_id) FROM wish_lists WHERE user_id=$1 AND product_id=$2  ;", user_id, product_id).Scan(&count)
	if count != 0 {
		return errors.New("this product is already in your wishlist")
	}
	result := database.DB.Raw("INSERT INTO wish_lists(user_id,product_id)VALUES($1,$2);", user_id, product_id).Scan(&models.WishList{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func ViewWishList(user_id int) ([]models.ViewWishList, error) {
	body := []models.ViewWishList{}
	// result := database.DB.Raw("SELECT * FROM wish_lists WHERE user_id=$1;", user_id).Scan(&body)
	// if result.Error != nil {
	// 	return nil, result.Error
	// }
	// for i := range body {
	// 	var name string
	// 	result = database.DB.Raw("SELECT product_name FROM products WHERE id=$1;", body[i].Product_id).Scan(&name)
	// 	body[i].Product_Name = name
	// }

	result := database.DB.Raw(`SELECT w.user_id, w.product_id, p.product_name ,p.product_price FROM wish_lists w
	INNER JOIN products p ON w.product_id = p.id
	WHERE w.user_id = $1;`, user_id).Scan(&body)
	if result.Error != nil {
		return nil, result.Error
	}
	return body, nil

}
func RemoveWishList(user_id, product_id int) error {
	var count int
	database.DB.Raw("SELECT COUNT(product_id) FROM wish_lists WHERE user_id =$1 AND product_id=$2 ;", user_id, product_id).Scan(&count)
	if count == 0 {
		return errors.New("this product is not available in the your wishlist")
	}
	result := database.DB.Raw("DELETE FROM wish_lists WHERE user_id=$1 AND product_id=$2;", user_id, product_id).Scan(&models.WishList{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
