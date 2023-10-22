package ast

type Meta struct {
	Fields []*MetaField
}

type MetaField struct {
	Name  string
	Type  string
	Value string
}
