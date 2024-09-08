package handler

import (
	"back2/internal/domain/entity"
	"back2/internal/usecase"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CustomerHandler struct {
	useCase      *usecase.CustomerUseCase
	useCaseOrder *usecase.OrderUseCase
}

func NewCustomerHandler(useCase *usecase.CustomerUseCase, useCaseOrder *usecase.OrderUseCase) *CustomerHandler {
	return &CustomerHandler{
		useCase:      useCase,
		useCaseOrder: useCaseOrder,
	}
}

func (h *CustomerHandler) Create(c *gin.Context) {
	var customer entity.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.useCase.Create(&customer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": customer})
}

func (h *CustomerHandler) GetAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	groupID, _ := strconv.Atoi(c.DefaultQuery("groupID", "0"))

	customers, err := h.useCase.List(groupID, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": customers})
}

func (h *CustomerHandler) Get(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	customer, err := h.useCase.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if customer == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": customer})
}

func (h *CustomerHandler) GetByPhone(c *gin.Context) {
	phoneNumber := c.Param("phone")
	customer, err := h.useCase.GetByPhoneNumber(phoneNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if customer == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": customer})
}

func (h *CustomerHandler) Update(c *gin.Context) {
	var customer entity.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	customer.ID = id

	if err := h.useCase.Update(&customer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": customer})
}

func (h *CustomerHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.useCase.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Customer deleted"})
}

func (h *CustomerHandler) List(c *gin.Context) {
	customers, err := h.useCase.List(0, 1, 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": customers})
}

func (h *CustomerHandler) GetDetail(c *gin.Context) {

	var customerOrder entity.CustomerOrder
	phoneNumber := c.Param("phone")

	customer, err := h.useCase.GetByPhoneNumber(phoneNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if customer == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "顧客が見つかりません"})
		return
	}

	orders, err := h.useCaseOrder.GetByCustomerID(int(customer.ID))
	if err != nil {
		log.Printf("注文情報の取得に失敗: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "注文情報の取得に失敗しました"})
		return
	}

	totalPrice, totalUseTime, err := h.useCaseOrder.GetTotalPriceAndUseTime(int(customer.ID))
	if err != nil {
		log.Printf("注文情報の取得に失敗: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "注文情報の取得に失敗しました"})
		return
	}

	orderSlice := make([]entity.Orders, len(orders))
	for i, order := range orders {
		orderSlice[i] = *order
	}

	customerOrder = entity.CustomerOrder{
		ID:           customer.ID,
		CustomerName: customer.CustomerName,
		PhoneNumber:  customer.PhoneNumber,
		Memo:         customer.Memo,
		City1:        customer.City1,
		Address1:     customer.Address1,
		City2:        customer.City2,
		Address2:     customer.Address2,
		City3:        customer.City3,
		Address3:     customer.Address3,
		TotalPrice:   totalPrice,
		TotalUseTime: totalUseTime,
		OrderList:    orderSlice,
	}
	log.Printf("レスポンス用のデータ: %+v", customerOrder)

	c.JSON(http.StatusOK, gin.H{"data": customerOrder})
}

func (h *CustomerHandler) GetSearchList(c *gin.Context) {
	var customerSearch entity.CustomerSearch
	customerSearch.PhoneLast4 = c.Query("phoneLast4")
	customerSearch.CastID = c.Query("castID")
	customerSearch.StoreID = c.Query("storeID")
	customerSearch.CreatedFrom = c.Query("createdFrom")
	customerSearch.CreatedTo = c.Query("createdTo")

	log.Printf("受け取ったパラメータ: %+v", customerSearch)

	customers, err := h.useCase.GetSearchList(customerSearch)
	if err != nil {
		log.Printf("顧客検索エラー: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": customers})
}
