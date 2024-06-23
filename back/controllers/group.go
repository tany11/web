package controllers

import (
	"back/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
)

type GroupLoginController struct {
	DB *xorm.Engine
}

// 新しいGroupLoginを作成します
func (ctrl GroupLoginController) CreateGroupLogin(c *gin.Context) {
	var groupLogin models.GroupLogin
	if err := c.ShouldBindJSON(&groupLogin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := ctrl.DB.Insert(&groupLogin); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": groupLogin})
}

// すべてのGroupLoginを取得します
func (ctrl GroupLoginController) GetAllGroupLogins(c *gin.Context) {
	var groupLogins []models.GroupLogin
	if err := ctrl.DB.Find(&groupLogins); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": groupLogins})
}

// 特定のGroupLoginを取得します
func (ctrl GroupLoginController) GetGroupLogin(c *gin.Context) {
	var groupLogin models.GroupLogin
	if _, err := ctrl.DB.Where("email = ?", c.Param("email")).Get(&groupLogin); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": groupLogin})
}

// GroupLoginを更新します
func (ctrl GroupLoginController) UpdateGroupLogin(c *gin.Context) {
	var groupLogin models.GroupLogin
	if _, err := ctrl.DB.Where("id = ?", c.Param("id")).Get(&groupLogin); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&groupLogin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := ctrl.DB.Id(groupLogin.ID).Update(&groupLogin); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": groupLogin})
}

// GroupLoginを削除します
func (ctrl GroupLoginController) DeleteGroupLogin(c *gin.Context) {
	if _, err := ctrl.DB.Id(c.Param("id")).Delete(&models.GroupLogin{}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "GroupLogin deleted"})
}
