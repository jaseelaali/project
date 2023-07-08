package routes

import (
	"github/jaseelaali/orchid/handlers"
	"github/jaseelaali/orchid/middleware"

	//"github/jaseelaali/orchid/middleware"

	"github.com/gin-gonic/gin"
)

func User(r *gin.Engine) {
	r.LoadHTMLGlob("template/*.html")
	user := r.Group("/user")

	{
		// user signup+login
		user.POST("/signup", handlers.UserSignUp)
		user.POST("/loginuser", handlers.UserLogin)
		//profile
		user.GET("/aboutme", middleware.RequiredAuthenticationUser, handlers.Profile)
		user.PATCH("/editmyprofile", middleware.RequiredAuthenticationUser, handlers.EditProfile) //////
		user.DELETE("/deleteprofile", middleware.RequiredAuthenticationUser, handlers.DeleteProfile)/////
		// cart management
		user.POST("/addcart", middleware.RequiredAuthenticationUser, handlers.AddCart) ////
		user.GET("/viewcart", middleware.RequiredAuthenticationUser, handlers.ViewCart)
		user.DELETE("/deleteitem", middleware.RequiredAuthenticationUser, handlers.DeleteItem)
		//product
		user.GET("/viewproduct", middleware.RequiredAuthenticationUser, handlers.ViewProduct)
		//change password
		user.POST("/changepassword", middleware.RequiredAuthenticationUser, handlers.ChangePassword)
		user.POST("/verifyotp", middleware.RequiredAuthenticationUser, handlers.VerifyOtp)////
		// address
		user.POST("/addaddress", middleware.RequiredAuthenticationUser, handlers.Address)///
		user.PATCH("/editaddress", middleware.RequiredAuthenticationUser, handlers.EditAddress)//
		user.DELETE("/deleteaddress", middleware.RequiredAuthenticationUser, handlers.DeleteAddress)///
		user.GET("/viewaddress", middleware.RequiredAuthenticationUser, handlers.ViewAddress)///
		// order management
		user.POST("/addorder", middleware.RequiredAuthenticationUser, handlers.AddOrder)
		user.GET("/showorder", middleware.RequiredAuthenticationUser, handlers.ShowOrder)
		user.DELETE("/cancelorder", middleware.RequiredAuthenticationUser, handlers.CancelOrder)
		//payment
		user.GET("/razorpay", middleware.RequiredAuthenticationUser, handlers.Razorpay)
		user.GET("/payment-success", middleware.RequiredAuthenticationUser, handlers.Payment_Success)
		//coupens
		user.GET("/listcoupens", middleware.RequiredAuthenticationUser, handlers.Listcoupens)
		user.POST("/applycoupens", middleware.RequiredAuthenticationUser, handlers.ApplyCoupen)
		//wishlist
		user.POST("/addwishlist", middleware.RequiredAuthenticationUser, handlers.AddWishList)
		user.GET("/listwishlist", middleware.RequiredAuthenticationUser, handlers.ListWishlist)
		user.DELETE("/removewishlist", middleware.RequiredAuthenticationUser, handlers.RemoveWishlist)
		//Return my product
		user.POST("/returnproduct", middleware.RequiredAuthenticationUser, handlers.ReturnMyProduct)
		//wallet
		user.GET("/mywallet", middleware.RequiredAuthenticationUser, handlers.MyWallet)

	}
}
func Admin(r *gin.Engine) {
	admin := r.Group("/admin")
	{
		// super admin login
		admin.POST("/superadminlogin", handlers.AdminLogin)
		//admin management
		// admin.POST("/addadmins", middleware.RequiredAuthenticationAdmin, handlers.AddAdmins)
		// admin.GET("/viewadmins", middleware.RequiredAuthenticationAdmin, handlers.ViewAdmins)
		// admin.DELETE("/deleteadmins", middleware.RequiredAuthenticationAdmin, handlers.DeleteAdmins)
		// admin login
		//admin.POST("/adminlogin", handlers.AdminLogin)

		//user management
		admin.GET("/view", middleware.RequiredAuthenticationAdmin, handlers.ViewUser)
		admin.GET("/speacificuser", middleware.RequiredAuthenticationAdmin, handlers.SpeacificUser)
		admin.POST("/block", middleware.RequiredAuthenticationAdmin, handlers.BlockUser)
		admin.POST("/unblock", middleware.RequiredAuthenticationAdmin, handlers.UnBlockUser)
		admin.GET("/viewblockedusers", middleware.RequiredAuthenticationAdmin, handlers.BlockedUsers)
		admin.GET("/viewunblockedusers", middleware.RequiredAuthenticationAdmin, handlers.ActiveUsers)
		// category mangement
		admin.POST("/addcategory", middleware.RequiredAuthenticationAdmin, handlers.AddCategory)
		admin.PATCH("/editcategory", middleware.RequiredAuthenticationAdmin, handlers.EditCategory)
		admin.DELETE("/deletecategory", middleware.RequiredAuthenticationAdmin, handlers.DeleteCategory)
		admin.GET("/viewcategory", middleware.RequiredAuthenticationAdmin, handlers.ViewCategory)
		// sub category management
		admin.POST("/addsubcategory", middleware.RequiredAuthenticationAdmin, handlers.AddSubCategory)
		admin.PATCH("/editsubcategory", middleware.RequiredAuthenticationAdmin, handlers.EditSubCategory)
		admin.DELETE("/deletesubcategory", middleware.RequiredAuthenticationAdmin, handlers.DeleteSubCategory)
		admin.GET("/viewsubcategory", middleware.RequiredAuthenticationAdmin, handlers.ViewSubCategory)
		//product managemant
		admin.POST("/addproduct", middleware.RequiredAuthenticationAdmin, handlers.AddProducts)
		admin.PATCH("/editproducts", middleware.RequiredAuthenticationAdmin, handlers.EditProducts)
		admin.DELETE("/deleteproducts", middleware.RequiredAuthenticationAdmin, handlers.DeleteProducts)
		admin.GET("/viewproducts", middleware.RequiredAuthenticationAdmin, handlers.ViewProducts)
		//coupen
		admin.POST("/addcoupen", middleware.RequiredAuthenticationAdmin, handlers.AddCoupens)
		admin.GET("/listcoupen", middleware.RequiredAuthenticationAdmin, handlers.ListCoupen)
		//delivery complete status change
		admin.POST("/returnstatus", middleware.RequiredAuthenticationAdmin, handlers.ReturnStatus)

		//sale report
		admin.GET("/salesreport", middleware.RequiredAuthenticationAdmin, handlers.SalesReport)

	}
}
