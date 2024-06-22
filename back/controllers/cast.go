package controllers

import (
	"back/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
)

type CastLoginController struct {
	DB *xorm.Engine
}

// 新しいCastLoginを作成します
func (ctrl CastLoginController) CreateCastLogin(c *gin.Context) {
	var castLogin models.CastLogin
	if err := c.ShouldBindJSON(&castLogin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := ctrl.DB.Insert(&castLogin); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": castLogin})
}

// すべてのCastLoginを取得します
func (ctrl CastLoginController) GetAllCastLogins(c *gin.Context) {
	var castLogins []models.CastLogin
	if err := ctrl.DB.Find(&castLogins); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": castLogins})
}

// 特定のCastLoginを取得します
func (ctrl CastLoginController) GetCastLogin(c *gin.Context) {
	var castLogin models.CastLogin
	if _, err := ctrl.DB.Where("id = ?", c.Param("id")).Get(&castLogin); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": castLogin})
}

// CastLoginを更新します
func (ctrl CastLoginController) UpdateCastLogin(c *gin.Context) {
	var castLogin models.CastLogin
	if _, err := ctrl.DB.Where("id = ?", c.Param("id")).Get(&castLogin); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&castLogin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := ctrl.DB.Id(castLogin.ID).Update(&castLogin); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": castLogin})
}

// CastLoginを削除します
func (ctrl CastLoginController) DeleteCastLogin(c *gin.Context) {
	if _, err := ctrl.DB.Id(c.Param("id")).Delete(&models.CastLogin{}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "CastLogin deleted"})
}
