package handlers

import (
	"github/jaseelaali/orchid/repository"
	"strconv"

	"github.com/gin-gonic/gin"
)

// MYWALLET
// @summary My wallet
// @ID wallet
// @Description  	Use can check the amount in wallet
// @Tags User
//@Tags my wallet
// @Produce json
// @Sucess 200
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
