package repository

import (
	"fmt"
	"github/jaseelaali/orchid/database"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func GetId(r *gin.Context) int {
	temp := fmt.Sprint(r.Get("user_id"))
	id := strings.Split(temp, " ")
	Id, _ := strconv.Atoi(id[0])
	return Id
}
func GenerateID() string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	id := make([]rune, 10)
	for i := range id {
		id[i] = chars[rand.Intn(len(chars))]
	}
	return string(id)
}
func Price(id, quantity int) (int, error) {
	var price int
	err := database.DB.Raw("SELECT product_price FROM products WHERE id=$1;", id).Scan(&price)
	total := price * quantity
	if err != nil {
		return total, err.Error
	}
	return total, nil
}
