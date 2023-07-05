package handlers

import (
	"github/jaseelaali/orchid/repository"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AddWishList
// @Summary Add a product to the wishlist
// @ID addwishlist
// @Description Adds a product to the user's wishlist based on the provided ID.
// @Tags User
// @Tags Wishlist
// @Accept json
// @Produce json
// @Param id query int true "ID of the product to add to the wishlist"
// @Success 200
// @Failure 400
// @Router /user/addwishlist [post]

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

// ListWishlist gets the user's wishlist.
// @Summary Get the user's wishlist
// @Description Retrieves the wishlist for the authenticated user.
//@Tags User
// @Tags Wishlist
// @Accept json
// @Produce json

// @Success 200
// @Failure 400
// @Router /user/listwishlist [get]
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

// REMOVE WISHLIST
// @Summary REMOVE WISHLIST
// @ID removewishlist
// @Description user can delete items of their wishlist
// @Tags User
// @Tags wishlist
// @Accept json
// @Produce json
// @Success 200
// @Failure 400
// @Router /user/removewishlist [delete]
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
