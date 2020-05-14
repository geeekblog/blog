package models

type BlogCategoryID uint8

// BlogCategory 博客分类
type BlogCategory struct {
	ID   BlogCategoryID //博客分类ID
	Name string         //博客分类名称
}
