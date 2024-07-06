package controllers

import (
	"back/models"
	"back/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type StaffLoginController struct {
	Service *service.StaffLoginService
}

func NewStaffLoginController(service *service.StaffLoginService) *StaffLoginController {
	return &StaffLoginController{Service: service}
}

// 新しいStaffLoginを作成します
func (ctrl *StaffLoginController) CreateStaffLogin(c *gin.Context) {
	var staffLogin models.StaffLogin
	if err := c.ShouldBindJSON(&staffLogin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.Service.Create(&staffLogin); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": staffLogin})
}

// すべてのStaffLoginを取得します
func (ctrl *StaffLoginController) GetAllStaffLogins(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	staffLogins, err := ctrl.Service.List(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": staffLogins})
}

// 特定のStaffLoginを取得します
func (ctrl *StaffLoginController) GetStaffLogin(c *gin.Context) {
	staffLogin, err := ctrl.Service.GetByID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if staffLogin == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "StaffLogin not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": staffLogin})
}

// StaffLoginを更新します
func (ctrl *StaffLoginController) UpdateStaffLogin(c *gin.Context) {
	var staffLogin models.StaffLogin
	if err := c.ShouldBindJSON(&staffLogin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	staffLogin.ID = id

	if err := ctrl.Service.Update(&staffLogin); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": staffLogin})
}

// StaffLoginを削除します
func (ctrl *StaffLoginController) DeleteStaffLogin(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := ctrl.Service.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "StaffLogin deleted"})
}
