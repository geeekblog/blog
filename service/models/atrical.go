package models

import "time"

type Artical struct {
	ID         string       //文章ID：uuid
	UserID     string       //用户ID: uuid
	Title      string       //文章标题
	Content    string       //文章内容
	Tags       []int        //文章tag数组
	Pics       []string     //文章图片数组
	CreateTime time.Time    //第一次创建时间
	UpdateTime time.Time    //最后一次修改的时间
	Category   BlogCategory //文章分类
}
