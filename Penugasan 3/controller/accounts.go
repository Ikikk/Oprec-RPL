package controller

import (
	"Penugasan-3/config"
	"Penugasan-3/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AccountControl struct {
	DB *gorm.DB
}

func (c *AccountControl) PostAccount(ctx *gin.Context) {
	var account models.Accounts

	err := ctx.ShouldBindJSON(&account)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, config.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	tx := c.DB.Create(&account)
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
		Message: "account created successfully",
		Data:    account,
	})
}

// func (c *CustomersControl) UpdateBalance(ctx *gin.Context) {
// 	var account models.Accounts

// 	find := c.DB.Preload("Accounts").First(&account)
// 	if find.Error != nil {
// 		ctx.JSON(http.StatusBadRequest, config.Response{
// 			Status:  false,
// 			Message: find.Error.Error(),
// 			Data:    nil,
// 		})
// 		return
// 	}

// 	balance := ctx.Params.ByName("balance")

// 	update := c.DB.Select("balance")

// 	err := ctx.ShouldBindJSON(&account)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, config.Response{
// 			Status:  false,
// 			Message: err.Error(),
// 			Data:    nil,
// 		})
// 		return
// 	}
// 	c.DB.Save(&account)

// 	ctx.JSON(http.StatusOK, config.Response{
// 		Status:  true,
// 		Message: "cash out successfully",
// 		Data:    account,
// 	})
// }

func (c *AccountControl) GetAllAccountByUser(ctx *gin.Context) {
	var account models.Accounts

	id := ctx.Params.ByName("cust_id")
	tx := c.DB.Where("cust_id = ?", id).Take(&account)
	if tx.Error != nil {
		ctx.JSON(http.StatusBadRequest,
			config.Response{
				Status:  false,
				Message: tx.Error.Error(),
				Data:    nil,
			})
		return
	}

	ctx.JSON(http.StatusOK,
		config.Response{
			Status:  true,
			Message: "get account successfully",
			Data:    account,
		})
}

func (c *AccountControl) DeleteAccount(ctx *gin.Context) {
	var account models.Accounts

	custid := ctx.Params.ByName("cust_id")
	find := c.DB.Where("cust_id= ?", custid).Delete(&account)
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
		Message: "account by id # " + custid + "deleted successfully",
		Data:    account,
	})
}
