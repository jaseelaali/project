package repository

import (
	"errors"
	"fmt"
	"github/jaseelaali/orchid/database"
	"github/jaseelaali/orchid/models"
	"github/jaseelaali/orchid/utils"

	"gorm.io/gorm"
)

func Addproduct(product models.Products) error {
	var count int
	err := database.DB.Raw("SELECT COUNT(product_name) FROM products WHERE product_name=$1;", product.Product_Name).Scan(&count)
	if count == 0 {
		err = database.DB.Create(&models.Products{
			Product_Name:   product.Product_Name,
			Product_Colour: product.Product_Colour,
			Product_Size:   product.Product_Size,
			Product_Brand:  product.Product_Brand,
			Product_Price:  product.Product_Price,
			Stock:          product.Stock,
		})
		// err := database.DB.Raw("INSERT INTO products (product_name,product_colour,product_size,product_brand,product_price) VALUES ($1,$2,$3,$4,$5);", product.Product_Name,
		// 	product.Product_Colour, product.Product_Size, product.Product_Brand, product.Product_Price).Scan(&newproduct)
		if err.Error != nil {
			return err.Error
		}
		return nil
	} else {
		return errors.New("this product already occured")
	}
}
func EditProductName(name string, Id int) error {
	var count int
	err := database.DB.Raw("SELECT COUNT(product_name) FROM products WHERE product_name=$1;", name).Scan(&count)
	if count == 0 {
		fmt.Println(".........................", name)
		err = database.DB.Raw("UPDATE products SET product_name=$1 WHERE id =$2;", name, Id).Scan(&models.Products{})
		return err.Error
	} else {
		return errors.New("this product already occured")
	}
}
func EditProductColour(colour string, Id int) error {
	fmt.Println(".........................", colour, Id)

	err := database.DB.Raw("UPDATE products SET product_colour=$1 WHERE id =$2;", colour, Id).Scan(&models.Products{})
	return err.Error
}
func EditProductSize(size, Id int) error {
	err := database.DB.Raw("UPDATE products SET product_size=$1 WHERE id =$2;", size, Id).Scan(&models.Products{})
	return err.Error
}
func EditProductBrand(Brand string, Id int) error {
	err := database.DB.Raw("UPDATE products SET product_brand=$1 WHERE id =$2;", Brand, Id).Scan(&models.Products{})
	return err.Error
}
func EditProductPrice(price, Id int) error {
	err := database.DB.Raw("UPDATE products SET product_price=$1 WHERE id =$2;", price, Id).Scan(&models.Products{})
	return err.Error
}
func EditProductStock(stock, Id int) error {
	err := database.DB.Raw("UPDATE products SET stock=$1 WHERE id =$2;", stock, Id).Scan(&models.Products{})
	return err.Error
}
func Deleteproduct(Id int) error {
	var exist int
	err := database.DB.Raw("SELECT COUNT(id) FROM products WHERE id=$1;", Id).Scan(&exist)
	if exist == 0 {
		return errors.New("this product not occured")
	}

	err = database.DB.Raw("DELETE FROM products WHERE id=$1", Id).Scan(&models.Products{})
	if err != nil {
		return err.Error
	}
	return nil

}
func Viewproduct(page, perPage int) ([]models.Products, utils.MetaData, error) {
	Products := []models.Products{}
	var totalRecords int64
	err := database.DB.Raw("SELECT COUNT(id) FROM products  ;").Scan(&totalRecords).Error
	if err != nil {
		return Products, utils.MetaData{}, err
	}
	metaData, offset, err := utils.ComputeMetaData(page, perPage, int(totalRecords))
	if err != nil {
		return Products, metaData, err
	}
	err = database.DB.Raw("SELECT* FROM products OFFSET ? LIMIT ?;", offset, perPage).Scan(&Products).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return Products, metaData, errors.New("Record not found")
		}
		return Products, metaData, err
	}
	return Products, metaData, nil

}
