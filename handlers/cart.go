package handlers

import (
	"github/jaseelaali/orchid/database"
	"github/jaseelaali/orchid/repository"
	"strconv"

	"github.com/gin-gonic/gin"
	//"github.com/go-playground/validator/v10/translations/id"
)

// AddCart
// @Summary Add Cart
// @ID add-cart
// @Description User can add items to their cart
// @Tags User
// @Tags cart management
// @Accept json
// @Produce json
// @Param product_id query integer true "Product_ID"
// @Param product_quantity query integer true "Product_Quantity"
// @Success 200
// @Failure 400
// @Router /user/addcart [post]
func AddCart(r *gin.Context) {
	product_id, err := strconv.Atoi(r.Query("product_id"))
	if err != nil {
		r.JSON(400, gin.H{
			"message": "product id not geted",
		})
		return
	}
	product_quantity, err := strconv.Atoi(r.Query("product_quantity"))
	if err != nil {
		r.JSON(400, gin.H{
			"message": "quantity not geted",
		})
		return
	}

	user_id := repository.GetId(r)
	var stock int
	err = database.DB.Raw("SELECT stock FROM products WHERE id=$1;", product_id).Scan(&stock).Error
	if stock == 0 {
		r.JSON(400, gin.H{
			"message": "no product",
		})
		return
	} else if product_quantity > stock {
		r.JSON(400, gin.H{
			"message": "the product is out of stock",
		})
		return
	}
	var alreadyexist int
	database.DB.Raw("SELECT quantity FROM cart_items WHERE product_id=$1 AND user_id=$2;", product_id, user_id).Scan(&alreadyexist)
	if alreadyexist != 0 {
		err = repository.ADDcart(alreadyexist, product_id, product_quantity, user_id)
		if err != nil {
			r.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		sum, err := repository.SumCart(user_id)
		if err != nil {
			r.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		r.JSON(200, gin.H{
			"sum of the products is": sum,
		})

	} else {
		//database.DB.Raw("UPDATE PRODUCTS SET stock=$1 WHERE id=$2;", stock-product_quantity, product_id).Scan(&models.Products{})
		err = repository.Addcart(product_id, product_quantity, user_id)
		if err != nil {
			r.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		sum, err := repository.SumCart(user_id)
		if err != nil {
			r.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		r.JSON(200, gin.H{
			"sum of the products is": sum,
		})
	}
	r.JSON(200, gin.H{
		"message": "product added to cart successfully",
	})
}

// ViewCart
// @Summary View Cart
// @ID view-cart
// @Description User can view their cart
// @Tags User
// @Tags cart management
// @Accept json
// @Produce json
// @Param page query string true "Page "
// @Param perpage query string true "perpage"
// @Success 200
// @Failure 400
// @Router /user/viewcart [get]
func ViewCart(r *gin.Context) {
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
	userid := repository.GetId(r)
	cart, metadata, err := repository.Viewcart(userid, page, perpage)
	if err != nil {
		r.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	r.JSON(200, gin.H{
		"cart":     cart,
		"metadata": metadata})

}

// DELETE CART
// @Summary delete-cart items
// @ID delete-cart
// @Description User delete cart items
// @Tags User
// @Tags cart management
// @Accept json
// @Produce json
// @Param productId query string true "id"
// @Param quantity query string true "quantity"
// @Success 200
// @Failure 400
// @Router /user/deleteitem [delete]
func DeleteItem(r *gin.Context) {
	id, err := strconv.Atoi(r.Query("id"))
	if err != nil {
		r.JSON(400, gin.H{
			"message": "didn't get id",
		})
		return
	}
	quantity, err := strconv.Atoi(r.Query("quantity"))
	if err != nil {
		r.JSON(400, gin.H{
			"message": "didn't get quantity",
		})
		return
	}
	userid := repository.GetId(r)
	err = repository.Deleteitem(id, quantity, userid)
	if err != nil {
		r.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	r.JSON(200, gin.H{
		"message": "product delete from cart  successfully",
	})
}
