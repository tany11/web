package handler

import (
	"back2/internal/domain/entity"
	"back2/internal/usecase"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"back2/internal/infrastructure/websocket"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	orderUseCase    usecase.OrderUseCase
	websocketServer *websocket.Server
}

func NewOrderHandler(orderUseCase *usecase.OrderUseCase, websocketServer *websocket.Server) *OrderHandler {
	return &OrderHandler{
		orderUseCase:    *orderUseCase,
		websocketServer: websocketServer,
	}
}

func (h *OrderHandler) Create(c *gin.Context) {
	var orders entity.Orders
	if err := c.ShouldBindJSON(&orders); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	groupID, exists := c.Get("group_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "グループIDが見つかりません"})
		return
	}
	log.Printf("Retrieved group_id: %v, exists: %v", groupID, exists)

	// float64からint64への変換
	groupIDFloat, ok := groupID.(float64)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "グループIDの型が不正です"})
		return
	}
	orders.GroupID = int64(groupIDFloat)

	orders.CompletionFlg = "0"
	orders.IsDeleted = "0"

	if orders.ScheduledTime.IsZero() {
		orders.ScheduledTime = time.Now()
	}

	err := h.orderUseCase.Create(&orders)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	newOrderJSON, _ := json.Marshal(map[string]interface{}{
		"type":  "order_update",
		"order": orders,
	})
	h.websocketServer.BroadcastToGroup(int(orders.GroupID), string(newOrderJSON))

	c.JSON(http.StatusOK, gin.H{"data": orders})
}

func (h *OrderHandler) GetAll(c *gin.Context) {
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

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	orders, err := h.orderUseCase.List(int(groupIDInt), page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "注文リストの取得に失敗しました: " + err.Error()})
		return
	}

	if len(orders) == 0 {
		c.JSON(http.StatusOK, gin.H{"data": []entity.Orders{}, "message": "注文が見つかりません"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": orders})
}

func (h *OrderHandler) ListReserved(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	groupID := 0

	orders, err := h.orderUseCase.ListReserved(groupID, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "予約リストの取得に失敗しました: " + err.Error()})
		return
	}

	if len(orders) == 0 {
		c.JSON(http.StatusOK, gin.H{"data": []entity.Orders{}, "message": "予約が見つかりません"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": orders})
}

func (h *OrderHandler) Get(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	order, err := h.orderUseCase.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if order == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": order})
}

func (h *OrderHandler) Update(c *gin.Context) {
	var order entity.Orders
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	order.ID = id

	if err := h.orderUseCase.Update(&order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	updatedOrderJSON, _ := json.Marshal(map[string]interface{}{
		"type":  "order_update",
		"order": order,
	})
	h.websocketServer.BroadcastToGroup(int(order.GroupID), string(updatedOrderJSON))

	c.JSON(http.StatusOK, gin.H{"data": order})
}

func (h *OrderHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.orderUseCase.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Order deleted"})
}

func (h *OrderHandler) UpdateCompletionFlg(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無効なID"})
		return
	}

	_, err = h.orderUseCase.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "注文の取得に失敗しました: " + err.Error()})
		return
	}

	if err := h.orderUseCase.UpdateCompletionFlg(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "注文の更新に失敗しました: " + err.Error()})
		return
	}

	updatedOrder, err := h.orderUseCase.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新後の注文の取得に失敗しました: " + err.Error()})
		return
	}

	updatedOrderJSON, _ := json.Marshal(map[string]interface{}{
		"type":  "order_update",
		"order": updatedOrder,
	})
	h.websocketServer.BroadcastToGroup(int(updatedOrder.GroupID), string(updatedOrderJSON))

	c.JSON(http.StatusOK, gin.H{"data": "注文の完了フラグが更新されました"})
}

func (h *OrderHandler) DeleteFlg(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	order, err := h.orderUseCase.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "注文の取得に失敗しました: " + err.Error()})
		return
	}

	if err := h.orderUseCase.UpdateIsDeleted(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	deletedOrderJSON, _ := json.Marshal(map[string]interface{}{
		"type":    "order_delete",
		"orderId": id,
	})
	h.websocketServer.BroadcastToGroup(int(order.GroupID), string(deletedOrderJSON))

	c.JSON(http.StatusOK, gin.H{"data": "注文の削除フラグが更新されました"})
}

func (h *OrderHandler) ListSchedule(c *gin.Context) {
	startDate := c.Query("start_time")
	endDate := c.Query("end_time")
	groupID, exists := c.Get("group_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "グループIDが見つかりません"})
		return
	}

	orders, err := h.orderUseCase.ListSchedule(int(groupID.(float64)), startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	activeOrders := make([]entity.Orders, 0)
	for _, order := range orders {
		if order.IsDeleted != "1" {
			activeOrders = append(activeOrders, *order)
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": activeOrders})
}

func (h *OrderHandler) UpdateSchedule(c *gin.Context) {
	id := c.Param("id")
	actualModel := c.Query("actual_model")
	scheduledTimeStr := c.Query("scheduled_time")

	if id == "" || scheduledTimeStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "必要なパラメータが不足しています"})
		return
	}

	scheduledTime, err := time.Parse(time.RFC3339, scheduledTimeStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無効な日時形式です"})
		return
	}

	orderID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無効なオーダーIDです"})
		return
	}

	order, err := h.orderUseCase.GetByID(orderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "オーダーの取得に失敗しました: " + err.Error()})
		return
	}

	order.ActualModel = actualModel
	order.ScheduledTime = scheduledTime

	done := make(chan bool)
	errChan := make(chan error)
	go func() {
		log.Println("Starting UpdateSchedule operation")
		err := h.orderUseCase.UpdateSchedule(order)
		if err != nil {
			log.Printf("Error in UpdateSchedule: %v", err)
			errChan <- err
		} else {
			log.Println("UpdateSchedule operation completed successfully")
			done <- true
		}
	}()

	select {
	case <-done:
		log.Println("UpdateSchedule completed within timeout")
		updatedOrderJSON := map[string]interface{}{
			"type":  "order_update",
			"order": order,
		}
		log.Printf("Broadcasting message: %+v", updatedOrderJSON)
		err := h.websocketServer.BroadcastToGroup(int(order.GroupID), updatedOrderJSON)
		if err != nil {
			log.Printf("Error broadcasting message: %v", err)
		}
		c.JSON(http.StatusOK, gin.H{"message": "スケジュールが更新されました"})
	case err := <-errChan:
		log.Printf("Error occurred during UpdateSchedule: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	case <-time.After(10 * time.Second):
		log.Println("UpdateSchedule operation timed out")
		c.JSON(http.StatusRequestTimeout, gin.H{"error": "処理がタイムアウトしました"})
	}
}
