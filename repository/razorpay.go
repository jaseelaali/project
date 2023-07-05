package repository

import "github/jaseelaali/orchid/database"

/*import "github/jaseelaali/orchid/models"

func RazorPay(paymentid, orderid string,home interface{}) error {
	payment:=models.Payment{}


}*/
/*err := repository.AdCart(body.Product_Id)

if err != nil {
	r.JSON(400, gin.H{
		"message": err.Error(),
	})
	return
}
r.JSON(200, gin.H{
	"message": "product added to cart successfully",
})
return
} //111111111111111111111*/
func Value(user_id int) string {
	var value string
	database.DB.Raw("SELECT id FROM orders WHERE user_id=$1;", user_id).Scan(&value)
	return value

}
func Address(user_id int) string {
	var Address string
	database.DB.Raw("SELECT id FROM addresses WHERE user_id=$1;", user_id).Scan(&Address)
	return Address

}
