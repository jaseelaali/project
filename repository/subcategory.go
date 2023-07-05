package repository

import (
	"errors"
	"github/jaseelaali/orchid/database"
	"github/jaseelaali/orchid/models"
)

func AddSubcategory(Subcategory models.SubCategory) error {
	SubCategory := models.SubCategory{}
	var count int
	err := database.DB.Raw("SELECT COUNT(sub_category_name) FROM sub_categories WHERE sub_category_name=$1;", Subcategory.SubCategory_Name).Scan(&count)
	if count == 0 {
		err = database.DB.Raw("INSERT INTO sub_categories(sub_category_name)VALUES ($1);", Subcategory.SubCategory_Name).Scan(&SubCategory)
		// fmt.Println("******* %v", Subcategory.SubCategory_Name)
		if err != nil {
			return err.Error
		}
		return nil
	} else {
		return errors.New("this category already occured")
	}
}
