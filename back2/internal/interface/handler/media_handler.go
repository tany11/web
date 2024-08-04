package handler

import (
	"back2/internal/domain/entity"
	"back2/internal/usecase"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MediaHandler struct {
	useCase *usecase.MediaUseCase
}

func NewMediaHandler(useCase *usecase.MediaUseCase) *MediaHandler {
	return &MediaHandler{useCase: useCase}
}

func (h *MediaHandler) Create(c *gin.Context) {
	var media entity.Media
	if err := c.ShouldBindJSON(&media); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.useCase.Create(&media); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": media})
}

func (h *MediaHandler) GetAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	media, err := h.useCase.List(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": media})
}

func (h *MediaHandler) Get(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("email"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	media, err := h.useCase.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if media == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Media not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": media})
}

func (h *MediaHandler) Update(c *gin.Context) {
	var media entity.Media
	if err := c.ShouldBindJSON(&media); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	media.ID = id

	if err := h.useCase.Update(&media); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": media})
}

func (h *MediaHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.useCase.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Group deleted"})
}

func (h *MediaHandler) ListForDropdown(c *gin.Context) {
	// GroupIDを仮に1に設定
	groupID := 1

	medias, err := h.useCase.ListForDropdown(groupID)
	if err != nil {
		log.Printf("Error in ListForDropdown: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	simplifiedStaffs := make([]gin.H, len(medias))
	for i, media := range medias {
		simplifiedStaffs[i] = gin.H{
			"id":   media.ID,
			"name": media.MediaName,
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": simplifiedStaffs})
}
