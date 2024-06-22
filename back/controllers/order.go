package controllers

import (
	"back/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
)

type OrderController struct {
	DB *xorm.Engine
}

// 新しいOrderを作成します
func (ctrl OrderController) CreateOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := ctrl.DB.Insert(&order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": order})
}

// すべてのOrderを取得します
func (ctrl OrderController) GetAllOrders(c *gin.Context) {
	var orders []models.Order
	if err := ctrl.DB.Find(&orders); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": orders})
}

// 特定のOrderを取得します
func (ctrl OrderController) GetOrder(c *gin.Context) {
	var order models.Order
	if _, err := ctrl.DB.Where("id = ?", c.Param("id")).Get(&order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": order})
}

// Orderを更新します
func (ctrl OrderController) UpdateOrder(c *gin.Context) {
	var order models.Order
	if _, err := ctrl.DB.Where("id = ?", c.Param("id")).Get(&order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := ctrl.DB.Id(order.ID).Update(&order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": order})
}

// Orderを削除します
func (ctrl OrderController) DeleteOrder(c *gin.Context) {
	if _, err := ctrl.DB.Id(c.Param("id")).Delete(&models.Order{}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Order deleted"})
}
