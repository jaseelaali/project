package handlers

import (
	"fmt"
	"github/jaseelaali/orchid/database"
	"github/jaseelaali/orchid/repository"
	"strconv"

	"github.com/gin-gonic/gin"
)


func AddOrder(r *gin.Context) {
	user_id := repository.GetId(r)
	var Address_id int
	result := database.DB.Raw("SELECT id FROM addresses WHERE user_id=$1;", user_id).Scan(&Address_id)
	if result.Error != nil {
		return
	}
	if Address_id == 0 {
		r.JSON(400, gin.H{
			"error": "enter your address",
		})
		return
	}

	Data, err := repository.Add_Order(user_id, Address_id)
	//fmt.Print(o_id)
	if err != nil {
		r.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	data1 := "your total price is:" + fmt.Sprint(Data)
	r.JSON(200, gin.H{
		"data":    data1,
		"success": "placed order successfully",
	})
	err = repository.OrderViewUpdation(user_id)
	if err != nil {
		r.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	repository.ClearCart(user_id)
}

func ShowOrder(r *gin.Context) {
	page, err := strconv.Atoi(r.Query("page"))
	perpage, err := strconv.Atoi(r.Query("perpage"))
	if err != nil {
		r.JSON(400, gin.H{
			"message": "page not geted",
		})
		return
	}
	if err != nil {
		r.JSON(400, gin.H{
			"message": "perpage not geted",
		})
		return
	}
	user_id := repository.GetId(r)
	order, metadata, err := repository.Show_Order(user_id, page, perpage)
	if err != nil {
		r.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	r.JSON(200, gin.H{
		"success":  order,
		"metadata": metadata,
	})
}

func CancelOrder(r *gin.Context) {
	user_id := repository.GetId(r)
	err := repository.Cancel_Order(user_id)
	if err != nil {
		r.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	r.JSON(200, gin.H{
		"success": "order deleted",
	})
}
