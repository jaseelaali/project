package handlers

import (
	"encoding/csv"
	"fmt"
	"github/jaseelaali/orchid/database"
	"github/jaseelaali/orchid/models"
	"github/jaseelaali/orchid/repository"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// AdminLogin
// @Summary Admin Login
// @ID Admin-login
// @Description Admin can sign up with email and password
// @Tags Admin
// @Tags home
// @Accept json
// @Produce json
// @Param login body models.Admin true "Login credentials"
// @Success 200
// @Failure 400
// @Router /admin/superadminlogin [post]
func AdminLogin(r *gin.Context) {
	var login struct {
		Email    string
		Password string
	}
	if err := r.Bind(&login); err != nil {
		r.JSON(400, gin.H{"message": "error in binding data"})
	}
	//var password string
	admin := models.Admin{}
	database.DB.Where("email=?", login.Email).First(&admin)
	//err := database.DB.Raw("SELECT password from admins where email='$1'", login.Email).Scan(&password).Error
	err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(login.Password))
	if err != nil {
		r.JSON(400, gin.H{
			"message": "passwords or email are not matching",
		})
		return
	}
	/*-----------------------------------------------------------------------------------------------------------------------------*/
	/*password, _ := bcrypt.GenerateFromPassword(([]byte(login.Password)), 11)
	Password := string(password)
	fmt.Printf("****password want to databse:%v", Password)
	database.DB.Raw("INSERT INTO admins (password,email)VALUES($1,$2);", Password, login.Email).Scan(&models.Admin{})
	*/
	/*---------------------------------------------------------------------------------------------------------------------------------*/
	//generate jwt token
	//here call token function
	token, err := repository.Token(login.Email)

	//sign and get the complete encoded token as a string using the secret key
	tokenstring, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		r.JSON(400, gin.H{"message": "unable to create token"})
		return
	}
	// //send it back
	r.SetSameSite(http.SameSiteLaxMode)
	r.SetCookie("Authorization", tokenstring, 3600*24*30, "", "", false, true)
	r.JSON(200, gin.H{
		"token":   tokenstring,
		"message": "login successfully",
	})

	// func Validate(r *gin.Context){
	// 	r.JSON(200,gin.H{
	// 		"message":"loged in"
	// 	})
	// }

}

// sales report
// @Summary sales report
// @ID sales report
// @Description Admin can download sales report from here
// @Tags Admin
// @Tags sales report
// @Accept json
// @Produce json
// @Success 200
// @Failure 400
// @Router /admin/salesreport [get]
func SalesReport(r *gin.Context) {

	sales, err := repository.SalesReport()
	if err != nil {
		r.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	r.Header("Content-Type", "text/csv")
	r.Header("Content-Disposition", "attachment;filename=sales.csv")

	// Create a CSV writer using our response writer as our io.Writer
	wr := csv.NewWriter(r.Writer)

	// Write CSV header row
	headers := []string{"Order ID", "User ID", "Product ID", "Payment ID", "Delivery Status"}
	if err := wr.Write(headers); err != nil {
		r.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Write data rows
	for _, sale := range sales {
		row := []string{
			strconv.Itoa(int(sale.ID)),
			strconv.Itoa(sale.User_id),
			fmt.Sprintf("%d", sale.Product_id),
			string(sale.Payment_Id),
			sale.Delivery,
		}
		if err := wr.Write(row); err != nil {
			r.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}

	// Flush the writer's buffer to ensure all data is written to the client
	wr.Flush()
	if err := wr.Error(); err != nil {
		r.AbortWithError(http.StatusInternalServerError, err)
		return
	}
}
