package entity

import (
	"time"
)

type Tips struct {
	ID            int64     //メモのID
	GroupID       int       //グループID
	ActualModel   string    //実際のモデル
	Title         string    //タイトル
	Notes         string    //メモ
	ScheduledTime time.Time //予定時間
	CreatedAt     time.Time //作成日
	UpdatedAt     time.Time //更新日
	CompletionFlg string    `default:"0"` //完了フラグ
	IsDeleted     string    `default:"0"` //削除フラグ
}
