package ast

type Service struct {
	Name string
	RPCs []*RPC
}

type RPC struct {
	Name   string
	Fields []*StructField
	Return string
}
