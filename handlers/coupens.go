package handlers

import (
	"fmt"
	"github/jaseelaali/orchid/repository"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// ADD COUPEN
// @Summary ADD COUPEN
// @ID add-coupen
// @Description Admin can add coupen here
// @Tags Admin
// @Tags coupen management
// @Accept json
// @Produce json
// @Param code query string true "code"
// @Param minamount query string true "minamount"
// @Param amount query string true "amount"
// @Success 200
// @Failure 400
// @Router /admin/addcoupen [post]
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

// LIST COUPEN
// @Summary LIST COUPEN
// @ID list-coupen
// @Description Admin can list coupen here
// @Tags Admin
// @Tags coupen management
// @Accept json
// @Produce json
// @Success 200
// @Failure 400
// @Router /admin/listcoupen [get]
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

// APPLY COUPEN
// @Summary APPLY COUPEN
// @ID applycoupen
// @Description user can apply coupen here
// @Tags User
// @Tags coupen
// @Accept json
// @Produce json
// @Param coupenname query string true "coupenname"
// @Success 200
// @Failure 400
// @Router /user/applycoupens [post]
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

// LIST COUPEN
// @Summary LIST COUPEN
// @ID listcoupen
// @Description user can list coupen here
// @Tags User
// @Tags coupen
// @Accept json
// @Produce json
// @Success 200
// @Failure 400
// @Router /user/listcoupens [get]
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
