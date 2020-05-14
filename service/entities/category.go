package entities

import (
	"blog/models"
	"time"
)

type BlogCategory struct {
	ID         models.BlogCategoryID `xorm:"id pk"`
	Name       string                `xorm:"name"`
	CreateTime time.Time             `xorm:"create_time"`
}
