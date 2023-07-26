package handlers

import (
	"github/jaseelaali/orchid/repository"
	"strconv"

	"github.com/gin-gonic/gin"
)

// My Wallet
// @Summary  My Wallet
// @ID  My_Wallet
// @Description user can view wallet
// @Tags User
// @Tags  My Wallet
// @Accept json
// @Produce json
// @Success 200
// @Failure 400
// @Router /user/mywallet [get]
func MyWallet(r *gin.Context) {
	user_id := repository.GetId(r)
	data, err := repository.Mywallet(user_id)
	if err != nil {
		r.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	r.JSON(200, gin.H{
		"message": "your balance is " + strconv.Itoa(data),
	})
}
