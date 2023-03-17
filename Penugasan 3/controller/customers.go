package controller

import (
	"Penugasan-3/config"
	"Penugasan-3/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CustomersControl struct {
	DB *gorm.DB
}

func (c *CustomersControl) PostCustomer(ctx *gin.Context) {
	var cust models.Customers

	err := ctx.ShouldBindJSON(&cust)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, config.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	tx := c.DB.Create(&cust)
	if tx.Error != nil {
		ctx.JSON(http.StatusBadRequest, config.Response{
			Status:  false,
			Message: tx.Error.Error(),
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, config.Response{
		Status:  true,
		Message: "user created successfully",
		Data:    cust,
	})
}

func (c *CustomersControl) UpdateCustomer(ctx *gin.Context) {
	var cust models.Customers

	id := ctx.Params.ByName("id")
	find := c.DB.Where("id = ?", id).First(&cust)
	if find.Error != nil {
		ctx.JSON(http.StatusBadRequest, config.Response{
			Status:  false,
			Message: find.Error.Error(),
			Data:    nil,
		})
		return
	}

	err := ctx.ShouldBindJSON(&cust)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, config.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	c.DB.Save(&cust)

	ctx.JSON(http.StatusOK, config.Response{
		Status:  true,
		Message: "user updated successfully",
		Data:    cust,
	})
}

func (c *CustomersControl) DeleteCustomer(ctx *gin.Context) {
	var cust models.Customers

	id := ctx.Params.ByName("id")
	find := c.DB.Where("id = ?", id).Delete(&cust)
	if find.Error != nil {
		ctx.JSON(http.StatusBadRequest, config.Response{
			Status:  false,
			Message: find.Error.Error(),
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, config.Response{
		Status:  true,
		Message: "id #" + id + "deleted successfully",
		Data:    cust,
	})
}
