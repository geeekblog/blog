package models

//页面的内容
type Page struct {
	UUID  string
	Title string
	Uri   string
	Body  []byte
}
