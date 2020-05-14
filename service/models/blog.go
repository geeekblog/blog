package models

import "time"

type Blog struct {
	ID         string       //博客ID：uuid
	UserID     string       //用户ID: uuid
	CreateTime time.Time    //第一次创建时间
	UpdateTime time.Time    //最后一次修改的时间
	Title      string       //博客标题
	Content    string       //博客内容
	Category   BlogCategory //博客分类
	Tags       []int        //博客tag数组
	Pics       []string     //博客图片数组
}
