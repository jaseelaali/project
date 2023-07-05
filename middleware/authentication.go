package middleware

import (
	//"bytes"

	"fmt"
	"github/jaseelaali/orchid/database"
	"github/jaseelaali/orchid/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func RequiredAuthenticationUser(r *gin.Context) {
	tokenString, err := r.Cookie("Authorization")
	if err != nil {
		r.JSON(http.StatusUnauthorized, gin.H{"error": "Plese login first"})
		r.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	token, _ := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])

		}
		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			r.JSON(http.StatusUnauthorized, gin.H{"error": "Autherization failed"})
			r.Abort()
		}
		user := models.User{}
		result := database.DB.Where("email", claims["sub"]).First(&user)
		if result.Error != nil {
			r.JSON(http.StatusUnauthorized, gin.H{"error": "not user found!"})
		}
		r.Set("user_id", user.ID)
		r.Next()
	} else {
		r.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Autherization failed !"})
	}
}
func RequiredAuthenticationAdmin(r *gin.Context) {

	// Get the cookie off req
	tokenString, err := r.Cookie("Authorization")
	if err != nil {
		r.JSON(http.StatusUnauthorized, gin.H{"error": "Plese login first"})
		r.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	// Pzarse takes the token string and a function for looking up the key. The latter is especially

	token, _ := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		// Check the expm
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			r.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		// Find the user with token sub
		var admin models.Admin
		database.DB.Model(&models.Admin{}).Where("email", claims["sub"]).First(&admin)
		if admin.ID == 0 {
			fmt.Println("check")
			r.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Attach to request
		r.Set("admin_id", admin.ID)
		r.Next()

	} else {
		r.JSON(http.StatusUnauthorized, gin.H{"error": "Autherization failed"})
		r.AbortWithStatus(http.StatusUnauthorized)
		return
	}

}
