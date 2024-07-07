package handler

import (
	"back2/internal/domain/entity"
	"back2/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type StaffHandler struct {
	useCase *usecase.StaffUseCase
}

func NewStaffHandler(useCase *usecase.StaffUseCase) *StaffHandler {
	return &StaffHandler{useCase: useCase}
}

func (h *StaffHandler) Create(c *gin.Context) {
	var staff entity.Staff
	if err := c.ShouldBindJSON(&staff); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.useCase.Create(&staff); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	staff.PasswordHash = "" // パスワードハッシュをレスポンスから除外
	c.JSON(http.StatusOK, gin.H{"data": staff})
}

func (h *StaffHandler) GetAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	groupID, _ := strconv.Atoi(c.DefaultQuery("groupID", "0"))

	staffs, err := h.useCase.List(groupID, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// パスワードハッシュをレスポンスから除外
	for _, s := range staffs {
		s.PasswordHash = ""
	}

	c.JSON(http.StatusOK, gin.H{"data": staffs})
}

func (h *StaffHandler) Get(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	staff, err := h.useCase.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if staff == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Staff not found"})
		return
	}

	staff.PasswordHash = "" // パスワードハッシュをレスポンスから除外
	c.JSON(http.StatusOK, gin.H{"data": staff})
}

func (h *StaffHandler) Update(c *gin.Context) {
	var staff entity.Staff
	if err := c.ShouldBindJSON(&staff); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	staff.ID = id

	if err := h.useCase.Update(&staff); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	staff.PasswordHash = "" // パスワードハッシュをレスポンスから除外
	c.JSON(http.StatusOK, gin.H{"data": staff})
}

func (h *StaffHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.useCase.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Staff deleted"})
}

func (h *StaffHandler) Authenticate(c *gin.Context) {
	var credentials struct {
		StaffID  string `json:"staff_id" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	staff, err := h.useCase.Authenticate(credentials.StaffID, credentials.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	staff.PasswordHash = "" // パスワードハッシュをレスポンスから除外
	c.JSON(http.StatusOK, gin.H{"data": staff})
}

func (h *StaffHandler) GetByEmail(c *gin.Context) {

	c.JSON(http.StatusOK, nil)
}