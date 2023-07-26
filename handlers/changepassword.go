package handlers

import (
	// "fmt"
	// "math/rand"
	// "strconv"

	"fmt"
	"github/jaseelaali/orchid/database"
	"github/jaseelaali/orchid/models"
	"github/jaseelaali/orchid/repository"
	"math/rand"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sfreiberg/gotwilio"
	"golang.org/x/crypto/bcrypt"
)

var OTP int
var twilio = gotwilio.NewTwilioClient("AC3037e122f46a35ae97b5a48f7413be56", "0bc43f0b4e4a492d46e26bf093c0fc40")

// change password
// @Summary Change password
// @ID change-password
// @Description User can change password
// @Tags User
// @Tags Change password
// @Accept json
// @Produce json
// @Success 200
// @Failure 400
// @Router /user/changepassword [post]
func ChangePassword(r *gin.Context) {
	user_id, _ := r.Get("user_id")
	User_Id, _ := strconv.Atoi(fmt.Sprint(user_id))
	var mobilenumber string
	err := database.DB.Raw("SELECT phone_number FROM users WHERE id=$1;", User_Id).Scan(&mobilenumber)
	if err.Error != nil {
		r.JSON(400, gin.H{
			"message": err.Error,
		})
		return
	}
	otp, Error := sendOTP(mobilenumber)
	OTP = otp
	if Error != nil {
		r.JSON(400, gin.H{
			"message": "failed to send message",
			"error":   Error,
		})
		return
	}
	r.JSON(200, gin.H{
		"message": "successfully send the otp",
		"data":    OTP,
	})
}
func sendOTP(phoneNumber string) (int, error) {
	otpCode := generateOTP()
	fmt.Println(otpCode)
	message := "Your OTP code is " + strconv.Itoa(otpCode)
	_, _, err := twilio.SendSMS("+15302036484", "+91"+phoneNumber, message, "", "")
	if err != nil {
		return -1, err
	}
	return otpCode, nil
}
func generateOTP() int {
	// Generate a random 4-digit OTP code
	otp := rand.Intn(8899) + 1000
	fmt.Println(otp)
	return otp
}

type VerifyOtpResponse struct {
	Otp              int
	New_Password     string
	Confirm_Password string
}
// VERIFYOTP
// @Summary VERIFY OTP
// @ID verify-otp
// @Description User can verify otp for password change
// @Tags User
// @Tags Change password
// @Accept json
// @Produce json
// @Param VerifyOtpResponse body VerifyOtpResponse true "otp"
// @Success 200
// @Failure 400
// @Router /user/verifyotp [post]
func VerifyOtp(r *gin.Context) {
	var body struct {
		Otp              int
		New_Password     string
		Confirm_Password string
	}
	err := r.Bind(&body)
	if err != nil {
		r.JSON(400, gin.H{
			"message": "error in binding data",
		})
		return
	}

	isValid := VerifyOTP(body.Otp, OTP)
	if isValid == true {
		//userid, _ := strconv.Atoi(fmt.Sprint(r.Get("user_id")))
		userid := repository.GetId(r)
		newpasword := body.New_Password
		confirmpasword := body.Confirm_Password
		if newpasword == confirmpasword {
			password, err := (bcrypt.GenerateFromPassword([]byte(newpasword), 11))
			Pasword := string(password)
			var poi string
			database.DB.Raw("SELECT password FROM users WHERE id=$1", userid).Scan(&poi)
			//database.DB.Model(&User).Select("password").Updates(map[string]interface{}{"password": "hello", "age": 18, "active": false})
			// UPDATE users SET name='hello' WHERE id=111;
			result := database.DB.Raw("UPDATE users SET password=$1 WHERE id=$2", Pasword, userid).Scan(&models.User{})
			fmt.Println(result.Error)
			if result.Error != nil {
				r.JSON(400, gin.H{
					"message": err.Error(),
				})
				return
			}
			r.JSON(200, gin.H{
				"message": "password changed successfully",
			})
			return
		}
		r.JSON(400, gin.H{
			"message": "passwords are not matched",
		})
		return
	} else {
		r.JSON(400, gin.H{
			"message": "invalid otp",
		})
		return
	}
}
func VerifyOTP(otpCode, expectedCode int) bool {
	if otpCode == expectedCode {
		return true
	}
	return false
}
