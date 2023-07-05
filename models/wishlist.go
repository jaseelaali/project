package models

type WishList struct {
	User_id    int `json:"user_id"`
	Product_id int `json:"product_id"`
}
type ViewWishList struct {
	User_id      int `json:"user_id"`
	Product_id   int `json:"product_id"`
	Product_Name string
}
type Wallet struct {
	User_id int
	Money   int
}

