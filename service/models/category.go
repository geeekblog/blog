package models

type BlogCategoryID uint8

// BlogCategory 博客分类
type BlogCategory struct {
	ID   BlogCategoryID //文章分类ID
	Name string         //文章分类名称
	Uri  string         //文章地址
}
