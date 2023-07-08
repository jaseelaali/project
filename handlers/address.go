package handlers

import (
	"fmt"
	"github/jaseelaali/orchid/models"
	"github/jaseelaali/orchid/repository"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	_ "gorm.io/gorm"
)

type AddAddressResponse struct {
	House_name string
	Place      string
	District   string
	State      string
	Pin_Number int
}

func Address(r *gin.Context) {
	var Body struct {
		House_name string
		Place      string
		District   string
		State      string
		Pin_Number int
	}
	temp := fmt.Sprint(r.Get("user_id"))
	id := strings.Split(temp, " ")
	Id, _ := strconv.Atoi(id[0])
	fmt.Println("id string is :", id, "id is :", Id)
	err := r.Bind(&Body)
	if err != nil {
		r.JSON(400, gin.H{
			"message": "error in binding data",
		})
		return
	}
	address := models.Address{

		House_name: Body.House_name,
		Place:      Body.Place,
		District:   Body.District,
		State:      Body.State,
		Pin_Number: Body.Pin_Number,
	}
	err = repository.AddAddress(address, Id)
	if err != nil {
		r.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	r.JSON(200, gin.H{
		"message": "successfully added address",
	})
}

type EditAddressResponse struct {
	House_name string
	Place      string
	District   string
	State      string
	Pin_Number int
}

func EditAddress(r *gin.Context) {
	Id := repository.GetId(r)
	var body struct {
		House_name string `json:"house_name"`
		Place      string `json:"place"`
		District   string `json:"district"`
		State      string `json:"state"`
		Pin_Number int    `json:"pin_number"`
	}
	err := r.Bind(&body)
	if err != nil {
		r.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	if body.House_name != "" {
		err := repository.EditHousename(body.House_name, Id)
		if err != nil {
			r.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}

	}
	if body.Place != "" {
		err = repository.EditPlace(body.Place, Id)
		if err != nil {
			r.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
	}
	if body.District != "" {
		err = repository.EditDistrict(body.District, Id)
		if err != nil {
			r.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
	}
	if body.State != "" {
		err = repository.EditState(body.State, Id)
		if err != nil {
			r.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
	}
	if body.Pin_Number != 0 {
		err = repository.EditPin(body.Pin_Number, Id)
		if err != nil {
			r.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
	}
	r.JSON(200, gin.H{
		"message": "addressupdated",
	})

}

func DeleteAddress(r *gin.Context) {
	User_Id := repository.GetId(r)
	var Body struct {
		Address_id int `json:"address_id" binding:"required"`
	}
	err := r.Bind(&Body)
	if err != nil {
		r.JSON(400, gin.H{
			"MESSAGE": "error in binding data",
		})
		return
	}
	err = repository.Deleteaddress(User_Id, Body.Address_id)
	if err != nil {
		r.JSON(400, gin.H{
			"MESSAGE": err.Error(),
		})
		return
	}
	r.JSON(200, gin.H{
		"MESSAGE": "address deleted",
	})

}


func ViewAddress(r *gin.Context) {
	user_id := repository.GetId(r)
	address, err := repository.Viewaddress(user_id)
	if err != nil {
		r.JSON(400, gin.H{
			"MESSAGE": err.Error(),
		})
		return
	}

	r.JSON(200, gin.H{
		"MESSAGE": address,
	})
}
