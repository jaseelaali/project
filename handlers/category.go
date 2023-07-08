package handlers

import (
	"github/jaseelaali/orchid/models"
	"github/jaseelaali/orchid/repository"
	"strconv"

	"github.com/gin-gonic/gin"
)


func AddCategory(r *gin.Context) {
	category := models.Category{}
	err := r.Bind(&category)
	if err != nil {
		r.JSON(400, gin.H{
			"message": "error in binding data",
		})
		return
	}
	err = repository.Addcategory(category)
	if err != nil {
		r.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	r.JSON(200, gin.H{
		"message": "category added successfully",
	})
}


func EditCategory(r *gin.Context) {
	id, err := strconv.Atoi(r.Query("id"))
	if err != nil {
		r.JSON(400, gin.H{
			"message": "didn't get id",
		})
		return
	}
	name := r.Query("name")
	if name == "" {
		r.JSON(400, gin.H{
			"message": "didn't get name",
		})
		return
	}
	err = repository.Editcategory(name, id)
	if err != nil {
		r.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	r.JSON(200, gin.H{
		"message": "category updated",
	})

}


func DeleteCategory(r *gin.Context) {
	id, err := strconv.Atoi(r.Query("id"))
	if err != nil {
		r.JSON(400, gin.H{
			"message": "didn't get name",
		})
		return
	}
	err = repository.Deletecategory(id)
	if err != nil {
		r.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	r.JSON(200, gin.H{
		"message": "category deleted",
	})
}


func ViewCategory(r *gin.Context) {

	Category, err := repository.Viewcategory()
	if err != nil {
		r.JSON(400, gin.H{
			"message": "failed to list category",
		})
		return
	}
	r.JSON(200, gin.H{
		"message": Category,
	})

}


func EditSubCategory(r *gin.Context) {
	id, err := strconv.Atoi(r.Query("id"))
	if err != nil {
		r.JSON(400, gin.H{
			"message": "didn't get id",
		})
		return
	}
	name := r.Query("name")
	if name == "" {
		r.JSON(400, gin.H{
			"message": "didn't get name",
		})
		return
	}
	err = repository.EditSubcategory(name, id)
	if err != nil {
		r.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	r.JSON(200, gin.H{
		"message": "subcategory updated",
	})

}


func DeleteSubCategory(r *gin.Context) {
	id, err := strconv.Atoi(r.Query("id"))
	if err != nil {
		r.JSON(400, gin.H{
			"message": "didn't get name",
		})
		return
	}
	err = repository.DeleteSubcategory(id)
	if err != nil {
		r.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	r.JSON(200, gin.H{
		"message": "subcategory deleted",
	})
}


func ViewSubCategory(r *gin.Context) {
	SubCategory, err := repository.ViewSubcategory()
	if err != nil {
		r.JSON(400, gin.H{
			"message": "failed to list subcategory",
		})
		return
	}
	r.JSON(400, gin.H{
		"message": SubCategory,
	})

}
