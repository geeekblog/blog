package entities

import (
	"blog/models"
	"time"
)

type Blog struct {
	ID         string                `xorm:"id char(36) pk"`
	UserID     string                `xorm:"user_id char(36)"`
	CreateTime time.Time             `xorm:"create_time"`
	UpdateTime time.Time             `xorm:"update_time"`
	Title      string                `xorm:"title"`
	Content    string                `xorm:"body"`
	Category   models.BlogCategoryID `xorm:"category"`
	Tags       []int                 `xorm:"tags"`
	Pics       []string              `xorm:"pics"`
}
