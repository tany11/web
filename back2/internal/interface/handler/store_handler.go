package handler

import (
	"back2/internal/domain/entity"
	"back2/internal/usecase"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type StoreHandler struct {
	useCase *usecase.StoreUseCase
}

func NewStoreHandler(useCase *usecase.StoreUseCase) *StoreHandler {
	return &StoreHandler{useCase: useCase}
}

func (h *StoreHandler) Create(c *gin.Context) {
	var store entity.Store
	store.GroupID = 1
	if err := c.ShouldBindJSON(&store); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.useCase.Create(&store); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": store})
}

func (h *StoreHandler) GetAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	groups, err := h.useCase.List(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": groups})
}

func (h *StoreHandler) Get(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("email"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	group, err := h.useCase.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if group == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Store not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": group})
}

func (h *StoreHandler) Update(c *gin.Context) {
	var store entity.Store
	if err := c.ShouldBindJSON(&store); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	store.ID = id

	if err := h.useCase.Update(&store); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": store})
}

func (h *StoreHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.useCase.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Store deleted"})
}

func (h *StoreHandler) ListForDropdown(c *gin.Context) {
	// GroupIDを仮に1に設定
	groupID := 1

	stores, err := h.useCase.ListForDropdown(groupID)
	if err != nil {
		log.Printf("Error in ListForDropdown: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	simplifiedStaffs := make([]gin.H, len(stores))
	for i, store := range stores {
		simplifiedStaffs[i] = gin.H{
			"id":   store.ID,
			"name": store.StoreName,
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": simplifiedStaffs})
}
