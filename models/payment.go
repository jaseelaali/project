package models

import "time"

type Payment struct {
	Created_at      time.Time
	User_Id         int
	Order_Id        int
	Applied_Coupons string
	Discount        int
	Total_Amount    int
	Balance_Amount  int
	PaymentMethod   string
	Payment_Status  string
	Order_Status    string
	Address_Id      int
}
