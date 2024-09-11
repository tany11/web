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

	orders.CompletionFlg = "0"
	orders.IsDeleted = "0"

	// ScheduledTimeが設定されていない場合、現在時刻を使用
	if orders.ScheduledTime.IsZero() {
		orders.ScheduledTime = time.Now()
	}

	err := h.orderUseCase.Create(&orders)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// WebSocket通を送信
	newOrderJSON, _ := json.Marshal(map[string]interface{}{
		"type":  "order_update",
		"order": orders,
	})
	h.websocketServer.BroadcastMessage(newOrderJSON)

	c.JSON(http.StatusOK, gin.H{"data": orders})
}

func (h *OrderHandler) GetAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	groupID := 0 // 固定値として1を使用

	orders, err := h.orderUseCase.List(groupID, page, pageSize)
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

	groupID := 0 // 固定値として1を使用

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

	// WebSocket通知を送信
	updatedOrderJSON, _ := json.Marshal(map[string]interface{}{
		"type":  "order_update",
		"order": order,
	})
	h.websocketServer.BroadcastMessage(updatedOrderJSON)

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

	// 現在の注文を取得
	_, err = h.orderUseCase.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "注文の取得に失敗しました: " + err.Error()})
		return
	}

	// 更新を実行
	if err := h.orderUseCase.UpdateCompletionFlg(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "注文の更新に失敗しました: " + err.Error()})
		return
	}

	// 更新後の注文を再取得
	updatedOrder, err := h.orderUseCase.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新後の注文の取得に失敗しました: " + err.Error()})
		return
	}

	// WebSocket通知を送信
	updatedOrderJSON, _ := json.Marshal(map[string]interface{}{
		"type":  "order_update",
		"order": updatedOrder,
	})
	h.websocketServer.BroadcastMessage(updatedOrderJSON)

	c.JSON(http.StatusOK, gin.H{"data": "注文の完了フラグが更新されました"})
}

func (h *OrderHandler) DeleteFlg(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.orderUseCase.UpdateIsDeleted(id); err != nil {
		return
	}

	// WebSocket通知を送信
	deletedOrderJSON, _ := json.Marshal(map[string]interface{}{
		"type":    "order_delete",
		"orderId": id,
	})
	h.websocketServer.BroadcastMessage(deletedOrderJSON)

	if err := h.orderUseCase.UpdateIsDeleted(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "注文の削除フラグが更新されました"})
}

func (h *OrderHandler) ListSchedule(c *gin.Context) {
	startDate := c.Query("start_time")
	endDate := c.Query("end_time")

	orders, err := h.orderUseCase.ListSchedule(startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 削除フラグが立っていないオーダーのみをフィルタリング
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

	// パラメータの検証
	if id == "" || scheduledTimeStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "必要なパラメータが不足しています"})
		return
	}

	// 日時の解析
	scheduledTime, err := time.Parse(time.RFC3339, scheduledTimeStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無効な日時形式です"})
		return
	}

	// IDの解析
	orderID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無効なオーダーIDです"})
		return
	}

	// オーダーの取得
	order, err := h.orderUseCase.GetByID(orderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "オーダーの取得に失敗しました: " + err.Error()})
		return
	}

	// オーダーの更新
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
		// オーダーの更新が成功した場合、WebSocketを通じて通知
		updatedOrderJSON, _ := json.Marshal(map[string]interface{}{
			"type":  "order_update",
			"order": order,
		})
		log.Printf("Broadcasting message: %s", string(updatedOrderJSON))
		h.websocketServer.BroadcastMessage(updatedOrderJSON)
		c.JSON(http.StatusOK, gin.H{"message": "スケジュールが更新されました"})
	case err := <-errChan:
		log.Printf("Error occurred during UpdateSchedule: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	case <-time.After(10 * time.Second):
		log.Println("UpdateSchedule operation timed out")
		c.JSON(http.StatusRequestTimeout, gin.H{"error": "処理がタイムアウトしました"})
	}
}
