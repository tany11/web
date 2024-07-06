package handler

import (
	"back2/internal/domain/entity"
	"back2/internal/usecase"
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
	var cast entity.Cast
	if err := c.ShouldBindJSON(&cast); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.useCase.Create(&cast); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": cast})
}

func (h *CastHandler) GetAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	groupID, _ := strconv.Atoi(c.DefaultQuery("groupID", "0"))

	casts, err := h.useCase.List(groupID, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
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
