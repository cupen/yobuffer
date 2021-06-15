package models

type Struct struct {
	ID       int
	Name     string
	Desc     string
	Encoding string
	Fields   []*Field
	Source   string
}
