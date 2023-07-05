package handlers

import (
	"github/jaseelaali/orchid/models"
	"github/jaseelaali/orchid/repository"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ADDCATEGORY
// @Summary add category
// @ID add-category
// @Description admin can add categories here
// @Tags Admin
// @Tags category management
// @Accept json
// @Produce json
// param admin body models.Category{} true"add product"
// @Success 200
// @Failure 400
// @Router /admin/addcategory [post]
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

// EditCategory
// @Summary Edit Category
// @ID Edit-category
// @Description admin can edit category
// @Tags Admin
// @Tags category management
// @Accept json
// @Produce json
// @Success 200
// @Failure 400
// @Router /admin/editcategory [patch]
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

// Delete Category
// @Summary Delete Category
// @ID delete-category
// @Description admin can delete category
// @Tags Admin
// @Tags category management
// @Accept json
// @Produce json
// @Success 200
// @Failure 400
// @Router /admin/deletecategory [delete]
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

// View Category
// @Summary viewcategory
// @ID view-category
// @Description admin can view category
// @Tags Admin
// @Tags category management
// @Accept json
// @Produce json
// @Success 200
// @Failure 400
// @Router /admin/viewcategory [get]
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

// EditSubCategory
// @Summary Edit SUb Category
// @ID Edit-subcategory
// @Description admin can edit sub category
// @Tags Admin
// @Tags subcategory management
// @Accept json
// @Produce json
// @Success 200
// @Failure 400
// @Router /admin/editsubcategory [patch]
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

// Delete SubCategory
// @Summary Delete SubCategory
// @ID delete-subcategory
// @Description admin can delete subcategory
// @Tags Admin
// @Tags subcategory management
// @Accept json
// @Produce json
// @Success 200
// @Failure 400
// @Router /admin/deletesubcategory [delete]
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

// View SubCategory
// @Summary viewsubcategory
// @ID view-subcategory
// @Description admin can view subcategory
// @Tags Admin
// @Tags subcategory management
// @Accept json
// @Produce json
// @Success 200
// @Failure 400
// @Router /admin/viewsubcategory [get]
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
