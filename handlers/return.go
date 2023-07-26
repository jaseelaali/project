package handlers

import (
	"github/jaseelaali/orchid/repository"

	"github.com/gin-gonic/gin"
)

//ReturnStatus
//@Summary ReturnStatus
//@ID return status
//@Description admin can see delivery status
//@Tags Admin
//@Tags delivery status
//@Accept json
//@Produce json
//@Param paymentid query string true "paymentid"
//@Success 200
//@Failure 400
//@Router /admin/returnstatus [POST]
func ReturnStatus(r *gin.Context) {
	id := r.Query("paymentid")
	if id == "" {
		r.JSON(400, gin.H{
			"message": "didn't get payment id",
		})
		return
	}
	err := repository.ReturnStatusChange(id)
	if err != nil {
		r.JSON(400, gin.H{
			"message": "couldn't update the data",
		})
		return
	}
	r.JSON(200, gin.H{
		"message": "delivery completed",
	})
}
// RETURN PRODUCT
// @Summary RETURN PRODUCT
// @ID return product
// @Description user can return product
//@Tags User
//@Tags return product
// @Accept json
// @Produce json
//@Param  orderid query string true "orderid"
//@Param  productid query string true "productid"
// @Success 200
// @Failure 400
// @Router /user/returnproduct [post]
func ReturnMyProduct(r *gin.Context) {
	orderid := r.Query("orderid")
	if orderid == "" {
		r.JSON(400, gin.H{
			"message": "didn't get order id",
		})
		return
	}
	productid := r.Query("productid")
	if productid == "" {
		r.JSON(400, gin.H{
			"message": "didn't get product id",
		})
		return
	}
	user_id := repository.GetId(r)
	err := repository.ReturnProduct(user_id, orderid, productid)
	if err != nil {
		r.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	r.JSON(200, gin.H{
		"message": "your product collected soon",
	})
}
