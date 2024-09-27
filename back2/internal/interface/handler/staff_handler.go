package handler

import (
	"back2/internal/domain/entity"
	"back2/internal/usecase"
	"log"
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
	groupID, exists := c.Get("group_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "グループIDが見つかりません"})
		return
	}
	staff.GroupID = int64(groupID.(float64))

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
	groupID, exists := c.Get("group_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "グループIDが見つかりません"})
		return
	}

	groupIDInt, ok := groupID.(float64)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "グループIDの型が不正です"})
		return
	}

	staffs, err := h.useCase.List(int(groupIDInt))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "スタッフの取得中にエラーが発生しました"})
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

func (h *StaffHandler) ListForDropdown(c *gin.Context) {
	// GroupIDを仮に1に設定
	groupID := 1

	staffs, err := h.useCase.ListForDropdown(groupID)
	if err != nil {
		log.Printf("Error in ListForDropdown: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	simplifiedStaffs := make([]gin.H, len(staffs))
	for i, staff := range staffs {
		simplifiedStaffs[i] = gin.H{
			"id":         staff.ID,
			"staff_id":   staff.StaffID,
			"name":       staff.StaffLastName,
			"office_flg": staff.OfficeFlg,
			"driver_flg": staff.DriverFlg,
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": simplifiedStaffs})
}
