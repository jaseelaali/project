package handlers

import (
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

type UserSignupResponse struct {
	User_Name    string
	Email        string
	Phone_Number string
	Password     string
}

// UserSignup
// @Summary User Signup
// @ID user-signup
// @Description User can sign up with email and password
// @Tags User
// @Tags Home
// @Accept json
// @Produce json
// @Param newUser body UserSignupResponse true "user credentials for creating new account"
// @Success 200
// @Failure 400
// @Router /user/signup [post]
func UserSignUp(r *gin.Context) {
	newUser := models.User{}
	if err := r.Bind(&newUser); err != nil {
		r.JSON(400, gin.H{"message": "error in binding data", "error": err})
		return
	}
	Password, err := (bcrypt.GenerateFromPassword([]byte(newUser.Password), 11))
	if err != nil {
		r.JSON(400, gin.H{"message": "error in hashing password"})
		return
	}
	newUser.Password = string(Password)
	newUser.Status = "active"
	err = repository.CreateUser(newUser)
	if err != nil {
		r.JSON(400, err.Error())
		return
	}
	r.JSON(200, gin.H{"success": "Created new user successfully "})
}

type UserLoginResponse struct {
	Email    string
	Password string
}

// UserLogin
// @Summary User LOGIN
// @ID user-LOGIN
// @Description User can login here
// @Tags User
// @Tags Home
// @Accept json
// @Produce json
// @Param newUser body UserLoginResponse{} true "user credentials for creating new account"
// @Success 200
// @Failure 400
// @Router /user/loginuser [post]
func UserLogin(r *gin.Context) {
	var login struct {
		Email    string
		Password string
	}
	if err := r.Bind(&login); err != nil {
		r.JSON(400, gin.H{"message": "error in binding data"})
		return
	}
	//var password string
	user := &models.User{}
	database.DB.Where(&models.User{Email: login.Email, Status: "active"}).First(&user)
	fmt.Println()
	if user == nil {
		r.JSON(400, gin.H{
			"message": "innalid user",
		})
		return
	}

	//database.DB.Raw("SELECT password FROM users WHERE email='$1'AND status='active';", login.Email).Scan(&password)
	password := user.Password

	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(login.Password))

	if err != nil {
		r.JSON(400, gin.H{"message": "passwords are not matching "})
		return
	}
	//r.JSON(200, gin.H{"message": "login successfully"})

	//generate jwt token
	//here call token function
	token, err := repository.Token(login.Email)

	//sign and get the complete encoded token as a string using the secret key
	tokenstring, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		r.JSON(400, gin.H{"message": "unable to create token"})
		return
	}
	//send it back
	r.SetSameSite(http.SameSiteLaxMode)
	r.SetCookie("Authorization", tokenstring, 3600*24*30, "", "", false, true)
	r.JSON(200, gin.H{
		"token":   tokenstring,
		"message": "login successfully",
	})

}

// VIEW USER
// @Summary VIEW USER
// @ID view-user
// @Description admin can view user here
// @Tags Admin
// @Tags user management
// @Accept json
// @Produce json
// @Param page query integer true "page"
// @Param perpage query integer  true "perpage"
// @Success 200
// @Failure 400
// @Router /admin/view [get]
func ViewUser(r *gin.Context) {

	page := r.Query("page")
	if page == "" {
		r.JSON(400, gin.H{
			"message": "didn't get page number",
		})
		return
	}
	perpage := r.Query("perpage")
	if perpage == "" {
		r.JSON(400, gin.H{
			"message": "didn't get perpage number",
		})
		return
	}
	pagenumber, _ := strconv.Atoi(page)
	perpagenumber, _ := strconv.Atoi(perpage)

	users, metaData, err := repository.View(pagenumber, perpagenumber)
	if err != nil {
		r.JSON(400, gin.H{
			"message": err.Error()})
	} else {
		r.JSON(200, gin.H{
			//"message":"the list
			"List of users": users,
			"metadata":      metaData})
	}
}

// VIEW SPECIFIC USER
// @Summary VIEW SPECIFIC USER
// @ID view-specific user
// @Description admin can view specific user here
// @Tags Admin
// @Tags user management
// @Accept json
// @Produce json
// @Param id query integer false "Id"
// @Param User_Name query string false "User_Name"
// @Param Email query string false "Email"
// @Param Phone_Number query string false "Phone_Number"
// @Success 200
// @Failure 400
// @Router /admin/speacificuser [get]
func SpeacificUser(r *gin.Context) {
	var body struct {
		Id           int    `json:"id"`
		User_Name    string `json:"user_name"`
		Email        string `json:"email"`
		Phone_Number string `json:"phone_number"`
	}
	if err := r.Bind(&body); err != nil {
		r.JSON(400, gin.H{"message": "error in binding data"})
		return
	}
	fmt.Printf(".................%v", body.Id)
	if body.Id != 0 {
		information, err := repository.UserById(body.Id)

		if err != nil {
			r.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		r.JSON(200, gin.H{
			"success": information,
		})
	}
	if body.User_Name != "" {
		information, err := repository.UserByName(body.User_Name)
		if err != nil {
			r.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		r.JSON(200, gin.H{
			"success": information,
		})
	}
	if body.Email != "" {
		information, err := repository.UserByEmail(body.Email)
		if err != nil {
			r.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		r.JSON(200, gin.H{
			"success": information,
		})
	}
	if body.Phone_Number != "" {
		information, err := repository.UserByNumber(body.Phone_Number)
		if err != nil {
			r.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		r.JSON(200, gin.H{
			"success": information,
		})
	}

}

// BlockUser
// @Summary BLOCK USER
// @ID block user
// @Description admin can block user here
// @Tags Admin
// @Tags user management
// @Accept json
// @Produce json
// @Param id query string  true "id"
// @Success 200
// @Failure 400
// @Router /admin/block [post]
func BlockUser(r *gin.Context) {
	ID, _ := strconv.Atoi(r.Query("id"))
	err := repository.BlockUser(ID)
	if err != nil {
		r.JSON(400, gin.H{"error": err.Error()})
		return
	}
	r.JSON(200, gin.H{"success": "Blocked user successfully"})
}

// UnBlockUser
// @Summary UnBLOCK USER
// @ID unblock user
// @Description admin can unblock user here
// @Tags Admin
// @Tags user management
// @Accept json
// @Produce json
// @Param id query string  true "id"
// @Success 200
// @Failure 400
// @Router /admin/unblock [post]
func UnBlockUser(r *gin.Context) {
	ID, _ := strconv.Atoi(r.Query("id"))
	err := repository.UnBlocUser(ID)
	if err != nil {
		r.JSON(400, gin.H{"error": err.Error()})
	} else {
		r.JSON(200, gin.H{"success": "Unblocked user successfully"})
	}
}

// BlockedUsers
// @Summary Get blocked users
// @ID blockedusers
// @Description Retrieve a list of blocked users with pagination
// @Security BearerAuth
// @Tags Admin
// @Tags user management
// @Accept json
// @Produce json
// @Param page query integer true "page"
// @Param perpage query integer true "perpage"
// @Success 200
// @Failure 400
// @Router /admin/viewblockedusers [get]
func BlockedUsers(r *gin.Context) {
	var Body struct {
		Page    int
		Perpage int
	}

	err := r.ShouldBind(&Body)
	if err != nil {
		r.JSON(400, gin.H{
			"message": err.Error()})
		return
	}

	users, metadata, err := repository.BlocUsers(Body.Page, Body.Perpage)
	if err != nil {
		r.JSON(400, gin.H{
			"message": err.Error()})
	} else {
		r.JSON(200, gin.H{
			//"message":"the list
			"message":  users,
			"metadata": metadata})
	}

}

// UNBlockedUsers
// @Summary Get un blocked users
// @ID unblockedusers
// @Description Retrieve a list of active users with pagination
// @Security BearerAuth
// @Tags Admin
// @Tags user management
// @Accept json
// @Produce json
// @Param page query integer true "page"
// @Param perpage query integer true "perpage"
// @Success 200
// @Failure 400
// @Router /admin/viewunblockedusers [get]
func ActiveUsers(r *gin.Context) {
	var Body struct {
		Page    int `json:"page"`
		Perpage int `json:"perpage"`
	}
	err := r.ShouldBind(&Body)
	if err != nil {
		r.JSON(400, gin.H{
			"message": err.Error()})
		return
	}

	users, metadata, err := repository.ActiveUser(Body.Page, Body.Perpage)
	if err != nil {
		r.JSON(400, gin.H{
			"message": err.Error()})
	} else {
		r.JSON(200, gin.H{
			"message":  users,
			"metadata": metadata})
	}
}

// func ChangePassword(r *gin.Context) {
// 	var Password struct {
// 		//sEmail        string
// 		Phone_Number string
// 		NewPassword  string
// 	}
// 	if err := r.Bind(&Password); err != nil {
// 		r.JSON(400, gin.H{"message": "error in binding data"})
// 	}
// 	var email string
// 	//database.DB.Where(&models.User{Email: Password.Email, Phone_Number: Password.Phone_Number}).First(&Id)

// 	database.DB.Raw("SELECT email FROM user WHERE  phone_number=$1;", Password.Phone_Number).Scan(&email)

// 	//Id, _ := strconv.Atoi(r.Query("ID"))
// 	fmt.Println("-------------------------")
// 	fmt.Println(email)

// 	fmt.Println("-------------------------")

// 	err := repository.NewPassword(Password.NewPassword, email)
// 	if err != nil {
// 		r.JSON(400, gin.H{
// 			"message": err.Error})
// 	} else {
// 		r.JSON(200, gin.H{
// 			"message": "password changed"})
// 	}
// }
