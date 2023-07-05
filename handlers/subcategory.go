package handlers

import (
	"github/jaseelaali/orchid/models"
	"github/jaseelaali/orchid/repository"

	"github.com/gin-gonic/gin"
)
// ADDSUBCATEGORY
// @Summary add subcategory
// @ID add-subcategory
// @Description admin can add subcategories here
// @Tags Admin
//@Tags subcategory management
// @Accept json
// @Produce json
//param admin body models.SubCategory{} true"add product"
// @Success 200
// @Failure 400
// @Router /admin/addsubcategory [post]
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
