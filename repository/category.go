package repository

import (
	"errors"
	"github/jaseelaali/orchid/database"
	"github/jaseelaali/orchid/models"
)

func Addcategory(category models.Category) error {
	Category := models.Category{}
	var count int
	err := database.DB.Raw("SELECT COUNT(category_name) FROM categories WHERE category_name=$1;", category.Category_Name).Scan(&count)
	if count == 0 {
		err = database.DB.Raw("INSERT INTO categories(category_name)VALUES ($1) ;", category.Category_Name).Scan(&Category)
		//fmt.Println("*1****** %v", category.Category_Name)
		if err != nil {
			return err.Error
		}
		return nil
	} else {
		return errors.New("this category already occured")
	}
}
func Editcategory(name string, Id int) error {
	var count int
	err := database.DB.Raw("SELECT COUNT(category_name) FROM categories WHERE category_name=$1;", name).Scan(&count)
	if count == 0 {
		err = database.DB.Raw("UPDATE categories SET category_name=$1 WHERE id=$2;", name, Id).Scan(&models.Category{})
		if err != nil {
			return err.Error
		}
		return nil
	} else {
		return errors.New("this category already occured")
	}
}
func Deletecategory(Id int) error {
	var exist int
	err := database.DB.Raw("SELECT COUNT(id) FROM categories WHERE id=$1;", Id).Scan(&exist)
	if exist == 0 {
		return errors.New("this category not occured")
	}

	err = database.DB.Raw("DELETE FROM categories WHERE id=$1", Id).Scan(&models.Category{})
	if err != nil {
		return err.Error
	}
	return nil

}
func Viewcategory() ([]models.Category, error) {
	category := []models.Category{}

	err := database.DB.Raw("SELECT * FROM categories;").Scan(&category)
	if err != nil {
		return category, err.Error
	}
	return category, nil
}
func EditSubcategory(name string, Id int) error {
	var count int
	err := database.DB.Raw("SELECT COUNT(sub_category_name) FROM sub_categories WHERE sub_category_name=$1;", name).Scan(&count)
	if count == 0 {

		err = database.DB.Raw("UPDATE sub_categories SET sub_category_name=$1 WHERE id=$2;", name, Id).Scan(&models.SubCategory{})
		if err != nil {
			return err.Error
		}
		return nil
	} else {
		return errors.New("this category already occured")
	}
}
func DeleteSubcategory(Id int) error {
	var exist int
	err := database.DB.Raw("SELECT COUNT(id) FROM sub_categories WHERE id=$1;", Id).Scan(&exist)
	if exist == 0 {
		return errors.New("this sub category not occured")
	}
	err = database.DB.Raw("DELETE FROM sub_categories WHERE id=$1", Id).Scan(&models.SubCategory{})
	if err != nil {
		return err.Error
	}
	return nil
}
func ViewSubcategory() ([]models.SubCategory, error) {
	subcategory := []models.SubCategory{}
	err := database.DB.Raw("SELECT * FROM sub_categories;").Scan(&subcategory)
	if err != nil {
		return subcategory, err.Error
	}
	return subcategory, nil
}
