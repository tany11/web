package handler

import (
	"back2/internal/domain/entity"
	"back2/internal/usecase"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CastHandler struct {
	useCase *usecase.CastUseCase
}

func NewCastHandler(useCase *usecase.CastUseCase) *CastHandler {
	return &CastHandler{useCase: useCase}
}

func (h *CastHandler) Create(c *gin.Context) {
	var castInput struct {
		entity.Cast
	}
	if err := c.ShouldBindJSON(&castInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cast := castInput.Cast
	cast.IsDeleted = "0"

	if err := h.useCase.Create(&cast); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": cast})
}

func (h *CastHandler) GetAll(c *gin.Context) {
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

	casts, err := h.useCase.List(int(groupIDInt))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// パスワードハッシュをレスポンスから除外
	for _, c := range casts {
		c.PasswordHash = ""
	}

	c.JSON(http.StatusOK, gin.H{"data": casts})
}

func (h *CastHandler) Get(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	cast, err := h.useCase.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if cast == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cast not found"})
		return
	}
	cast.PasswordHash = ""

	c.JSON(http.StatusOK, gin.H{"data": cast})
}

func (h *CastHandler) Update(c *gin.Context) {
	var cast entity.Cast
	if err := c.ShouldBindJSON(&cast); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	cast.ID = id

	log.Printf("cast: %v", cast)

	if err := h.useCase.Update(&cast); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": cast})
}

func (h *CastHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.useCase.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Cast deleted"})
}

func (h *CastHandler) ListForDropdown(c *gin.Context) {
	// GroupIDを仮に1に設定
	groupID := 1

	casts, err := h.useCase.ListForDropdown(groupID)
	if err != nil {
		log.Printf("Error in ListForDropdown: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	simplifiedStaffs := make([]gin.H, len(casts))
	for i, cast := range casts {
		simplifiedStaffs[i] = gin.H{
			"id":          cast.ID,
			"cast_id":     cast.CastID,
			"name":        cast.CastName,
			"working_flg": cast.WorkingFlg,
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": simplifiedStaffs})
}

func (h *CastHandler) UpdateWorkingFlg(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.useCase.UpdateWorkingFlg(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Cast updated"})
}
