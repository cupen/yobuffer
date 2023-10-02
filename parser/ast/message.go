package ast

type Struct struct {
	Name   string
	Meta   *Meta
	Fields []*StructField
}

type BeginPackage struct {
	Name string
}

type StructBegin struct {
	Name string
}

type StructField struct {
	Name string
	Type string
	Tag  int
}

type ExternalStruct struct {
	Name string
}
