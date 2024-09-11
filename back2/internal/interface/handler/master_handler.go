package handler

import (
	"back2/internal/domain/entity"
	"back2/internal/usecase"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MasterHandler struct {
	useCase *usecase.MasterUseCase
}

func NewMasterHandler(useCase *usecase.MasterUseCase) *MasterHandler {
	return &MasterHandler{useCase: useCase}
}

func (h *MasterHandler) Create(c *gin.Context) {
	var master entity.Master
	if err := c.ShouldBindJSON(&master); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.useCase.Create(&master); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": master})
}

func (h *MasterHandler) GetAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	master, err := h.useCase.List(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": master})
}

func (h *MasterHandler) Get(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("email"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	master, err := h.useCase.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if master == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Master not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": master})
}

func (h *MasterHandler) Update(c *gin.Context) {
	var master entity.Master
	if err := c.ShouldBindJSON(&master); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	master.ID = id

	if err := h.useCase.Update(&master); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": master})
}

func (h *MasterHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.useCase.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Master deleted"})
}

func (h *MasterHandler) ListForDropdown(c *gin.Context) {
	// GroupIDを仮に1に設定
	groupID := 1

	masters, err := h.useCase.ListForDropdown(groupID)
	if err != nil {
		log.Printf("Error in ListForDropdown: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	simplifiedStaffs := make([]gin.H, len(masters))
	for i, master := range masters {
		simplifiedStaffs[i] = gin.H{
			"id":   master.ClassificationCode,
			"name": master.ClassificationName,
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": simplifiedStaffs})
}

func (h *MasterHandler) ListForUsage(c *gin.Context) {
	groupID := 1
	masters, err := h.useCase.ListForUsage(groupID)
	if err != nil {
		log.Printf("Error in ListForUsage: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": masters})
}
