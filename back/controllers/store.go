package controllers

import (
	"back/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
)

type StoreController struct {
	DB *xorm.Engine
}

// 新しいStoreを作成します
func (ctrl StoreController) CreateStore(c *gin.Context) {
	var store models.Store
	if err := c.ShouldBindJSON(&store); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := ctrl.DB.Insert(&store); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": store})
}

// すべてのStoreを取得します
func (ctrl StoreController) GetAllStores(c *gin.Context) {
	var stores []models.Store
	if err := ctrl.DB.Find(&stores); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": stores})
}

// 特定のStoreを取得します
func (ctrl StoreController) GetStore(c *gin.Context) {
	var store models.Store
	if _, err := ctrl.DB.Where("id = ?", c.Param("id")).Get(&store); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": store})
}

// Storeを更新します
func (ctrl StoreController) UpdateStore(c *gin.Context) {
	var store models.Store
	if _, err := ctrl.DB.Where("id = ?", c.Param("id")).Get(&store); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&store); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := ctrl.DB.Id(store.ID).Update(&store); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": store})
}

// Storeを削除します
func (ctrl StoreController) DeleteStore(c *gin.Context) {
	if _, err := ctrl.DB.Id(c.Param("id")).Delete(&models.Store{}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Store deleted"})
}
