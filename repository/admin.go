package repository

import (
	"github/jaseelaali/orchid/database"
	"github/jaseelaali/orchid/models"
)

func SalesReport() ([]models.OrderStatus, error) {
	var salesData []models.OrderStatus

	salesReportQuery := `SELECT os.id, os.user_id,os.product_id, os.payment_id, os.delivery FROM order_statuses as os;`

	if err := database.DB.Raw(salesReportQuery).Scan(&salesData).Error; err != nil {
		return []models.OrderStatus{}, err
	}
	return salesData, nil
}

// func AddAdmin(email, name, password string) error {
// 	var model []models.Admin
// 	result := database.DB.Raw("SELECT * FROM admins WHERE email=$1;", email).Scan(model)
// 	if model != nil {
// 		return errors.New("email already occured")
// 	}
// 	Pasword, err := (bcrypt.GenerateFromPassword([]byte(password), 11))
// 	if err != nil {
// 		return err
// 	}
// 	password = string(Pasword)
// 	result = database.DB.Raw("INSERT INTO admins(admin_name,email,password)VALUES($1,$2,$3);", name, email, password).Scan(&models.Admin{})
// 	if result.Error != nil {
// 		return result.Error
// 	}
// 	return nil
// }
// func ViewAllAdmins(page, perpage int) ([]models.Admin, utils.MetaData, error) {
// 	admins := []models.Admin{}
// 	var totalRecords int64
// 	err := database.DB.Raw("SELECT COUNT(id) FROM admins  ;").Scan(&totalRecords).Error
// 	if err != nil {
// 		return admins, utils.MetaData{}, err
// 	}
// 	metaData, offset, err := utils.ComputeMetaData(page, perpage, int(totalRecords))
// 	if err != nil {
// 		return admins, metaData, err
// 	}
// 	err = database.DB.Raw("SELECT* FROM admins OFFSET ? LIMIT ?;", offset, perpage).Scan(&admins).Error
// 	if err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			return admins, metaData, errors.New("Record not found")
// 		}
// 		return admins, metaData, err
// 	}
// 	return admins, metaData, nil
// }
// func DeleteAnAdmin(id int) error {
// 	result := database.DB.Raw("DELETE FROM admins WHERE id=$1;", id).Scan(&models.Admin{})
// 	if result.Error != nil {
// 		return result.Error
// 	}
// 	return nil
// }
// func AdminsLogin(email, password string) error {
// 	var passwordfromdatabase string
// 	result := database.DB.Raw("SELECT password FROM admins WHERE email=$1;", email).Scan(&passwordfromdatabase)
// 	if result.Error != nil {
// 		return result.Error
// 	}
// 	err := bcrypt.CompareHashAndPassword([]byte(passwordfromdatabase), []byte(password))
// 	if err != nil {
// 		return result.Error
// 	}
// 	return nil
// }
