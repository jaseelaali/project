package handlers

import (
	"github/jaseelaali/orchid/models"
	"github/jaseelaali/orchid/repository"

	"github.com/gin-gonic/gin"
)

func AddSubCategory(r *gin.Context) {
	Subcategory := models.SubCategory{}
	
	err := r.Bind(&Subcategory)
	if err != nil {
		r.JSON(400, gin.H{
			"message": "error in binding data",
		})
		return
	}
	err = repository.AddSubcategory(Subcategory)
	if err != nil {
		r.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	r.JSON(200, gin.H{
		"message": "sub-category added successfully",
	})
}
