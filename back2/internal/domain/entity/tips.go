package entity

import (
	"time"
)

type Tips struct {
	ID            int64     //メモのID
	GroupID       int64     //グループID
	ActualModel   string    //実際のモデル
	Title         string    //タイトル
	Notes         string    //メモ
	ScheduledTime time.Time //予定時間
	ScheduledBox  string    `xorm:"varchar(10) notnull default '60'"` //予定ボックス
	CreatedAt     time.Time //作成日
	UpdatedAt     time.Time //更新日
	CompletionFlg string    `xorm:"varchar(1) notnull default '0'"` //完了フラグ
	IsDeleted     string    `xorm:"varchar(1) notnull default '0'"` //削除フラグ
}
