package entity

import "time"

type Master struct {
	ID                 int64     `xorm:"pk autoincr"`
	ClassificationName string    `xorm:"varchar(30) notnull"`
	DisplayName        string    `xorm:"varchar(30) notnull"`
	ClassificationCode string    `xorm:"varchar(30) notnull"`
	CreatedAt          time.Time `xorm:"created"`
	UpdatedAt          time.Time `xorm:"updated"`
}
