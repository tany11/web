package controllers

import (
	"back/models"
	"back/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
)

type OrderController struct {
	DB              *xorm.Engine
	OrderService    *service.OrderService
	CustomerService *service.CustomerService
}

// CreateOrder は新しい注文を作成し、必要に応じて顧客も作成します
func (ctrl OrderController) CreateOrder(c *gin.Context) {
	var orderRequest struct {
		PhoneNumber       string `json:"phoneNumber"`
		CustomerName      string `json:"customerName"`
		ModelName         string `json:"modelName"`
		Price             string `json:"price"`
		PostalCode        string `json:"postalCode"`
		Address           string `json:"address"`
		NominationFee     string `json:"nominationFee"`
		TransportationFee string `json:"transportationFee"`
		TravelExpenses    string `json:"travelExpenses"`
	}

	if err := c.ShouldBindJSON(&orderRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無効なリクエストデータ: " + err.Error()})
		return
	}

	// 顧客を電話番号で検索
	customer, err := ctrl.CustomerService.GetByPhone(orderRequest.PhoneNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "顧客の検索中にエラーが発生しました: " + err.Error()})
		return
	}

	// 顧客が存在しない場合は新規作成
	if customer == nil {
		customer = &models.Customer{
			CustomerName: orderRequest.CustomerName,
			PhoneNumber:  orderRequest.PhoneNumber,
			Address:      orderRequest.Address,
		}
		if err := ctrl.CustomerService.Create(customer); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "顧客の作成中にエラーが発生しました: " + err.Error()})
			return
		}
	}

	// 注文を作成
	order := &models.Order{
		GroupID:        1, // 適切な値を設定
		StoreID:        1, // 適切な値を設定
		CustomerID:     int(customer.ID),
		CustomerName:   customer.CustomerName,
		ModelName:      orderRequest.ModelName,
		OrderAmount:    parseIntOrZero(orderRequest.Price),
		PostalCode:     orderRequest.PostalCode,
		Address:        orderRequest.Address,
		ReservationFee: parseIntOrZero(orderRequest.NominationFee),
		TravelCost:     parseIntOrZero(orderRequest.TransportationFee),
		OtherFee:       parseIntOrZero(orderRequest.TravelExpenses),
	}

	if err := ctrl.OrderService.Create(order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "注文の作成中にエラーが発生しました: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": order})
}

// parseIntOrZero は文字列を整数に変換し、エラーの場合は0を返します
func parseIntOrZero(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

// すべてのOrderを取得します
func (ctrl OrderController) GetAllOrders(c *gin.Context) {
	var orders []models.Order
	if err := ctrl.DB.Find(&orders); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": orders})
}

// 特定のOrderを取得します
func (ctrl OrderController) GetOrder(c *gin.Context) {
	var order models.Order
	if _, err := ctrl.DB.Where("id = ?", c.Param("id")).Get(&order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": order})
}

// Orderを更新します
func (ctrl OrderController) UpdateOrder(c *gin.Context) {
	var order models.Order
	if _, err := ctrl.DB.Where("id = ?", c.Param("id")).Get(&order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := ctrl.DB.Id(order.ID).Update(&order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": order})
}

// Orderを削除します
func (ctrl OrderController) DeleteOrder(c *gin.Context) {
	if _, err := ctrl.DB.Id(c.Param("id")).Delete(&models.Order{}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Order deleted"})
}
