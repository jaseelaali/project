package handlers

import (
	"fmt"
	"github/jaseelaali/orchid/repository"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)



func AddCoupens(r *gin.Context) {
	code := r.Query("code")
	if code == "" {
		r.JSON(400, gin.H{
			"message": "code is missing",
		})
		return
	}
	minamount, err := strconv.Atoi(r.Query("minamount"))
	if err != nil {
		r.JSON(400, gin.H{
			"message": "minimum amount is missing",
		})
		return
	}
	amount, err := strconv.Atoi(r.Query("amount"))
	if err != nil {
		r.JSON(400, gin.H{
			"message": "discount amount is missing",
		})
		return
	}
	var expiry time.Time
	expiry = time.Now().Add(time.Hour * 24 * 1)
	err = repository.Addcoupen(code, expiry, minamount, amount)
	if err != nil {
		r.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	r.JSON(200, gin.H{
		"success": "coupen added successfully",
	})
}

func ListCoupen(r *gin.Context) {
	coupen, result := repository.Listcoupen()
	if result != nil {
		r.JSON(400, gin.H{
			"error": result.Error(),
		})
		return
	}
	r.JSON(200, gin.H{
		"coupens": coupen,
	})
}

func ApplyCoupen(r *gin.Context) {
	var body struct {
		Coupenname string `json:"coupenname"`
	}
	err := r.Bind(&body)
	if err != nil {
		r.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	user_id := repository.GetId(r)
	fmt.Printf("+++++++:%v", user_id)
	data, err := repository.Applycoupen(user_id, body.Coupenname)
	if err != nil {
		r.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	r.JSON(200, gin.H{
		"data":    data,
		"success": "coupen apply succesfully",
	})

}

func Listcoupens(r *gin.Context) {
	coupen, result := repository.Listcoupen()
	if result != nil {
		r.JSON(400, gin.H{
			"error": result.Error(),
		})
		return
	}
	r.JSON(200, gin.H{
		"coupens": coupen,
	})
}
