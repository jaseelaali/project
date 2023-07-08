package handlers

import (
	"github/jaseelaali/orchid/repository"
	"strconv"

	"github.com/gin-gonic/gin"
)



func AddWishList(r *gin.Context) {
	id, err := strconv.Atoi(r.Query("id"))
	if err != nil {
		r.JSON(400, gin.H{
			"message": "didn't get id",
		})
		return
	}
	user_id := repository.GetId(r)
	err = repository.AddWishlist(user_id, id)
	if err != nil {
		r.JSON(400, gin.H{
			"errror": err.Error(),
		})
		return
	}
	r.JSON(200, gin.H{
		"success": "product added to wishlist",
	})
}


func ListWishlist(r *gin.Context) {
	user_id := repository.GetId(r)
	list, err := repository.ViewWishList(user_id)
	if err != nil {
		r.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	r.JSON(200, gin.H{
		"WishList": list,
	})
}


func RemoveWishlist(r *gin.Context) {
	id, err := strconv.Atoi(r.Query("id"))
	if err != nil {
		r.JSON(400, gin.H{
			"message": "didn't get id",
		})
		return
	}
	user_id := repository.GetId(r)
	err = repository.RemoveWishList(user_id, id)
	if err != nil {
		r.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	r.JSON(400, gin.H{
		"success": "product delete from user wishlist",
	})
}
