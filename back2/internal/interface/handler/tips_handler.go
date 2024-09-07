package handler

import (
	"back2/internal/domain/entity"
	"back2/internal/usecase"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type TipsHandler struct {
	useCase *usecase.TipsUseCase
}

func NewTipsHandler(useCase *usecase.TipsUseCase) *TipsHandler {
	return &TipsHandler{useCase: useCase}
}

func (h *TipsHandler) Create(c *gin.Context) {
	var tips entity.Tips
	if err := c.ShouldBindJSON(&tips); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(tips)
	// ScheduledTimeが設定されていない場合、現在時刻を使用
	if tips.ScheduledTime.IsZero() {
		tips.ScheduledTime = time.Now()
	}

	err := h.useCase.Create(&tips)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": tips})
}

func (h *TipsHandler) List(c *gin.Context) {
	groupID := 0
	// groupID, err := strconv.ParseInt(c.Query("group_id"), 10, 64)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid group ID"})
	// 	return
	// }
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
		return
	}
	pageSize, err := strconv.Atoi(c.Query("page_size"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page size"})
		return
	}

	tips, err := h.useCase.List(int(groupID), page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": tips})
}

func (h *TipsHandler) Update(c *gin.Context) {
	var tips entity.Tips
	if err := c.ShouldBindJSON(&tips); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	tips.ID = id

	if err := h.useCase.Update(&tips); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": tips})
}

func (h *TipsHandler) Delete(c *gin.Context) {
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

func (h *TipsHandler) UpdateCompletionFlg(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無効なID"})
		return
	}

	// 現在の注文を取得
	tips, err := h.useCase.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "注文の取得に失敗しました: " + err.Error()})
		return
	}
	if tips == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ヒントが見つかりません"})
		return
	}

	tips.CompletionFlg = "1"

	// 更新を実行
	if err := h.useCase.Update(tips); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ヒントの更新に失敗しました: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "ヒントの完了フラグが更新されました"})
}

func (h *TipsHandler) DeleteFlg(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無効なID"})
		return
	}

	tips, err := h.useCase.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ヒントの取得に失敗しました: " + err.Error()})
		return
	}

	tips.IsDeleted = "1"

	if err := h.useCase.Update(tips); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ヒントの更新に失敗しました: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "ヒントの削除フラグが更新されました"})
}

func (h *TipsHandler) ListSchedule(c *gin.Context) {
	startDate := c.Query("start_time")
	endDate := c.Query("end_time")
	groupID := 0

	tips, err := h.useCase.ListSchedule(int(groupID), startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": tips})
}

func (h *TipsHandler) UpdateSchedule(c *gin.Context) {
	id := c.Param("id")
	actualModel := c.Query("actual_model")
	scheduledTimeStr := c.Query("scheduled_time")
	fmt.Println("あくちゅあるもでる", actualModel)
	fmt.Println("すけじゅーるたいむ", scheduledTimeStr)

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
	tipsID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無効なオーダーIDです"})
		return
	}

	// オーダーの取得
	tips, err := h.useCase.GetByID(tipsID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "オーダーの取得に失敗しました: " + err.Error()})
		return
	}

	// オーダーの更新
	tips.ActualModel = actualModel
	tips.ScheduledTime = scheduledTime
	fmt.Println("tips", tips)

	// 更新処理
	err = h.useCase.Update(tips)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "スケジュールが更新されました"})
}

func (h *TipsHandler) UpdateMemo(c *gin.Context) {
	id := c.Param("id")
	actualModel := c.Query("actual_model")
	title := c.Query("title")
	notes := c.Query("notes")
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
	tipsID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無効なヒントIDです"})
		return
	}

	// オーダーの取得
	tips, err := h.useCase.GetByID(tipsID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ヒントの取得に失敗しました: " + err.Error()})
		return
	}

	// オーダーの更新
	tips.ActualModel = actualModel
	tips.ScheduledTime = scheduledTime
	tips.Title = title
	tips.Notes = notes

	// 更新処理
	err = h.useCase.Update(tips)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "スケジュールが更新されました"})
}
