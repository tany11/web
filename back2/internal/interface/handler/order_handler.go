package handler

import (
	"back2/internal/domain/entity"
	"back2/internal/usecase"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	useCase *usecase.OrderUseCase
}

func NewOrderHandler(useCase *usecase.OrderUseCase) *OrderHandler {
	return &OrderHandler{useCase: useCase}
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

	err := h.useCase.Create(&orders)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": orders})
}

func (h *OrderHandler) GetAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	groupID := 0 // 固定値として1を使用

	orders, err := h.useCase.List(groupID, page, pageSize)
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

	orders, err := h.useCase.ListReserved(groupID, page, pageSize)
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

	order, err := h.useCase.GetByID(id)
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

	if err := h.useCase.Update(&order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": order})
}

func (h *OrderHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.useCase.Delete(id); err != nil {
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
	if err := h.useCase.UpdateCompletionFlg(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "注文の更新に失敗しました: " + err.Error()})
		return
	}

	// 更新を実行
	if err := h.useCase.UpdateCompletionFlg(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "注文の更新に失敗しました: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "注文の完了フラグが更新されました"})
}

func (h *OrderHandler) DeleteFlg(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.useCase.UpdateIsDeleted(id); err != nil {
		return
	}

	if err := h.useCase.UpdateIsDeleted(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "注文の削除フラグが更新されました"})
}

func (h *OrderHandler) ListSchedule(c *gin.Context) {
	startDate := c.Query("start_time")
	endDate := c.Query("end_time")

	orders, err := h.useCase.ListSchedule(startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": orders})
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
	order, err := h.useCase.GetByID(orderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "オーダーの取得に失敗しました: " + err.Error()})
		return
	}

	// オーダーの更新
	order.ActualModel = actualModel
	order.ScheduledTime = scheduledTime

	// 更新処理
	err = h.useCase.UpdateSchedule(order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "スケジュールが更新されました"})
}
