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
	result := database.DB.Raw("SELECT * FROM wish_lists WHERE user_id=$1;", user_id).Scan(&body)
	if result.Error != nil {
		return nil, result.Error
	}
	for i := range body {
		var name string
		result = database.DB.Raw("SELECT product_name FROM products WHERE id=$1;", body[i].Product_id).Scan(&name)
		//result = database.DB.Raw("INSERT INTO view_wish_lists(product_name)VALUES($1)WHERE product_id=$2 AND user_id=$3;", name, body[i].Product_id, user_id).Scan(&models.ViewWishList{})
		result := database.DB.Exec("UPDATE view_wish_lists SET product_name = $1 WHERE product_id = $2 AND user_id = $3", name, body[i].Product_id, user_id)
		if result.Error != nil {
			return nil, result.Error
		}
		result = database.DB.Raw("SELECT * FROM view_wish_lists WHERE user_id=$1;", user_id).Scan(&body)
		if result.Error != nil {
			return nil, result.Error
		}
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
