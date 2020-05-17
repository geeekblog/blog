package entities

import (
	"blog/models"
	"time"
)

type Blog struct {
	ID         string                `xorm:"id char(36) pk"`
	UserID     string                `xorm:"user_id char(36)"`
	Title      string                `xorm:"title"`
	Content    string                `xorm:"body"`
	Tags       []int                 `xorm:"tags"`
	Pics       []string              `xorm:"pics"`
	Category   models.BlogCategoryID `xorm:"category"`
	CreateTime time.Time             `xorm:"create_time"`
	UpdateTime time.Time             `xorm:"update_time"`
}
