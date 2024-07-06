package controllers

import (
	"back/models"
	"back/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CastLoginController struct {
	Service *service.CastLoginService
}

func NewCastLoginController(service *service.CastLoginService) *CastLoginController {
	return &CastLoginController{Service: service}
}

// 新しいCastLoginを作成します
func (ctrl *CastLoginController) CreateCastLogin(c *gin.Context) {
	var castLogin models.CastLogin
	if err := c.ShouldBindJSON(&castLogin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.Service.Create(&castLogin); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": castLogin})
}

// すべてのCastLoginを取得します
func (ctrl *CastLoginController) GetAllCastLogins(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	castLogins, err := ctrl.Service.List(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": castLogins})
}

// 特定のCastLoginを取得します
func (ctrl *CastLoginController) GetCastLogin(c *gin.Context) {
	castLogin, err := ctrl.Service.GetByID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if castLogin == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "CastLogin not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": castLogin})
}

// CastLoginを更新します
func (ctrl *CastLoginController) UpdateCastLogin(c *gin.Context) {
	var castLogin models.CastLogin
	if err := c.ShouldBindJSON(&castLogin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	castLogin.ID = id

	if err := ctrl.Service.Update(&castLogin); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": castLogin})
}

// CastLoginを削除します
func (ctrl *CastLoginController) DeleteCastLogin(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := ctrl.Service.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "CastLogin deleted"})
}
