package repository

import "github/jaseelaali/orchid/database"

func Mywallet(user_id int) (int, error) {
	var data int
	result := database.DB.Raw("SELECT money FROM wallets WHERE user_id=$1; ", user_id).Scan(&data)
	if result.Error != nil {
		return 0, result.Error
	}
	return data, nil
}
