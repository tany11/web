package controllers

import (
	"back/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
)

type CustomerController struct {
	DB *xorm.Engine
}

// 新しいCustomerを作成します
func (ctrl CustomerController) CreateCustomer(c *gin.Context) {
	var customer models.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := ctrl.DB.Insert(&customer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": customer})
}

// すべてのCustomerを取得します
func (ctrl CustomerController) GetAllCustomers(c *gin.Context) {
	var customers []models.Customer
	if err := ctrl.DB.Find(&customers); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": customers})
}

// 特定のCustomerを取得します
func (ctrl CustomerController) GetCustomer(c *gin.Context) {
	var customer models.Customer
	if _, err := ctrl.DB.Where("id = ?", c.Param("id")).Get(&customer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": customer})
}

func (ctrl CustomerController) GetCustomerPhone(c *gin.Context) {
	phoneNumber := c.Param("phone")
	if phoneNumber == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "電話番号が指定されていません"})
		return
	}

	var customer models.Customer
	has, err := ctrl.DB.Where("phone_number = ?", phoneNumber).Get(&customer)
	log.Printf("Query result: has=%v, customer=%+v, error=%v", has, customer, err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !has {
		c.JSON(http.StatusNotFound, gin.H{"error": "顧客が見つかりません"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": customer})
}

// Customerを更新します
func (ctrl CustomerController) UpdateCustomer(c *gin.Context) {
	var customer models.Customer
	if _, err := ctrl.DB.Where("id = ?", c.Param("id")).Get(&customer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := ctrl.DB.Id(customer.ID).Update(&customer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": customer})
}

// Customerを削除します
func (ctrl CustomerController) DeleteCustomer(c *gin.Context) {
	if _, err := ctrl.DB.Id(c.Param("id")).Delete(&models.Customer{}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Customer deleted"})
}
